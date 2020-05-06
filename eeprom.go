package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/wobcom/ethtool/eeprom"
	"gitlab.com/wobcom/ethtool/eeprom/sff8472"
	"gitlab.com/wobcom/ethtool/eeprom/sff8636"
	"unsafe"
)

type EthtoolModinfo struct {
	Command    uint32
	EepromType uint32
	Length     uint32
	Reserved   [8]uint32
}

type EthtoolEeprom struct {
	Command uint32
	Magic   uint32
	Offset  uint32
	Length  uint32
	Data    [EEPROM_MAX_LENGTH]byte
}

const (
	ETHTOOL_GMODULEINFO   = 0x00000042 /* Get plug-in module information */
	ETHTOOL_GMODULEEEPROM = 0x00000043 /* Get plug-in module eeprom */
	ETHTOOL_GEEPROM       = 0x0000000b /* Get EEPROM data */
	ETHTOOL_SEEPROM       = 0x0000000c /* Set EEPROM data. */
	EEPROM_MAX_LENGTH     = 32768
)

func (i *Interface) WriteEEPROM(data []byte, offset uint32) error {
	// TODO check lengths instead of letting ioctl fail
	/* Retrieve magic cookie used to avoid accidental changes */
	ethtoolEeprom := EthtoolEeprom{
		Command: ETHTOOL_GEEPROM,
		Offset:  offset,
		Length:  uint32(len(data)),
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&ethtoolEeprom))); err != nil {
		return errors.Wrapf(err, "ioctl ETHTOOL_GEEPROM returned error")
	}
	fmt.Printf("read: %v\n", ethtoolEeprom.Data[0:ethtoolEeprom.Length])

	/* Write data to eeprom */
	for i := 0; i < len(data); i++ {
		ethtoolEeprom.Data[i] = data[i]
	}

	ethtoolEeprom.Command = ETHTOOL_SEEPROM

	if err := i.performIoctl(uintptr(unsafe.Pointer(&ethtoolEeprom))); err != nil {
		return errors.Wrapf(err, "iotctl ETHTOOL_SEEPROM returend error")
	}
	return nil
}

func (i *Interface) getEEPROM() (eeprom.EEPROM, error) {
	ethtoolModInfo, err := i.getEEPROMModuleInfo()
	if err != nil {
		return nil, errors.Wrapf(err, "Could not retrieve module info for interface %s", i.Name)
	}

	if ethtoolModInfo.Length == 0 {
		return nil, errors.New("EERPOM of length 0 reported")
	}

	ethtoolEeprom, err := i.getModuleEEPROM(ethtoolModInfo.Length)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not retrieve module raw data")
	}

	eepromType := eeprom.EEPROMType(ethtoolModInfo.EepromType)
	data := ethtoolEeprom.Data[:ethtoolModInfo.Length]

	switch eepromType {
	case eeprom.ETH_MODULE_SFF_8472:
		return SFF8472.NewEEPROM(data)
	case eeprom.ETH_MODULE_SFF_8436:
		return SFF8636.NewEEPROM(data)
	case eeprom.ETH_MODULE_SFF_8636:
		return SFF8636.NewEEPROM(data)
	default:
		return nil, errors.New(fmt.Sprintf("EEPROM Type %v not supported", eepromType.String()))
	}
}

func (i *Interface) getEEPROMModuleInfo() (*EthtoolModinfo, error) {
	ethtoolModInfo := &EthtoolModinfo{
		Command: ETHTOOL_GMODULEINFO,
	}
	if err := i.performIoctl(uintptr(unsafe.Pointer(ethtoolModInfo))); err != nil {
		return nil, errors.Wrapf(err, "Error running ioctl ETHTOOL_GMODULEINFO")
	}
	return ethtoolModInfo, nil
}

func (i *Interface) getModuleEEPROM(length uint32) (*EthtoolEeprom, error) {
	ethtoolEeprom := &EthtoolEeprom{
		Command: ETHTOOL_GMODULEEEPROM,
		Length:  length,
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(ethtoolEeprom))); err != nil {
		return nil, errors.Wrapf(err, "Error running ioctl ETHTOOL_GMODULEEEPROM")
	}
	return ethtoolEeprom, nil
}
