package SFF8079

import (
	"fmt"
)

type RateIdentifier byte

func (r RateIdentifier) String() string {
	mapping := map[RateIdentifier]string{
		0x00: "unspecified",
		0x01: "4/2/1G Rate_Select & AS0/AS1",
		0x02: "8/4/2G Rx Rate_Select only",
		0x03: "8/4/2G Independent Rx & Tx Rate_Select",
		0x04: "8/4/2G Tx Rate_Select only",
		0x06: "8/4/2G Independent Rx & Tx Rate_select",
		0x08: "16/8/4G Rx Rate_select only",
		0x0A: "16/8/4G Independent Rx, Tx Rate_select",
		0x0C: "32/16/8G Independent Rx, Tx Rate_Select",
	}

	str, found := mapping[r]
	if found {
		return str
	}
	return "unknown"
}

func (r RateIdentifier) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"ascii": r.String(),
		"hex":   fmt.Sprintf("%#02x", byte(r)),
	}
}
