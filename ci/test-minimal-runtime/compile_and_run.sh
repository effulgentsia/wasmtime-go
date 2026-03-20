#!/usr/bin/env bash
# Two-step integration test: full Wasmtime produces a precompiled .cwasm artifact;
# minimal Wasmtime (Go build tag `min`) loads and runs it.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$ROOT"

rm -f "$SCRIPT_DIR/module.cwasm"
export WASMTIME_TEST_AOT_DIR="$SCRIPT_DIR"

go test -run '^TestWasmtimeMinAOT_Produce$' ./ci/test-minimal-runtime/
go test -tags min -run '^TestWasmtimeMinAOT_Load$' ./ci/test-minimal-runtime/

echo "test-minimal-runtime: success"
