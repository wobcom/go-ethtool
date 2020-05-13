package ethtool

import (
	"unsafe"
)

const (
	// Reset hardware
	resetIoctl = 0x00000034
)

type ethtoolReset struct {
	cmd   uint32
	flags uint32
}

// PerformReset performs an interface reset
func (i *Interface) PerformReset() error {
	reset := ethtoolReset{
		cmd:   resetIoctl,
		flags: 0xFFFFFFFF,
	}

	return i.performIoctl(uintptr(unsafe.Pointer(&reset)))
}
