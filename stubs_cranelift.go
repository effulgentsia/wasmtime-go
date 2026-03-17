//go:build !wasmtime_cranelift
// +build !wasmtime_cranelift

package wasmtime

/*
#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>

extern void *wasmtime_error_new(const char *msg);

void *wasmtime_module_new(void *engine, const uint8_t *wasm, size_t len,
                          void **out) {
    (void)engine; (void)wasm; (void)len; (void)out;
    return wasmtime_error_new("module compilation unavailable (requires cranelift feature)");
}

void *wasmtime_module_validate(void *engine, const uint8_t *wasm, size_t len) {
    (void)engine; (void)wasm; (void)len;
    return wasmtime_error_new("module validation unavailable (requires cranelift feature)");
}

void *wasmtime_module_serialize(void *module, void *ret) {
    (void)module; (void)ret;
    return wasmtime_error_new("module serialization unavailable (requires cranelift feature)");
}

void wasmtime_config_strategy_set(void *cfg, uint8_t strategy) {
    (void)cfg; (void)strategy;
}

void wasmtime_config_cranelift_debug_verifier_set(void *cfg, bool enable) {
    (void)cfg; (void)enable;
}

void wasmtime_config_cranelift_nan_canonicalization_set(void *cfg, bool enable) {
    (void)cfg; (void)enable;
}

void wasmtime_config_cranelift_opt_level_set(void *cfg, uint8_t level) {
    (void)cfg; (void)level;
}

void wasmtime_config_cranelift_flag_enable(void *cfg, const char *flag) {
    (void)cfg; (void)flag;
}

void wasmtime_config_cranelift_flag_set(void *cfg, const char *flag,
                                        const char *value) {
    (void)cfg; (void)flag; (void)value;
}
*/
import "C"
