//go:build (minimal || no_threads) && !threads
// +build !threads
// +build minimal no_threads

package wasmtime

/*
#include <stdbool.h>

void wasmtime_config_wasm_threads_set(void *cfg, bool enable) {
    (void)cfg; (void)enable;
}
*/
import "C"
