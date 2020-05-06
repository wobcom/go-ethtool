package ethtool

import (
	"sync"

	"golang.org/x/sys/unix"
)

const (
	SIOCETHTOOL = 0x8946
)

type Ethtool struct {
	fd int
	mu *sync.Mutex
}

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

func (e *Ethtool) PerformIoctl(ifr uintptr) error {
	// TODO figure out if locking is necessary
	//    e.mu.Lock()
	_, _, ep := unix.Syscall(unix.SYS_IOCTL, uintptr(e.fd), SIOCETHTOOL, ifr)
	//    e.mu.Unlock()

	if ep != 0 {
		return ep
	}

	return nil
}

func (e *Ethtool) Close() {
	unix.Close(e.fd)
}
