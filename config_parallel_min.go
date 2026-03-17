//go:build !wasmtime_parallel_compilation
// +build !wasmtime_parallel_compilation

package wasmtime

// SetParallelCompilation is a no-op without the parallel-compilation feature.
func (cfg *Config) SetParallelCompilation(enabled bool) {}
