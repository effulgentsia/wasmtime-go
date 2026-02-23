#include "_cgo_export.h"
#include "shims.h"
#if !defined(_WIN32)
#include <dlfcn.h>
#endif

wasmtime_store_t *go_store_new(wasm_engine_t *engine, size_t env) {
  return wasmtime_store_new(engine, (void*) env, goFinalizeStore);
}

static wasm_trap_t* trampoline(
   void *env,
   wasmtime_caller_t *caller,
   const wasmtime_val_t *args,
   size_t nargs,
   wasmtime_val_t *results,
   size_t nresults
) {
    return goTrampolineNew(caller, (size_t) env,
        (wasmtime_val_t*) args, nargs,
        results, nresults);
}

static wasm_trap_t* wrap_trampoline(
   void *env,
   wasmtime_caller_t *caller,
   const wasmtime_val_t *args,
   size_t nargs,
   wasmtime_val_t *results,
   size_t nresults
) {
    return goTrampolineWrap(caller, (size_t) env,
        (wasmtime_val_t*) args, nargs,
        results, nresults);
}

void go_func_new(
    wasmtime_context_t *store,
    wasm_functype_t *ty,
    size_t env,
    int wrap,
    wasmtime_func_t *ret
) {
  wasmtime_func_callback_t callback = trampoline;
  if (wrap)
    callback = wrap_trampoline;
  return wasmtime_func_new(store, ty, callback, (void*) env, NULL, ret);
}

wasmtime_error_t *go_linker_define_func(
    wasmtime_linker_t *linker,
    const char *module,
    size_t module_len,
    const char *name,
    size_t name_len,
    const wasm_functype_t *ty,
    int wrap,
    size_t env
) {
  wasmtime_func_callback_t cb = trampoline;
  void(*finalizer)(void*) = goFinalizeFuncNew;
  if (wrap) {
    cb = wrap_trampoline;
    finalizer = goFinalizeFuncWrap;
  }
  return wasmtime_linker_define_func(linker, module, module_len, name, name_len, ty, cb, (void*) env, finalizer);
}

bool go_externref_new(wasmtime_context_t *cx, size_t env, wasmtime_externref_t *ref) {
  return wasmtime_externref_new(cx, (void*) env, goFinalizeExternref, ref);
}

typedef void (*go_gc_support_set_t)(wasm_config_t *, bool);

void go_wasmtime_config_gc_support_set(wasm_config_t *config, bool enable) {
#if !defined(_WIN32)
  static int looked_up = 0;
  static go_gc_support_set_t fn = NULL;
  if (!looked_up) {
    looked_up = 1;
    fn = (go_gc_support_set_t) dlsym(RTLD_DEFAULT, "wasmtime_config_gc_support_set");
  }
  if (fn != NULL) {
    fn(config, enable);
    return;
  }
#endif
  wasmtime_config_wasm_gc_set(config, enable);
}

#define UNION_ACCESSOR(name, field, ty) \
  ty go_##name##_##field##_get(const name##_t *val) { return val->of.field; } \
  void go_##name##_##field##_set(name##_t *val, ty i) { val->of.field = i; }

EACH_UNION_ACCESSOR(UNION_ACCESSOR)
