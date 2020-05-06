package SFF8079

import (
	"fmt"
)

type ExtendedIdentifier byte

const (
	ExtendedIdentifierNotCompliant ExtendedIdentifier = 0x00
	ExtendedIdentifierSFP          ExtendedIdentifier = 0x04
	ExtendedIdentifierModDef       ExtendedIdentifier = 0x07
)

func (e ExtendedIdentifier) String() string {
	if e == ExtendedIdentifierNotCompliant {
		return "GBIC not specified / not MOD_DEF compliant"
	} else if e == ExtendedIdentifierSFP {
		return "GBIC/SFP defined by 2-wire interface ID"
	} else if e <= ExtendedIdentifierModDef {
		return fmt.Sprintf("GBIC compliant with MOD_DEF %d", byte(e))
	} else {
		return "invalid or unknown"
	}
}

func (e ExtendedIdentifier) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"ascii": e.String(),
		"hex":   fmt.Sprintf("%#02x", byte(e)),
	}
}
