//go:build wasmtime_min_build
// +build wasmtime_min_build

package wasmtime

import "fmt"

// SetWasmThreads is a no-op in the headless runtime.
func (cfg *Config) SetWasmThreads(enabled bool) {}

// SetParallelCompilation is a no-op in the headless runtime.
func (cfg *Config) SetParallelCompilation(enabled bool) {}

// SetCraneliftNanCanonicalization is a no-op in the headless runtime.
func (cfg *Config) SetCraneliftNanCanonicalization(enabled bool) {}

// SetStrategy is a no-op in the headless runtime.
func (cfg *Config) SetStrategy(strat Strategy) {}

// SetCraneliftDebugVerifier is a no-op in the headless runtime.
func (cfg *Config) SetCraneliftDebugVerifier(enabled bool) {}

// SetCraneliftOptLevel is a no-op in the headless runtime.
func (cfg *Config) SetCraneliftOptLevel(level OptLevel) {}

// CacheConfigLoadDefault is unavailable in the headless runtime.
func (cfg *Config) CacheConfigLoadDefault() error {
	return fmt.Errorf("cache config unavailable (headless runtime)")
}

// CacheConfigLoad is unavailable in the headless runtime.
func (cfg *Config) CacheConfigLoad(path string) error {
	return fmt.Errorf("cache config unavailable (headless runtime)")
}

// EnableCraneliftFlag is a no-op in the headless runtime.
func (cfg *Config) EnableCraneliftFlag(flag string) {}

// SetCraneliftFlag is a no-op in the headless runtime.
func (cfg *Config) SetCraneliftFlag(name string, value string) {}
