//go:build !wasmtime_wasi
// +build !wasmtime_wasi

package wasmtime

import "runtime"

// SetWasi is a no-op without the wasi feature. The WasiConfig is consumed
// (closed) but no WASI state is actually configured.
func (store *Store) SetWasi(wasi *WasiConfig) {
	runtime.SetFinalizer(wasi, nil)
	wasi._closed = true
	runtime.KeepAlive(store)
}
