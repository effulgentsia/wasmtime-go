#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"

rm -f module.cwasm
go run create_cwasm.go
go test -tags min -run '^TestRun$' .
