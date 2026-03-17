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

# Headless / Minimal Build

This package supports linking against a minimal (headless) Wasmtime library
that omits the compiler and WASI subsystems. To use this mode, build with the
wasmtime_min_build tag:

	go build -tags wasmtime_min_build

In this mode the following restrictions apply:

  - Module compilation is unavailable: [NewModule], [NewModuleFromFile], and
    [ModuleValidate] return errors. Use [NewModuleDeserialize] or
    [NewModuleDeserializeFile] with a pre-compiled module instead.
  - [Module.Serialize] returns an error.
  - [Wat2Wasm] returns an error.
  - Compiler configuration methods ([Config.SetStrategy],
    [Config.SetCraneliftOptLevel], [Config.SetCraneliftDebugVerifier],
    [Config.SetCraneliftNanCanonicalization], [Config.SetParallelCompilation],
    [Config.EnableCraneliftFlag], [Config.SetCraneliftFlag],
    [Config.SetWasmThreads]) are no-ops.
  - [Config.CacheConfigLoadDefault] and [Config.CacheConfigLoad] return errors.
  - All WASI functions ([NewWasiConfig], [WasiConfig] methods,
    [Linker.DefineWasi], [Store.SetWasi]) are stubs that either no-op or return
    errors.
*/
package wasmtime
