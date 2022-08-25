package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wobcom/go-ethtool/eeprom"
	"github.com/wobcom/go-ethtool/eeprom/sff8472"
	"github.com/wobcom/go-ethtool/eeprom/sff8636"
	"github.com/wobcom/go-ethtool/util"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

type ethtoolModinfo struct {
	Command    uint32
	EepromType uint32
	Length     uint32
	Reserved   [8]uint32
}

type ethtoolEeprom struct {
	Command uint32
	Magic   uint32
	Offset  uint32
	Length  uint32
	Data    [eepromMaxLength]byte
}

const (
	// Get plug-in module information
	getModuleInfoIoctl = 0x00000042
	// Get plug-in module eeprom
	getModuleEepromIoctl = 0x00000043
	// Get EEPROM data
	getEepromDataIoctl = 0x0000000b
	// Set EEPROM data
	setEepromDataIoctl = 0x0000000c
	// Maximum support eeprom length
	eepromMaxLength = 32768
)

// WriteEEPROM writes the given data to the given offset
// Note that not all NIC drivers implement this IOCTL
// and also note, that not all sections of an EEPROM are writeable
func (i *Interface) WriteEEPROM(data []byte, offset uint32) error {
	// Retrieve magic cookie used to avoid accidental changes
	ethtoolEeprom := ethtoolEeprom{
		Command: getEepromDataIoctl,
		Offset:  offset,
		Length:  uint32(len(data)),
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&ethtoolEeprom))); err != nil {
		return errors.Wrapf(err, "ioctl getEepromDataIoctl returned error")
	}
	fmt.Printf("read: %v\n", ethtoolEeprom.Data[0:ethtoolEeprom.Length])

	/* Write data to eeprom */
	for i := 0; i < len(data); i++ {
		ethtoolEeprom.Data[i] = data[i]
	}

	ethtoolEeprom.Command = setEepromDataIoctl

	if err := i.performIoctl(uintptr(unsafe.Pointer(&ethtoolEeprom))); err != nil {
		return errors.Wrapf(err, "iotctl setEepromDataIoctl returend error")
	}
	return nil
}

func isASCII(s []byte) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func (i *Interface) getEEPROM() (eeprom.EEPROM, error) {
	ethtoolModInfo, err := i.getEEPROMModuleInfo()
	if err != nil {
		return nil, errors.Wrapf(err, "Could not retrieve module info for interface %s", i.Name)
	}

	for retryCount := 0; retryCount < 3; retryCount++ {
		if ethtoolModInfo.Length == 0 {
			err = errors.New("EERPOM of length 0 reported")
			continue
		}

		ethtoolEeprom, err := i.getModuleEEPROM(ethtoolModInfo.Length)
		if err != nil {
			continue
		}

		eepromType := eeprom.Type(ethtoolModInfo.EepromType)
		data := ethtoolEeprom.Data[:ethtoolModInfo.Length]

		switch eepromType {
		case eeprom.TypeSFF8472:
			// bypasses a glitch (maybe in the sx_netdev driver?) which would return just garbage
			rawVendorName := data[0x14:0x23]
			if !utf8.Valid(rawVendorName) || strings.HasPrefix(util.GetValidUtf8String(rawVendorName), "/") {
				fmt.Printf("Avoiding weird SFF8472 vendor name for interface  %s\n", i.Name)
			}
			e, err := sff8472.NewEEPROM(data)
			if err != nil {
				fmt.Printf("sff8472: %s\n", err.Error())
			}
			return e, err
		case eeprom.TypeSFF8436:
			return sff8636.NewEEPROM(data)
		case eeprom.TypeSFF8636:
			return sff8636.NewEEPROM(data)
		default:
			err = fmt.Errorf("EEPROM Type %v not supported", eepromType.String())
			continue
		}
	}
	return nil, fmt.Errorf("Could not read EEPROM for interface %s after 3 tries", i.Name)
}

func (i *Interface) getEEPROMModuleInfo() (*ethtoolModinfo, error) {
	ethtoolModInfo := &ethtoolModinfo{
		Command: getModuleInfoIoctl,
	}
	if err := i.performIoctl(uintptr(unsafe.Pointer(ethtoolModInfo))); err != nil {
		return nil, errors.Wrapf(err, "Error running ioctl getModuleInfoIoctl")
	}
	return ethtoolModInfo, nil
}

func (i *Interface) getModuleEEPROM(length uint32) (*ethtoolEeprom, error) {
	ethtoolEeprom := &ethtoolEeprom{
		Command: getModuleEepromIoctl,
		Length:  length,
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(ethtoolEeprom))); err != nil {
		return nil, errors.Wrapf(err, "Error running ioctl getModuleEepromIoctl")
	}

	time.Sleep(5 * time.Millisecond)
	return ethtoolEeprom, nil
}
