//go:build wasmtime_minimal && !wasmtime_parallel_compilation
// +build wasmtime_minimal,!wasmtime_parallel_compilation

package wasmtime

/*
#include <stdbool.h>

void wasmtime_config_parallel_compilation_set(void *cfg, bool enable) {
    (void)cfg; (void)enable;
}
*/
import "C"
