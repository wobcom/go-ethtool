package ethtool

import (
	"unsafe"
)

const (
	ETHTOOL_RESET = 0x00000034 /* Reset hardware */
)

type ethtoolReset struct {
	cmd   uint32
	flags uint32
}

// Interestingly this ioctl seems to be not implemented by any driver :/
func (i *Interface) PerformReset() error {
	reset := ethtoolReset{
		cmd:   ETHTOOL_RESET,
		flags: 0xFFFFFFFF,
	}

	return i.performIoctl(uintptr(unsafe.Pointer(&reset)))
}
