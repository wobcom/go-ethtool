package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wobcom/go-ethtool/eeprom"
)

const (
	// IFNAMSIZ is the maximum length of an interface name
	IFNAMSIZ = 16
)

// Interface representation of a network interface
type Interface struct {
	Name      string
	nameBytes [IFNAMSIZ]byte

	Eeprom     eeprom.EEPROM
	DriverInfo *DriverInfo

	ethtool *Ethtool
}

// NewInterface retrieves information about a network interface and returns an Interface instance
func (e *Ethtool) NewInterface(ifname string, ignoreEepromReadErrors bool) (*Interface, error) {
	if len([]byte(ifname)) >= IFNAMSIZ {
		return nil, errors.New(fmt.Sprintf("Interface name must not be longer than %d characters.", IFNAMSIZ))
	}
	iface := &Interface{
		ethtool: e,
		Name:    ifname,
	}

	driverInfo, err := iface.getDriverInfo()
	if err != nil {
		return iface, errors.Wrapf(err, "Could not retrieve driver info")
	}
	iface.DriverInfo = driverInfo

	eeprom, err := iface.getEEPROM()
	if err != nil {
		if ignoreEepromReadErrors {
			return iface, nil
		}
		return iface, err
	}
	iface.Eeprom = eeprom
	return iface, nil
}

type ifreq struct {
	ifrName [IFNAMSIZ]byte
	ifrData uintptr
}

func (i *Interface) performIoctl(data uintptr) error {
	var name [IFNAMSIZ]byte

	copy(name[:], []byte(i.Name))
	ifr := ifreq{
		ifrName: name,
		ifrData: data,
	}

	return i.ethtool.PerformIoctl(&ifr)
}

type ethtoolArbitraryCommand struct {
	cmd   uint32
	value uint32
}
