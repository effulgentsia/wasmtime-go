//go:build wasmtime_minimal && !wasmtime_threads
// +build wasmtime_minimal,!wasmtime_threads

package wasmtime

/*
#include <stdbool.h>

void wasmtime_config_wasm_threads_set(void *cfg, bool enable) {
    (void)cfg; (void)enable;
}
*/
import "C"
