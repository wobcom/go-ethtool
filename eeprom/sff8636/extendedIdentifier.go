package sff8636

import (
	"gitlab.com/wobcom/ethtool/eeprom"
)

// ExtendedIdentifier provides additional information about the free side device.device contains a CDR function and identifies the power cons
type ExtendedIdentifier struct {
	PowerClass             eeprom.PowerClass
	PowerClass8Implemented bool
	CLEICodePresent        bool
	TxCDRPresent           bool
	RxCDRPresent           bool
}

func parsePowerClass(raw byte) eeprom.PowerClass {
	switch raw {
	case 0b00000000:
		return eeprom.PowerClass1
	case 0b01000000:
		return eeprom.PowerClass2
	case 0b10000000:
		return eeprom.PowerClass3
	case 0b11000000:
		return eeprom.PowerClass4
	case 0b11000001:
		return eeprom.PowerClass5
	case 0b11000010:
		return eeprom.PowerClass6
	case 0b11000011:
		return eeprom.PowerClass7
	default:
		return eeprom.PowerClass1
	}
}

// NewExtendedIdentifier parses a byte into a new ExtendedIdentifier instance
func NewExtendedIdentifier(raw byte) *ExtendedIdentifier {
	return &ExtendedIdentifier{
		PowerClass:             parsePowerClass(raw & 0b11000011),
		PowerClass8Implemented: raw&(1<<5) > 0,
		CLEICodePresent:        raw&(1<<4) > 0,
		TxCDRPresent:           raw&(1<<3) > 0,
		RxCDRPresent:           raw&(1<<2) > 0,
	}
}
