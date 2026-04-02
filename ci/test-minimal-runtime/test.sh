#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
MODULE_PATH=$(awk '/^module /{print $2}' "$REPO_ROOT/go.mod")
MAJOR_VERSION=$(echo "$MODULE_PATH" | grep -oP '\d+$' || echo "0")
PLACEHOLDER="v${MAJOR_VERSION}.0.0-placeholder"

cd "$SCRIPT_DIR"
trap 'rm -rf go.mod go.sum vendor build module.cwasm' EXIT

go mod init test-minimal-runtime
cat >> go.mod <<EOF
require "$MODULE_PATH" ${PLACEHOLDER}
replace "$MODULE_PATH" ${PLACEHOLDER} => "$REPO_ROOT"
EOF
go mod tidy

# Step 1: Create a pre-compiled module using the full Wasmtime library.
go run create_cwasm.go

# Step 2: Download the minimal Wasmtime static libraries.
python3 "$REPO_ROOT/ci/download-wasmtime.py" --min

# Step 3: Vendor the module, then adjust the vendored copy so it compiles
# against the minimal Wasmtime library:
#   a) Remove Go source files that call C functions absent from the min binary.
#   b) Copy the min static libraries over the full ones so CGO links the right lib.
go mod vendor
VENDOR_PKG="vendor/${MODULE_PATH}"
rm -f "$VENDOR_PKG"/wat2wasm.go
rm -f "$VENDOR_PKG"/wasi.go
rm -f "$VENDOR_PKG"/config_feat_*.go
rm -f "$VENDOR_PKG"/module_feat_*.go
rm -f "$VENDOR_PKG"/module_feats_*.go
rm -f "$VENDOR_PKG"/linker_feat_*.go
rm -f "$VENDOR_PKG"/store_feat_*.go
cp -rf build/* "$VENDOR_PKG/build/"

# Step 4: Test that the minimal Wasmtime binary can deserialize and run a module.
go test -count=1 .
