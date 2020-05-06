package SFF8024

import (
	"encoding/json"
	"fmt"
)

type Identifier byte

const (
	IdentifierUnknown     Identifier = 0x00
	IdentifierGbic        Identifier = 0x01
	IdentifierSoldered    Identifier = 0x02
	IdentifierSfp         Identifier = 0x03
	Identifier300PinXbi   Identifier = 0x04
	IdentifierXenpak      Identifier = 0x05
	IdentifierXfp         Identifier = 0x06
	IdentifierXff         Identifier = 0x07
	IdentifierXfpE        Identifier = 0x08
	IdentifierXpak        Identifier = 0x09
	IdentifierX2          Identifier = 0x0A
	IdentifierDwdmSfp     Identifier = 0x0B
	IdentifierQsfp        Identifier = 0x0C
	IdentifierQsfpPlus    Identifier = 0x0D
	IdentifierCxp         Identifier = 0x0E
	IdentifierHd4x        Identifier = 0x0F
	IdentifierHd8x        Identifier = 0x10
	IdentifierQsfp28      Identifier = 0x11
	IdentifierCxp2        Identifier = 0x12
	IdentifierCdfp        Identifier = 0x13
	IdentifierHd4xFanout  Identifier = 0x14
	IdentifierHd8xFanout  Identifier = 0x15
	IdentifierCdfpStyle3  Identifier = 0x16
	IdentifierMicroQsfp   Identifier = 0x17
	IdentifierLast        Identifier = IdentifierMicroQsfp
	IdentifierVendorStart Identifier = 0x80
	IdentifierVendorEnd   Identifier = 0xFF
)

func (i Identifier) IsKnown() bool {
	return i <= IdentifierLast || (i >= IdentifierVendorStart && i <= IdentifierVendorEnd)
}

func (i Identifier) String() string {
	mapping := map[Identifier]string{
		IdentifierUnknown:    "No module present, unknown, or unspecified",
		IdentifierGbic:       "GBIC",
		IdentifierSoldered:   "Module soldered to motherboard",
		IdentifierSfp:        "SFP",
		Identifier300PinXbi:  "300 pin XBI",
		IdentifierXenpak:     "XENPAK",
		IdentifierXfp:        "XFP",
		IdentifierXff:        "XFF",
		IdentifierXfpE:       "XFP-E",
		IdentifierXpak:       "XPAK",
		IdentifierX2:         "X2",
		IdentifierDwdmSfp:    "DWDM-SFP",
		IdentifierQsfp:       "QSFP",
		IdentifierQsfpPlus:   "QSFP+",
		IdentifierCxp:        "CXP",
		IdentifierHd4x:       "Shielded Mini Multilane HD 4X",
		IdentifierHd8x:       "Shielded Mini Multilane HD 8X",
		IdentifierQsfp28:     "QSFP28",
		IdentifierCxp2:       "CXP2/CXP28",
		IdentifierCdfp:       "CDFP Style 1/Style 2",
		IdentifierHd4xFanout: "Shielded Mini Multilane HD 4X Fanout Cable",
		IdentifierHd8xFanout: "Shielded Mini Multilane HD 8X Fanout Cable",
		IdentifierCdfpStyle3: "CDFP Style 3",
		IdentifierMicroQsfp:  "MicroQSFP",
	}

	str, found := mapping[i]

	if found {
		return str
	} else if i >= IdentifierVendorStart && i <= IdentifierVendorEnd {
		return "Vendor specific"
	} else {
		return "Invalid or unknown"
	}
}

func (i Identifier) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"ascii": i.String(),
		"hex":   fmt.Sprintf("%#02X", byte(i)),
	}
}

func (i Identifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"asci": i.String(),
		"hex":  fmt.Sprintf("%#02X", byte(i)),
	})
}
