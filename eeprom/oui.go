package eeprom

import (
	"encoding/json"
	"fmt"
)

// OUI IEEE company ID
type OUI uint32

func (o OUI) String() string {
	raw := uint32(o)
	return fmt.Sprintf("%02X:%02X:%02X", (raw>>16)&0xFF, (raw>>8)&0xFF, raw&0xFF)
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (o OUI) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"hex":   fmt.Sprintf("%x", uint32(o)&0xFFFFFF),
		"ascii": o.String(),
	})
}

// NewOUI parses [3]byte into an OUI instance
func NewOUI(raw [3]byte) OUI {
	return OUI(uint32(raw[0])<<16 | uint32(raw[1])<<8 | uint32(raw[2]))
}

// TODO provide a function to perform a lookup against a local copy of http://standards-oui.ieee.org/oui.txt
