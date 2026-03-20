//go:build !min

package testminimalruntime_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytecodealliance/wasmtime-go/v42"
)

func TestWasmtimeMinAOT_Produce(t *testing.T) {
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

	require.NoError(t, os.WriteFile("module.cwasm", artifact, 0o644))
}
