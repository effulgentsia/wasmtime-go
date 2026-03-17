//go:build wasmtime_min_build
// +build wasmtime_min_build

package wasmtime

import "runtime"

// SetWasi is a no-op in the headless runtime. The WasiConfig is consumed
// (closed) but no WASI state is actually configured.
func (store *Store) SetWasi(wasi *WasiConfig) {
	runtime.SetFinalizer(wasi, nil)
	wasi._closed = true
	runtime.KeepAlive(store)
}
