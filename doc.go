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

By default this package expects a full Wasmtime C library with all features.
A plain "go build" links against the full library and all functions are
available.

To link against a Wasmtime C library built without some or all features, use
no_feat_ build tags to exclude the corresponding Go functions. The functions
are simply not compiled, so any accidental use becomes a compile-time error.

For example, to disable the cranelift feature:

	go build -tags no_feat_cranelift

To build with no optional features (minimal/headless):

	go build -tags "no_feat_cranelift,no_feat_wat,no_feat_wasi,no_feat_cache,no_feat_parallel_compilation,no_feat_threads"

The tags correspond to Wasmtime's Cargo features
(see https://github.com/bytecodealliance/wasmtime/blob/main/crates/c-api/Cargo.toml).
When a feature is excluded, the following functions are not available:

  - no_feat_cranelift: [NewModule], [ModuleValidate], [Module.Serialize],
    [Config.SetStrategy], [Config.SetCraneliftOptLevel],
    [Config.SetCraneliftDebugVerifier],
    [Config.SetCraneliftNanCanonicalization], [Config.EnableCraneliftFlag],
    [Config.SetCraneliftFlag]. Use [NewModuleDeserialize] or
    [NewModuleDeserializeFile] with a pre-compiled module instead.
  - no_feat_wat: [Wat2Wasm]. Note: [NewModuleFromFile] requires both
    cranelift and wat.
  - no_feat_wasi: [WasiConfig] and all its methods, [Linker.DefineWasi],
    [Store.SetWasi].
  - no_feat_cache: [Config.CacheConfigLoadDefault], [Config.CacheConfigLoad].
  - no_feat_parallel_compilation: [Config.SetParallelCompilation].
  - no_feat_threads: [Config.SetWasmThreads].
*/
package wasmtime
