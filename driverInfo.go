package ethtool

import (
	"bytes"
	"unsafe"
)

const (
	ETHTOOL_GDRVINFO = 0x00000003 /* Get driver info */
)

type ethtoolDrvInfo struct {
	cmd          uint32
	driver       [32]byte
	version      [32]byte
	fw_version   [32]byte
	bus_info     [32]byte
	erom_version [32]byte
	reserved2    [12]byte
	n_priv_flags uint32
	n_stats      uint32
	testinfo_len uint32
	eedump_len   uint32
	regdump_len  uint32
}

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

func NewDriverInfo(ethtoolDrvInfo *ethtoolDrvInfo) *DriverInfo {
	return &DriverInfo{
		DriverName:          string(bytes.Trim(ethtoolDrvInfo.driver[:], "\x00")),
		DriverVersion:       string(bytes.Trim(ethtoolDrvInfo.version[:], "\x00")),
		FirmwareVersion:     string(bytes.Trim(ethtoolDrvInfo.fw_version[:], "\x00")),
		BusInfo:             string(bytes.Trim(ethtoolDrvInfo.bus_info[:], "\x00")),
		ExpansionRomVersion: string(bytes.Trim(ethtoolDrvInfo.erom_version[:], "\x00")),
		EEPROMLength:        ethtoolDrvInfo.eedump_len,
		RegDumpLength:       ethtoolDrvInfo.regdump_len,
	}
}

func (i *Interface) getDriverInfo() (*DriverInfo, error) {
	ethtoolDriverInfo := ethtoolDrvInfo{
		cmd: ETHTOOL_GDRVINFO,
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&ethtoolDriverInfo))); err != nil {
		return nil, err
	}

	return NewDriverInfo(&ethtoolDriverInfo), nil
}
