#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"

rm -f module.cwasm
go test -run '^TestCompile$' .
go test -tags min -run '^TestRun$' .
