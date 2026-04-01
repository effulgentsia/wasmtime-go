#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
MODULE_PATH=$(awk '/^module /{print $2}' "$REPO_ROOT/go.mod")

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

# Step 2: Vendor the module, then adjust the vendored copy so it compiles
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
for dir in "$VENDOR_PKG"/build/*-min; do
	target="${dir%-min}"
	if [ -d "$target" ]; then
		cp -f "$dir"/* "$target"/
	fi
done

# Step 3: Test that the minimal Wasmtime binary can deserialize and run a module.
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
