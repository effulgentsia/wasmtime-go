//go:build wasmtime_min_build
// +build wasmtime_min_build

package wasmtime

import "fmt"

// Wat2Wasm is unavailable in the headless runtime because the WAT parser
// requires the compiler feature.
func Wat2Wasm(wat string) ([]byte, error) {
	return nil, fmt.Errorf("wat2wasm unavailable (headless runtime)")
}
