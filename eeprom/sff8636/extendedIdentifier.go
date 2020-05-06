package SFF8636

import (
	"gitlab.com/wobcom/golang-ethtool/eeprom"
)

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
		return eeprom.PWR_CLASS_1
	case 0b01000000:
		return eeprom.PWR_CLASS_2
	case 0b10000000:
		return eeprom.PWR_CLASS_3
	case 0b11000000:
		return eeprom.PWR_CLASS_4
	case 0b11000001:
		return eeprom.PWR_CLASS_5
	case 0b11000010:
		return eeprom.PWR_CLASS_6
	case 0b11000011:
		return eeprom.PWR_CLASS_7
	default:
		return eeprom.PWR_CLASS_1
	}
}

func NewExtendedIdentifier(raw byte) *ExtendedIdentifier {
	return &ExtendedIdentifier{
		PowerClass:             parsePowerClass(raw & 0b11000011),
		PowerClass8Implemented: raw&(1<<5) > 0,
		CLEICodePresent:        raw&(1<<4) > 0,
		TxCDRPresent:           raw&(1<<3) > 0,
		RxCDRPresent:           raw&(1<<2) > 0,
	}
}
