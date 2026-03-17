//go:build !wasmtime_cranelift
// +build !wasmtime_cranelift

package wasmtime

// SetStrategy is a no-op without the cranelift feature.
func (cfg *Config) SetStrategy(strat Strategy) {}

// SetCraneliftDebugVerifier is a no-op without the cranelift feature.
func (cfg *Config) SetCraneliftDebugVerifier(enabled bool) {}

// SetCraneliftOptLevel is a no-op without the cranelift feature.
func (cfg *Config) SetCraneliftOptLevel(level OptLevel) {}

// SetCraneliftNanCanonicalization is a no-op without the cranelift feature.
func (cfg *Config) SetCraneliftNanCanonicalization(enabled bool) {}

// EnableCraneliftFlag is a no-op without the cranelift feature.
func (cfg *Config) EnableCraneliftFlag(flag string) {}

// SetCraneliftFlag is a no-op without the cranelift feature.
func (cfg *Config) SetCraneliftFlag(name string, value string) {}
