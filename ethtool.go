package ethtool

import (
	"golang.org/x/sys/unix"
	"sync"
	"unsafe"
)

const (
	siocethtool = 0x8946
)

// Ethtool provides a wrapper around the Kernel's ethtool ioctls
type Ethtool struct {
	fd int
	mu *sync.Mutex
}

// NewEthtool initializes internal data structure (i.e. opens a socket) and returns a new Ethtool instance
func NewEthtool() (*Ethtool, error) {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_IP)

	if err != nil {
		return nil, err
	}

	return &Ethtool{
		fd: fd,
		mu: &sync.Mutex{},
	}, nil
}

// PerformIoctl performs an ethtool ioctl and passes the given pointer to the ioctl
func (e *Ethtool) PerformIoctl(ifr *ifreq) error {
	e.mu.Lock()
	_, _, ep := unix.Syscall(unix.SYS_IOCTL, uintptr(e.fd), siocethtool, uintptr(unsafe.Pointer(ifr)))
	e.mu.Unlock()

	if ep != 0 {
		return ep
	}

	return nil
}

// Close closes the internally used socket
func (e *Ethtool) Close() {
	unix.Close(e.fd)
}
