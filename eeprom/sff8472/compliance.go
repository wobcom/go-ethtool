package sff8472

import (
	"encoding/json"
	"fmt"
)

// Compliance to the SFF-8472 standard
type Compliance byte

const (
	// ComplianceRev9dot3 9.3
	ComplianceRev9dot3 Compliance = 0x01
	// ComplianceRev9dot5 9.5
	ComplianceRev9dot5 Compliance = 0x02
	// ComplianceRev10dot2 10.2
	ComplianceRev10dot2 Compliance = 0x03
	// ComplianceRev10dot4 10.4
	ComplianceRev10dot4 Compliance = 0x04
	// ComplianceRev11dot0 11.0
	ComplianceRev11dot0 Compliance = 0x05
	// ComplianceRev11dot3 11.3
	ComplianceRev11dot3 Compliance = 0x06
	// ComplianceRev11dot4 11.4
	ComplianceRev11dot4 Compliance = 0x07
	// ComplianceRev12dot3 12.3
	ComplianceRev12dot3 Compliance = 0x08
)

func (c Compliance) String() string {
	return map[Compliance]string{
		ComplianceRev9dot3:  "9.3",
		ComplianceRev9dot5:  "9.5",
		ComplianceRev10dot2: "10.2",
		ComplianceRev10dot4: "10.4",
		ComplianceRev11dot0: "11.0",
		ComplianceRev11dot3: "11.3",
		ComplianceRev11dot4: "11.4",
		ComplianceRev12dot3: "12.3",
	}[c]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (c Compliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": c.String(),
		"hex":   fmt.Sprintf("%#02X", byte(c)),
	})
}
