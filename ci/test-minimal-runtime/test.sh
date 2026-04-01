#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
MODULE_PATH=$(awk '/^module /{print $2}' "$REPO_ROOT/go.mod")
WASMTIME_VERSION=$(python3 -c "exec(open('$REPO_ROOT/ci/download-wasmtime.py').read()); print(version)")

WORK_DIR=$(mktemp -d)
trap 'rm -rf "$WORK_DIR"' EXIT
cd "$WORK_DIR"

go mod init test-minimal-runtime
cat >> go.mod <<EOF
require "$MODULE_PATH" v0.0.0
replace "$MODULE_PATH" v0.0.0 => "$REPO_ROOT"
EOF
go mod tidy

# Step 1: Create a pre-compiled module using the full Wasmtime library.
cat > create_cwasm.go <<GOEOF
//go:build ignore

package main

import (
	"log"
	"os"

	wasmtime "$MODULE_PATH"
)

func main() {
	wasm, err := wasmtime.Wat2Wasm(\`(module (func (export "test") (result i32) (i32.const 1)))\`)
	check(err)

	cfg := wasmtime.NewConfig()
	cfg.SetGCSupport(false)
	cfg.SetWasmThreads(false)
	cfg.SetWasmComponentModel(false)
	engine := wasmtime.NewEngineWithConfig(cfg)
	module, err := wasmtime.NewModule(engine, wasm)
	check(err)
	defer module.Close()

	artifact, err := module.Serialize()
	check(err)

	check(os.WriteFile("module.cwasm", artifact, 0o644))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
GOEOF
go run create_cwasm.go

# Step 2: Download the minimal Wasmtime static library for the current platform.
case "$(uname -s)-$(uname -m)" in
	Linux-x86_64)  ARCHIVE="wasmtime-${WASMTIME_VERSION}-x86_64-linux-c-api.tar.xz";  BUILD_DIR="linux-x86_64";;
	Linux-aarch64) ARCHIVE="wasmtime-${WASMTIME_VERSION}-aarch64-linux-c-api.tar.xz"; BUILD_DIR="linux-aarch64";;
	Darwin-x86_64) ARCHIVE="wasmtime-${WASMTIME_VERSION}-x86_64-macos-c-api.tar.xz";  BUILD_DIR="macos-x86_64";;
	Darwin-arm64)  ARCHIVE="wasmtime-${WASMTIME_VERSION}-aarch64-macos-c-api.tar.xz"; BUILD_DIR="macos-aarch64";;
	*) echo "Unsupported platform: $(uname -s)-$(uname -m)" >&2; exit 1;;
esac
URL="https://github.com/bytecodealliance/wasmtime/releases/download/${WASMTIME_VERSION}/${ARCHIVE}"
echo "Downloading min library from ${URL}"
curl -sSLf "$URL" | tar xJ --strip-components=2 "${ARCHIVE%.tar.xz}/min/lib"

# Step 3: Vendor the module, then adjust the vendored copy so it compiles
# against the minimal Wasmtime library:
#   a) Remove Go source files that call C functions absent from the min binary.
#   b) Copy the min static library over the full one so CGO links the right lib.
go mod vendor
VENDOR_PKG="vendor/${MODULE_PATH}"
rm -f "$VENDOR_PKG"/wat2wasm.go
rm -f "$VENDOR_PKG"/wasi.go
rm -f "$VENDOR_PKG"/config_feat_*.go
rm -f "$VENDOR_PKG"/module_feat_*.go
rm -f "$VENDOR_PKG"/module_feats_*.go
rm -f "$VENDOR_PKG"/linker_feat_*.go
rm -f "$VENDOR_PKG"/store_feat_*.go
cp -f lib/libwasmtime.a "$VENDOR_PKG/build/${BUILD_DIR}/libwasmtime.a"

# Step 4: Test that the minimal Wasmtime binary can deserialize and run a module.
cat > min_test.go <<GOEOF
package testminimalruntime_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	wasmtime "$MODULE_PATH"
)

func TestMinimalRuntime(t *testing.T) {
	cfg := wasmtime.NewConfig()
	cfg.SetGCSupport(false)
	engine := wasmtime.NewEngineWithConfig(cfg)
	module, err := wasmtime.NewModuleDeserializeFile(engine, "module.cwasm")
	require.NoError(t, err)
	defer module.Close()

	store := wasmtime.NewStore(engine)
	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	require.NoError(t, err)

	fn := instance.GetFunc(store, "test")
	require.NotNil(t, fn)

	result, err := fn.Call(store)
	require.NoError(t, err)
	require.Equal(t, int32(1), result)
}
GOEOF
go test -count=1 .
