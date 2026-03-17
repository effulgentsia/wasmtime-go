//go:build wasmtime_minimal && !wasmtime_cache
// +build wasmtime_minimal,!wasmtime_cache

package wasmtime

/*
extern void *wasmtime_error_new(const char *msg);

void *wasmtime_config_cache_config_load(void *cfg, const char *path) {
    (void)cfg; (void)path;
    return wasmtime_error_new("cache config unavailable (requires cache feature)");
}
*/
import "C"
