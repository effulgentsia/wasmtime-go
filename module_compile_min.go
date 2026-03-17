//go:build wasmtime_min_build
// +build wasmtime_min_build

package wasmtime

import "fmt"

var errHeadlessCompile = fmt.Errorf("module compilation unavailable (headless runtime)")

// NewModule is unavailable in the headless runtime. Use NewModuleDeserialize
// to load a pre-compiled module instead.
func NewModule(engine *Engine, wasm []byte) (*Module, error) {
	return nil, errHeadlessCompile
}

// ModuleValidate is unavailable in the headless runtime.
func ModuleValidate(engine *Engine, wasm []byte) error {
	return errHeadlessCompile
}

// Serialize is unavailable in the headless runtime.
func (m *Module) Serialize() ([]byte, error) {
	return nil, errHeadlessCompile
}
