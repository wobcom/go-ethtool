package SFF8636

import (
	"fmt"
)

type ExtendedRateSelectCompliance byte

const (
	ExtendedRateSelectComplianceVer1 ExtendedRateSelectCompliance = 1
	ExtendedRateSelectComplianceVer2 ExtendedRateSelectCompliance = 2
)

func (e ExtendedRateSelectCompliance) String() string {
	return map[ExtendedRateSelectCompliance]string{
		ExtendedRateSelectComplianceVer1: "Rate Select Version 1",
		ExtendedRateSelectComplianceVer2: "Rate Select Version 2",
	}[e]
}

func (e ExtendedRateSelectCompliance) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"raw":   fmt.Sprintf("%#02x", byte(e)),
		"ascii": e.String(),
	}
}
