//go:build wasmtime_minimal && !wasmtime_wat
// +build wasmtime_minimal,!wasmtime_wat

package wasmtime

/*
#include <stddef.h>

extern void *wasmtime_error_new(const char *msg);

void *wasmtime_wat2wasm(const char *wat, size_t wat_len, void *ret) {
    (void)wat; (void)wat_len; (void)ret;
    return wasmtime_error_new("wat2wasm unavailable (requires wat feature)");
}
*/
import "C"
