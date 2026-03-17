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
"no_feat_*" build tags to exclude the corresponding Go functions, preventing
compilation errors due to the missing C functions.

For example, to exclude the Go bindings for the cranelift feature:

	go build -tags no_feat_cranelift

To exclude the Go bindings of all optional features (e.g., if linking against
Wasmtime's minimal build):

	go build -tags "no_feat_cranelift,no_feat_wat,no_feat_wasi,no_feat_cache,no_feat_parallel_compilation,no_feat_threads"

The tags correspond to Wasmtime's Cargo features
(see https://github.com/bytecodealliance/wasmtime/blob/main/crates/c-api/Cargo.toml).
*/
package wasmtime
