//go:build wasmtime_min_build
// +build wasmtime_min_build

package wasmtime

import "fmt"

// DefineWasi is unavailable in the headless runtime.
func (l *Linker) DefineWasi() error {
	return fmt.Errorf("WASI unavailable (headless runtime)")
}
