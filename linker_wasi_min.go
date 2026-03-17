//go:build !wasmtime_wasi
// +build !wasmtime_wasi

package wasmtime

import "fmt"

// DefineWasi is unavailable without the wasi feature.
func (l *Linker) DefineWasi() error {
	return fmt.Errorf("WASI unavailable (requires wasi feature)")
}
