//go:build !wasmtime_threads
// +build !wasmtime_threads

package wasmtime

// SetWasmThreads is a no-op without the threads feature.
func (cfg *Config) SetWasmThreads(enabled bool) {}
