package ethtool

import (
	"bytes"
	"unsafe"
)

const (
    // GetDriverInfoIoctl ioctl command number for "Get driver info"
	GetDriverInfoIoctl = 0x00000003
)

type ethtoolDrvInfo struct {
	cmd          uint32
	driver       [32]byte
	version      [32]byte
	fwVersion    [32]byte
	busInfo      [32]byte
	eromVersion  [32]byte
	reserved2    [12]byte
	nPrivFlags   uint32
	nStats       uint32
	testinfoLen  uint32
	eedumpLen    uint32
	regdumpLen   uint32
}

// DriverInfo network interface driver information
type DriverInfo struct {
	// Driver short name. This should normally match the name
	// in its bus driver structure (e.g. pci_driver::name).  Must
	// not be an empty string.
	DriverName string
	// Driver version string; may be an empty string
	DriverVersion string
	// Firmware version string; may be an empty string
	FirmwareVersion string
	// Device bus address
	BusInfo string
	// Expansion ROM version string; may be an empty string
	ExpansionRomVersion string
	// Size of EEPROM dumps in byte
	EEPROMLength uint32
	// Size of register dumps in byte
	RegDumpLength uint32
}

func newDriverInfo(ethtoolDrvInfo *ethtoolDrvInfo) *DriverInfo {
	return &DriverInfo{
		DriverName:          string(bytes.Trim(ethtoolDrvInfo.driver[:], "\x00")),
		DriverVersion:       string(bytes.Trim(ethtoolDrvInfo.version[:], "\x00")),
		FirmwareVersion:     string(bytes.Trim(ethtoolDrvInfo.fwVersion[:], "\x00")),
		BusInfo:             string(bytes.Trim(ethtoolDrvInfo.busInfo[:], "\x00")),
		ExpansionRomVersion: string(bytes.Trim(ethtoolDrvInfo.eromVersion[:], "\x00")),
		EEPROMLength:        ethtoolDrvInfo.eedumpLen,
		RegDumpLength:       ethtoolDrvInfo.regdumpLen,
	}
}

func (i *Interface) getDriverInfo() (*DriverInfo, error) {
	ethtoolDriverInfo := ethtoolDrvInfo{
		cmd: GetDriverInfoIoctl,
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&ethtoolDriverInfo))); err != nil {
		return nil, err
	}

	return newDriverInfo(&ethtoolDriverInfo), nil
}
