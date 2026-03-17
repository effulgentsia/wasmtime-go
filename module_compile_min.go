//go:build !wasmtime_cranelift
// +build !wasmtime_cranelift

package wasmtime

import "fmt"

var errNoCranelift = fmt.Errorf("module compilation unavailable (requires cranelift feature)")

// NewModule is unavailable without the cranelift feature. Use
// NewModuleDeserialize to load a pre-compiled module instead.
func NewModule(engine *Engine, wasm []byte) (*Module, error) {
	return nil, errNoCranelift
}

// ModuleValidate is unavailable without the cranelift feature.
func ModuleValidate(engine *Engine, wasm []byte) error {
	return errNoCranelift
}

// Serialize is unavailable without the cranelift feature.
func (m *Module) Serialize() ([]byte, error) {
	return nil, errNoCranelift
}
