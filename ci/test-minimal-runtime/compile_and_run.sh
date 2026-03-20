#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/../.."

rm -f ci/test-minimal-runtime/module.cwasm
go test -run '^TestWasmtimeMinAOT_Produce$' ./ci/test-minimal-runtime/
go test -tags min -run '^TestWasmtimeMinAOT_Load$' ./ci/test-minimal-runtime/
