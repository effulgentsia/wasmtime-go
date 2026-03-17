/*
Package wasmtime is a WebAssembly runtime for Go powered by Wasmtime.

This package provides everything necessary to compile and execute WebAssembly
modules as part of a Go program. Wasmtime is a JIT compiler written in Rust,
and can be found at https://github.com/bytecodealliance/wasmtime. This package
is a binding to the C API provided by Wasmtime.

The API of this Go package is intended to mirror the Rust API
(https://docs.rs/wasmtime) relatively closely, so if you find something is
under-documented here then you may have luck consulting the Rust documentation
as well. As always though feel free to file any issues at
https://github.com/bytecodealliance/wasmtime-go/issues/new.

It's also worth pointing out that the authors of this package up to this point
primarily work in Rust, so if you've got suggestions of how to make this package
more idiomatic for Go we'd love to hear your thoughts!

# Feature Build Tags

This package uses Go build tags that correspond to Wasmtime's Cargo features
(see https://github.com/bytecodealliance/wasmtime/blob/main/crates/c-api/Cargo.toml).
When a feature tag is present it indicates the linked Wasmtime C library was
built with that feature. When absent, this package compiles C stubs that
provide the missing symbols so the application links cleanly without needing
to supply its own stubs.

Build with the tags that match the features your Wasmtime C library was built
with. For a full build:

	go build -tags "wasmtime_cranelift,wasmtime_wat,wasmtime_wasi,wasmtime_cache,wasmtime_parallel_compilation,wasmtime_threads"

For a minimal (headless) build with no features, simply omit all tags:

	go build

The available tags and their effects when absent are:

  - wasmtime_cranelift: Without this tag, [NewModule], [NewModuleFromFile],
    [ModuleValidate], and [Module.Serialize] return errors. Use
    [NewModuleDeserialize] or [NewModuleDeserializeFile] with a pre-compiled
    module instead. [Config.SetStrategy], [Config.SetCraneliftOptLevel],
    [Config.SetCraneliftDebugVerifier],
    [Config.SetCraneliftNanCanonicalization], [Config.EnableCraneliftFlag],
    and [Config.SetCraneliftFlag] become no-ops.
  - wasmtime_wat: Without this tag, [Wat2Wasm] returns an error.
  - wasmtime_wasi: Without this tag, [NewWasiConfig] returns a stub config,
    [WasiConfig] setter methods are no-ops or return errors,
    [Linker.DefineWasi] returns an error, and [Store.SetWasi] frees the
    config without configuring WASI.
  - wasmtime_cache: Without this tag, [Config.CacheConfigLoadDefault] and
    [Config.CacheConfigLoad] return errors.
  - wasmtime_parallel_compilation: Without this tag,
    [Config.SetParallelCompilation] is a no-op.
  - wasmtime_threads: Without this tag, [Config.SetWasmThreads] is a no-op.
*/
package wasmtime
