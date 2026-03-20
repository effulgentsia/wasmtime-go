#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"

rm -f module.cwasm
go test -run '^TestWasmtimeMinAOT_Produce$' .
go test -tags min -run '^TestWasmtimeMinAOT_Load$' .
