//go:build min

package wasmtime

// Integration test: deserialize a .cwasm produced by TestWasmtimeMinAOT_Produce.

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestWasmtimeMinAOT_Load deserializes module.cwasm produced by
// TestWasmtimeMinAOT_Produce and invokes export "test".
func TestWasmtimeMinAOT_Load(t *testing.T) {
	dir := os.Getenv("WASMTIME_TEST_AOT_DIR")
	if dir == "" {
		dir = "test-wasmtime-min"
	}
	path := filepath.Join(dir, "module.cwasm")
	if _, err := os.Stat(path); err != nil {
		t.Skipf("missing %s (run TestWasmtimeMinAOT_Produce or ci/test-wasmtime-min.sh first): %v", path, err)
	}

	cfg := NewConfig()
	cfg.SetGCSupport(false)
	cfg.SetWasmGC(false)
	engine := NewEngineWithConfig(cfg)
	module, err := NewModuleDeserializeFile(engine, path)
	require.NoError(t, err)
	defer module.Close()

	store := NewStore(engine)
	instance, err := NewInstance(store, module, []AsExtern{})
	require.NoError(t, err)

	fn := instance.GetFunc(store, "test")
	require.NotNil(t, fn)

	_, err = fn.Call(store)
	require.NoError(t, err)
}
