#!/usr/bin/env bash
# Two-step integration test: full Wasmtime produces a precompiled .cwasm artifact;
# minimal Wasmtime (Go build tag `min`) loads and runs it.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

rm -rf test-wasmtime-min
mkdir -p test-wasmtime-min
export WASMTIME_TEST_AOT_DIR="$ROOT/test-wasmtime-min"

go test -run '^TestWasmtimeMinAOT_Produce$' ./aotproduce/
go test -tags min -run '^TestWasmtimeMinAOT_Load$' ./minload/

echo "test-wasmtime-min: success"
