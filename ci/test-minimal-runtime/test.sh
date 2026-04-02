#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
MODULE_PATH=$(awk '/^module /{print $2}' "$REPO_ROOT/go.mod")
MAJOR_VERSION=$(echo "$MODULE_PATH" | grep -oP '\d+$' || echo "0")
PLACEHOLDER="v${MAJOR_VERSION}.0.0-placeholder"
WASMTIME_VERSION=$(grep -oP "^version\s*=\s*'\K[^']+" "$REPO_ROOT/ci/download-wasmtime.py")

WORK_DIR=$(mktemp -d)
trap 'rm -rf "$WORK_DIR"' EXIT
cd "$WORK_DIR"

# Copy the static Go source files.
cp "$SCRIPT_DIR/create_cwasm.go" .
cp "$SCRIPT_DIR/min_test.go" .

go mod init test-minimal-runtime
cat >> go.mod <<EOF
require "$MODULE_PATH" ${PLACEHOLDER}
replace "$MODULE_PATH" ${PLACEHOLDER} => "$REPO_ROOT"
EOF
go mod tidy

# Step 1: Create a pre-compiled module using the full Wasmtime library.
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
go test -count=1 .
