package sff8636

import (
	"encoding/json"
	"fmt"
)

// ExtendedRateSelectCompliance The Extended Rate Select Compliance field is used to allow a single free side device the flexibility to comply with single or multiple Extended Rate Select definitions.
type ExtendedRateSelectCompliance byte

const (
	// ExtendedRateSelectComplianceVer1 Rate Select Version 1
	ExtendedRateSelectComplianceVer1 ExtendedRateSelectCompliance = 1
	// ExtendedRateSelectComplianceVer2 Rate Select Version 2
	ExtendedRateSelectComplianceVer2 ExtendedRateSelectCompliance = 2
)

func (e ExtendedRateSelectCompliance) String() string {
	return map[ExtendedRateSelectCompliance]string{
		ExtendedRateSelectComplianceVer1: "Rate Select Version 1",
		ExtendedRateSelectComplianceVer2: "Rate Select Version 2",
	}[e]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (e ExtendedRateSelectCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"raw":   fmt.Sprintf("%#02x", byte(e)),
		"ascii": e.String(),
	})
}
