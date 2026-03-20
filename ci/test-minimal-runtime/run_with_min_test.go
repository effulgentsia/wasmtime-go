//go:build min

// External tests for package wasmtime. The import path matches go.mod’s module
// path; it refers to this repository’s root package, not a remote version.
// Running `go test -tags min ./ci/test-minimal-runtime/` compiles github.com/.../wasmtime-go/v42
// (the parent module) with the `min` build tag, so `//go:build !min` sources are
// omitted and the link matches the minimal C library.

package testminimalruntime_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytecodealliance/wasmtime-go/v42"
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
		t.Skipf("missing %s (run TestWasmtimeMinAOT_Produce or ci/test-minimal-runtime/compile_and_run.sh first): %v", path, err)
	}

	cfg := wasmtime.NewConfig()
	cfg.SetGCSupport(false)
	cfg.SetWasmGC(false)
	engine := wasmtime.NewEngineWithConfig(cfg)
	module, err := wasmtime.NewModuleDeserializeFile(engine, path)
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
