package testminimalruntime_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytecodealliance/wasmtime-go/v42"
)

// TestWasmtimeMinAOT_Produce compiles inline WAT and writes a serialized module
// (same artifact shape as wasmtime compile) for TestWasmtimeMinAOT_Load / CI.
func TestWasmtimeMinAOT_Produce(t *testing.T) {
	dir := os.Getenv("WASMTIME_TEST_AOT_DIR")
	if dir == "" {
		dir = "test-wasmtime-min"
	}
	require.NoError(t, os.MkdirAll(dir, 0o755))

	wasm, err := wasmtime.Wat2Wasm(`(module (func (export "test")))`)
	require.NoError(t, err)

	cfg := wasmtime.NewConfig()
	cfg.SetGCSupport(false)
	cfg.SetWasmGC(false)
	cfg.SetWasmThreads(false)
	cfg.SetWasmComponentModel(false)
	engine := wasmtime.NewEngineWithConfig(cfg)
	module, err := wasmtime.NewModule(engine, wasm)
	require.NoError(t, err)
	defer module.Close()

	artifact, err := module.Serialize()
	require.NoError(t, err)

	path := filepath.Join(dir, "module.cwasm")
	require.NoError(t, os.WriteFile(path, artifact, 0o644))
}
