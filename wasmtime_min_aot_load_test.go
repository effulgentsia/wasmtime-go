//go:build min

package wasmtime

// Integration test: deserialize a .cwasm produced by TestWasmtimeMinAOT_Produce.
//
// Wasmtime v42 minimal builds can panic when creating an Engine (GC collector
// availability). v43 is expected to add the C-side hooks needed to configure
// this; until then this test documents intent and may fail at runtime even when
// linking succeeds. See ci/wasmtime-min-aot.md.

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
