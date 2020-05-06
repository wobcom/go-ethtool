package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/wobcom/ethtool/eeprom"
	"unsafe"
)

const (
	IFNAMSIZ = 16
)

type Interface struct {
	Name      string
	nameBytes [IFNAMSIZ]byte

	Eeprom     eeprom.EEPROM
	DriverInfo *DriverInfo

	ethtool *Ethtool
}

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
	ifr_name [IFNAMSIZ]byte
	ifr_data uintptr
}

func (i *Interface) performIoctl(data uintptr) error {
	var name [IFNAMSIZ]byte

	copy(name[:], []byte(i.Name))
	ifr := ifreq{
		ifr_name: name,
		ifr_data: data,
	}

	return i.ethtool.PerformIoctl(uintptr(unsafe.Pointer(&ifr)))
}

type ethtoolArbitraryCommand struct {
	cmd   uint32
	value uint32
}

const ETHTOOL_PHYS_ID = 0x0000001c /* identify the NIC */

// Makes the interface LED blink
func (i *Interface) Identify(time uint32) error {
	cmd := ethtoolArbitraryCommand{
		cmd:   ETHTOOL_PHYS_ID,
		value: time,
	}
	return i.performIoctl(uintptr(unsafe.Pointer(&cmd)))
}
