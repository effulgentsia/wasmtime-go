#!/usr/bin/env bash
# Two-step integration test: full Wasmtime produces a precompiled .cwasm artifact;
# minimal Wasmtime (Go build tag `min`) loads and runs it.
#
# Known issue (Wasmtime v42): minimal static libs lack GC collectors; engine
# creation for the load step may panic until the C API exposes the needed
# gc_support / collector configuration (anticipated in v43). See ci/wasmtime-min-aot.md.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

rm -rf test-wasmtime-min
mkdir -p test-wasmtime-min
export WASMTIME_TEST_AOT_DIR="$ROOT/test-wasmtime-min"

go test -run '^TestWasmtimeMinAOT_Produce$' .
go test -tags min -run '^TestWasmtimeMinAOT_Load$' .

echo "test-wasmtime-min: success"
