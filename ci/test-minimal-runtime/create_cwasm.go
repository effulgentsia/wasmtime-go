//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/bytecodealliance/wasmtime-go/v42"
)

func main() {
	wasm, err := wasmtime.Wat2Wasm(`(module (func (export "test")))`)
	check(err)

	cfg := wasmtime.NewConfig()
	cfg.SetGCSupport(false)
	engine := wasmtime.NewEngineWithConfig(cfg)
	module, err := wasmtime.NewModule(engine, wasm)
	check(err)
	defer module.Close()

	artifact, err := module.Serialize()
	check(err)

	check(os.WriteFile("module.cwasm", artifact, 0o644))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
