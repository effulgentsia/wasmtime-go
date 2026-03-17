//go:build !wasmtime_wat
// +build !wasmtime_wat

package wasmtime

import "fmt"

// Wat2Wasm is unavailable without the wat feature.
func Wat2Wasm(wat string) ([]byte, error) {
	return nil, fmt.Errorf("wat2wasm unavailable (requires wat feature)")
}
