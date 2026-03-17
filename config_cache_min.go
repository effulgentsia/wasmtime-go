//go:build !wasmtime_cache
// +build !wasmtime_cache

package wasmtime

import "fmt"

// CacheConfigLoadDefault is unavailable without the cache feature.
func (cfg *Config) CacheConfigLoadDefault() error {
	return fmt.Errorf("cache config unavailable (requires cache feature)")
}

// CacheConfigLoad is unavailable without the cache feature.
func (cfg *Config) CacheConfigLoad(path string) error {
	return fmt.Errorf("cache config unavailable (requires cache feature)")
}
