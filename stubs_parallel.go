//go:build (minimal || no_parallel_compilation) && !parallel_compilation
// +build !parallel_compilation
// +build minimal no_parallel_compilation

package wasmtime

/*
#include <stdbool.h>

void wasmtime_config_parallel_compilation_set(void *cfg, bool enable) {
    (void)cfg; (void)enable;
}
*/
import "C"
