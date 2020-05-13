package sff8079

import (
	"encoding/json"
	"fmt"
)

// ExtendedIdentifier extended identifier of type of transceiver
type ExtendedIdentifier byte

const (
	// ExtendedIdentifierNotCompliant GBIC not specified / not MOD_DEF compliant
	ExtendedIdentifierNotCompliant ExtendedIdentifier = 0x00
	// ExtendedIdentifierModDef1 GBIC compliant with MOD_DEF1
	ExtendedIdentifierModDef1 ExtendedIdentifier = 0x01
	// ExtendedIdentifierModDef2 GBIC compliant with MOD_DEF2
	ExtendedIdentifierModDef2 ExtendedIdentifier = 0x02
	// ExtendedIdentifierModDef3 GBIC compliant with MOD_DEF3
	ExtendedIdentifierModDef3 ExtendedIdentifier = 0x03
	// ExtendedIdentifierSFP GBIC/SFP function is defined by two-wire interface ID only
	ExtendedIdentifierSFP ExtendedIdentifier = 0x04
	// ExtendedIdentifierModDef5 GBIC compliant with MOD_DEF5
	ExtendedIdentifierModDef5 ExtendedIdentifier = 0x05
	// ExtendedIdentifierModDef6 GBIC compliant with MOD_DEF6
	ExtendedIdentifierModDef6 ExtendedIdentifier = 0x06
	// ExtendedIdentifierModDef7 GBIC compliant with MOD_DEF7
	ExtendedIdentifierModDef7 ExtendedIdentifier = 0x07
)

func (e ExtendedIdentifier) String() string {
	return map[ExtendedIdentifier]string{
		ExtendedIdentifierNotCompliant: "GBIC not specified / not MOD_DEF compliant",
		ExtendedIdentifierModDef1:      "GBIC compliant with MOD_DEF1",
		ExtendedIdentifierModDef2:      "GBIC compliant with MOD_DEF2",
		ExtendedIdentifierModDef3:      "GBIC compliant with MOD_DEF3",
		ExtendedIdentifierSFP:          "GBIC/SFP function is defined by two-wire interface ID only",
		ExtendedIdentifierModDef5:      "GBIC compliant with MOD_DEF5",
		ExtendedIdentifierModDef6:      "GBIC compliant with MOD_DEF6",
		ExtendedIdentifierModDef7:      "GBIC compliant with MOD_DEF7",
	}[e]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (e ExtendedIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": e.String(),
		"hex":   fmt.Sprintf("%#02X", byte(e)),
	})
}
