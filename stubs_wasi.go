//go:build wasmtime_minimal && !wasmtime_wasi
// +build wasmtime_minimal,!wasmtime_wasi

package wasmtime

/*
#include <stddef.h>
#include <stdbool.h>
#include <stdlib.h>

extern void *wasmtime_error_new(const char *msg);

// Minimal opaque struct so that new/delete round-trip without crashing.
typedef struct { int dummy; } wasi_config_stub_t;

void *wasi_config_new(void) {
    return calloc(1, sizeof(wasi_config_stub_t));
}

void wasi_config_delete(void *cfg) {
    free(cfg);
}

bool wasi_config_set_argv(void *cfg, size_t argc, const char *argv[]) {
    (void)cfg; (void)argc; (void)argv;
    return true;
}

void wasi_config_inherit_argv(void *cfg) { (void)cfg; }

bool wasi_config_set_env(void *cfg, size_t envc, const char *names[],
                         const char *values[]) {
    (void)cfg; (void)envc; (void)names; (void)values;
    return true;
}

void wasi_config_inherit_env(void *cfg) { (void)cfg; }

bool wasi_config_set_stdin_file(void *cfg, const char *path) {
    (void)cfg; (void)path;
    return false;
}

void wasi_config_inherit_stdin(void *cfg) { (void)cfg; }

bool wasi_config_set_stdout_file(void *cfg, const char *path) {
    (void)cfg; (void)path;
    return false;
}

void wasi_config_inherit_stdout(void *cfg) { (void)cfg; }

bool wasi_config_set_stderr_file(void *cfg, const char *path) {
    (void)cfg; (void)path;
    return false;
}

void wasi_config_inherit_stderr(void *cfg) { (void)cfg; }

bool wasi_config_preopen_dir(void *cfg, const char *host_path,
                             const char *guest_path, size_t dir_perms,
                             size_t file_perms) {
    (void)cfg; (void)host_path; (void)guest_path;
    (void)dir_perms; (void)file_perms;
    return false;
}

void *wasmtime_linker_define_wasi(void *linker) {
    (void)linker;
    return wasmtime_error_new("WASI unavailable (requires wasi feature)");
}

void wasmtime_context_set_wasi(void *context, void *wasi) {
    (void)context;
    free(wasi);
}
*/
import "C"
