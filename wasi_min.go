//go:build wasmtime_min_build
// +build wasmtime_min_build

package wasmtime

import "fmt"

// WasiConfig represents WASI configuration. In the headless runtime all
// methods are stubs that return errors or are no-ops.
type WasiConfig struct {
	_closed bool
}

// NewWasiConfig creates a new WasiConfig. In the headless runtime the
// returned object is a lightweight stub.
func NewWasiConfig() *WasiConfig {
	return &WasiConfig{}
}

func (c *WasiConfig) ptr() uintptr {
	if c._closed {
		panic("object has been closed already")
	}
	return 0
}

// Close is a no-op in the headless runtime.
func (c *WasiConfig) Close() {
	c._closed = true
}

// SetArgv is a no-op in the headless runtime.
func (c *WasiConfig) SetArgv(argv []string) {}

// InheritArgv is a no-op in the headless runtime.
func (c *WasiConfig) InheritArgv() {}

// SetEnv is a no-op in the headless runtime.
func (c *WasiConfig) SetEnv(keys, values []string) {}

// InheritEnv is a no-op in the headless runtime.
func (c *WasiConfig) InheritEnv() {}

// SetStdinFile is unavailable in the headless runtime.
func (c *WasiConfig) SetStdinFile(path string) error {
	return fmt.Errorf("WASI unavailable (headless runtime)")
}

// InheritStdin is a no-op in the headless runtime.
func (c *WasiConfig) InheritStdin() {}

// SetStdoutFile is unavailable in the headless runtime.
func (c *WasiConfig) SetStdoutFile(path string) error {
	return fmt.Errorf("WASI unavailable (headless runtime)")
}

// InheritStdout is a no-op in the headless runtime.
func (c *WasiConfig) InheritStdout() {}

// SetStderrFile is unavailable in the headless runtime.
func (c *WasiConfig) SetStderrFile(path string) error {
	return fmt.Errorf("WASI unavailable (headless runtime)")
}

// InheritStderr is a no-op in the headless runtime.
func (c *WasiConfig) InheritStderr() {}

// WasiDirPerms represents directory permissions for WASI.
type WasiDirPerms uint8

// WasiFilePerms represents file permissions for WASI.
type WasiFilePerms uint8

const (
	DIR_READ   WasiDirPerms  = 0x1
	DIR_WRITE  WasiDirPerms  = 0x2
	FILE_READ  WasiFilePerms = 0x1
	FILE_WRITE WasiFilePerms = 0x2
)

// PreopenDir is unavailable in the headless runtime.
func (c *WasiConfig) PreopenDir(path, guestPath string, dirPerms WasiDirPerms, filePerms WasiFilePerms) error {
	return fmt.Errorf("WASI unavailable (headless runtime)")
}
