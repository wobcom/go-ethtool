package eeprom

import (
	"fmt"
)

type OUI uint32

func (o OUI) String() string {
	raw := uint32(o)
	return fmt.Sprintf("%02X:%02X:%02X", (raw>>16)&0xFF, (raw>>8)&0xFF, raw&0xFF)
}

func (o OUI) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"hex":   fmt.Sprintf("%x", uint32(o)&0xFFFFFF),
		"ascii": o.String(),
	}
}

func NewOUI(raw [3]byte) OUI {
	return OUI(uint32(raw[0])<<16 | uint32(raw[1])<<8 | uint32(raw[2]))
}

// TODO provide a function to perform a lookup against a local copy of http://standards-oui.ieee.org/oui.txt
