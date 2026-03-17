//go:build (!minimal && !no_threads) || threads
// +build !minimal,!no_threads threads

package wasmtime

// #include <wasmtime.h>
import "C"
import "runtime"

// SetWasmThreads configures whether the wasm threads proposal is enabled
func (cfg *Config) SetWasmThreads(enabled bool) {
	C.wasmtime_config_wasm_threads_set(cfg.ptr(), C.bool(enabled))
	runtime.KeepAlive(cfg)
}
