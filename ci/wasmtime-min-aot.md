# Minimal Wasmtime AOT integration (`test-wasmtime-min`)

## What it does

1. **`TestWasmtimeMinAOT_Produce`** (default build, no `min` tag): compile inline WAT with the full C API, serialize to `test-wasmtime-min/module.cwasm`.
2. **`TestWasmtimeMinAOT_Load`** (`-tags min`): link against the **minimal** static library, deserialize that file, instantiate, call export `test`.

`ci/test-wasmtime-min.sh` runs both steps in order (same pattern as `test-vendoring.sh` in CI).

The load path uses **`Config.SetGCSupport(false)`** so the engine does not require GC collectors shipped only in full Wasmtime builds.

The **produce** step configures the full compiler so the serialized artifact matches what the minimal runtime exposes: **`SetWasmThreads(false)`** and **`SetWasmComponentModel(false)`** (among **`SetGCSupport` / `SetWasmGC`**) so deserialization does not require threads or the component model, which the minimal build omits.

## Local use

```bash
python3 ci/download-wasmtime.py
./ci/test-wasmtime-min.sh
```

Artifact directory: `test-wasmtime-min/` (gitignored). Override with `WASMTIME_TEST_AOT_DIR`.
