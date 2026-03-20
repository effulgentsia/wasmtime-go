# Minimal Wasmtime AOT integration (`test-wasmtime-min`)

## What it does

1. **`TestWasmtimeMinAOT_Produce`** (default build, no `min` tag): compile inline WAT with the full C API, serialize to `test-wasmtime-min/module.cwasm`.
2. **`TestWasmtimeMinAOT_Load`** (`-tags min`): link against the **minimal** static library, deserialize that file, instantiate, call export `test`.

`ci/test-wasmtime-min.sh` runs both steps in order (same pattern as `test-vendoring.sh` in CI).

## v42 limitation (engine / GC)

Wasmtime’s **minimal** static build does not ship GC collector implementations (`gc-drc`, `gc-null`). Creating an `Engine` can still hit Rust-side checks that require a collector when GC-related support is enabled at the engine level.

In **v42** the C API does not expose a dedicated **GC collector / `gc_support` configuration** knob that would let minimal builds align the engine with what the precompiled artifact expects. That is expected to improve in **v43** (see project notes on `gc_support` / collector configuration).

Until then, `go test -tags min -run '^TestWasmtimeMinAOT_Load$'` may **panic during engine creation** even though linking succeeds.

The CI workflow step that runs `./ci/test-wasmtime-min.sh` uses **`continue-on-error: true`** so the rest of CI stays green while this integration is finalized against a newer Wasmtime release.

## Local use

```bash
python3 ci/download-wasmtime.py
./ci/test-wasmtime-min.sh
```

Artifact directory: `test-wasmtime-min/` (gitignored). Override with `WASMTIME_TEST_AOT_DIR`.
