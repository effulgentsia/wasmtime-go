# Minimal Wasmtime AOT integration (`test-wasmtime-min`)

## What it does

1. **`aotproduce.TestWasmtimeMinAOT_Produce`** (default build, no `min` tag): compile inline WAT with the full C API, serialize to `test-wasmtime-min/module.cwasm`.
2. **`minload.TestWasmtimeMinAOT_Load`** (`-tags min`): link against the **minimal** static library, deserialize that file, instantiate, call export `test`.

Tests live in separate packages (`aotproduce/`, `minload/`) so `go test -tags min ./minload/` only builds that test—**not** the main package’s tests (which would fail under `min` because they call APIs omitted by the `min` build tag).

`ci/test-wasmtime-min.sh` runs both steps in order (same pattern as `test-vendoring.sh` in CI).

The load path uses **`Config.SetGCSupport(false)`** so the engine does not require GC collectors shipped only in full Wasmtime builds.

The **produce** step configures the full compiler so the serialized artifact matches what the minimal runtime exposes: **`SetWasmThreads(false)`** and **`SetWasmComponentModel(false)`** (among **`SetGCSupport` / `SetWasmGC`**) so deserialization does not require threads or the component model, which the minimal build omits.

## Local use

```bash
python3 ci/download-wasmtime.py
./ci/test-wasmtime-min.sh
```

Artifact directory: `test-wasmtime-min/` (gitignored). Override with `WASMTIME_TEST_AOT_DIR`.

### Manual commands

```bash
go test -run '^TestWasmtimeMinAOT_Produce$' ./aotproduce/
go test -tags min -run '^TestWasmtimeMinAOT_Load$' ./minload/
```

Do **not** run `go test -tags min .` at the repo root: that compiles **all** `wasmtime` package tests together, which still require the full API.
