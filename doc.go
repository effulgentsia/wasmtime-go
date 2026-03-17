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

# Minimal / Feature-Reduced Builds

By default this package expects a full Wasmtime C library with all features.
A plain "go build" links against the full library with no stubs.

To link against a Wasmtime C library built without all features, use the
"minimal" build tag. This compiles C stubs for every absent feature so the
application links without needing to provide its own stubs:

	go build -tags minimal

Individual features can be added back with their corresponding tags. This is
analogous to Cargo's --no-default-features --features <list>. The feature
tags match Wasmtime's Cargo features
(see https://github.com/bytecodealliance/wasmtime/blob/main/crates/c-api/Cargo.toml):

	go build -tags "minimal,cranelift"

Individual features can also be disabled without using the "minimal" tag:

	go build -tags no_cranelift

The available feature tags and what their absence stubs out are:

  - cranelift: [NewModule], [NewModuleFromFile], [ModuleValidate], and
    [Module.Serialize] return errors. Use [NewModuleDeserialize] or
    [NewModuleDeserializeFile] with a pre-compiled module instead.
    [Config.SetStrategy], [Config.SetCraneliftOptLevel],
    [Config.SetCraneliftDebugVerifier],
    [Config.SetCraneliftNanCanonicalization], [Config.EnableCraneliftFlag],
    and [Config.SetCraneliftFlag] become no-ops.
  - wat: [Wat2Wasm] returns an error.
  - wasi: [NewWasiConfig] returns a stub config, [WasiConfig] setter
    methods are no-ops or return errors, [Linker.DefineWasi] returns an
    error, and [Store.SetWasi] frees the config without configuring WASI.
  - cache: [Config.CacheConfigLoadDefault] and [Config.CacheConfigLoad]
    return errors.
  - parallel_compilation: [Config.SetParallelCompilation] is a no-op.
  - threads: [Config.SetWasmThreads] is a no-op.
*/
package wasmtime
