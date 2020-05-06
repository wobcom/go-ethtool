package SFF8472

import (
	"encoding/json"
	"fmt"
)

type Compliance byte

const (
	ComplianceRev_9_3  Compliance = 0x01
	ComplianceRev_9_5  Compliance = 0x02
	ComplianceRev_10_2 Compliance = 0x03
	ComplianceRev_10_4 Compliance = 0x04
	ComplianceRev_11_0 Compliance = 0x05
	ComplianceRev_11_3 Compliance = 0x06
	ComplianceRev_11_4 Compliance = 0x07
	ComplianceRev_12_3 Compliance = 0x08
)

func (c Compliance) String() string {
	return map[Compliance]string{
		ComplianceRev_9_3:  "9.3",
		ComplianceRev_9_5:  "9.5",
		ComplianceRev_10_2: "10.2",
		ComplianceRev_10_4: "10.4",
		ComplianceRev_11_0: "11.0",
		ComplianceRev_11_3: "11.3",
		ComplianceRev_11_4: "11.4",
		ComplianceRev_12_3: "12.3",
	}[c]
}

func (c Compliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": c.String(),
		"hex":   fmt.Sprintf("%#02X", byte(c)),
	})
}
