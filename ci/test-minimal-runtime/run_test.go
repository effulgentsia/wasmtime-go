//go:build min

package testminimalruntime_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytecodealliance/wasmtime-go/v42"
)

func TestRun(t *testing.T) {
	cfg := wasmtime.NewConfig()
	cfg.SetGCSupport(false)
	engine := wasmtime.NewEngineWithConfig(cfg)
	module, err := wasmtime.NewModuleDeserializeFile(engine, "module.cwasm")
	require.NoError(t, err)
	defer module.Close()

	store := wasmtime.NewStore(engine)
	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	require.NoError(t, err)

	fn := instance.GetFunc(store, "test")
	require.NotNil(t, fn)

	_, err = fn.Call(store)
	require.NoError(t, err)
}
