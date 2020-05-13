package sff8024

import (
	"encoding/json"
	"fmt"
)

// ConnectorType physical transceiver's connector compliance
type ConnectorType byte

const (
	// ConnectorUnknown Unknown or unspecified
	ConnectorUnknown ConnectorType = 0x00
	// ConnectorSc SC
	ConnectorSc ConnectorType = 0x01
	// ConnectorFcStyle1 Fibre Channel style 1 copper
	ConnectorFcStyle1 ConnectorType = 0x02
	// ConnectorFcStyle2 Fibre Channel style 2 copper
	ConnectorFcStyle2 ConnectorType = 0x03
	// ConnectorBncTnc BNC/TNC
	ConnectorBncTnc ConnectorType = 0x04
	// ConnectorFcCoax Fibre Channel coaxial headers
	ConnectorFcCoax ConnectorType = 0x05
	// ConnectorFiberJack FibreJack
	ConnectorFiberJack ConnectorType = 0x06
	// ConnectorLc LC
	ConnectorLc ConnectorType = 0x07
	// ConnectorMtRj MT-RJ
	ConnectorMtRj ConnectorType = 0x08
	// ConnectorMu MU
	ConnectorMu ConnectorType = 0x09
	// ConnectorSg SG
	ConnectorSg ConnectorType = 0x0A
	// ConnectorOptPtail Optical pigtail
	ConnectorOptPtail ConnectorType = 0x0B
	// ConnectorMpo MPO Parallel Optic
	ConnectorMpo ConnectorType = 0x0C
	// ConnectorMpo2 ConnectorMpo2
	ConnectorMpo2 ConnectorType = 0x0D
	// ConnectorHssdcII HSSDC II
	ConnectorHssdcII ConnectorType = 0x20
	// ConnectorCopperPtail Copper pigtail
	ConnectorCopperPtail ConnectorType = 0x21
	// ConnectorRj45 RJ45
	ConnectorRj45 ConnectorType = 0x22
	// ConnectorNoSeparable No separable connector
	ConnectorNoSeparable ConnectorType = 0x23
	// ConnectorMxc2x16 MXC 2x16
	ConnectorMxc2x16 ConnectorType = 0x24
	// ConnectorVendorStart Start of vendor specific connector types
	ConnectorVendorStart ConnectorType = 0x80
	// ConnectorVendorEnd End of vendor specific connector types
	ConnectorVendorEnd ConnectorType = 0xFF
)

func (c ConnectorType) String() string {
	mapping := map[ConnectorType]string{
		ConnectorUnknown:     "Unknown or unspecified",
		ConnectorSc:          "SC",
		ConnectorFcStyle1:    "Fibre Channel style 1 copper",
		ConnectorFcStyle2:    "Fibre Channel style 2 copper",
		ConnectorBncTnc:      "BNC/TNC",
		ConnectorFcCoax:      "Fibre Channel coaxial headers",
		ConnectorFiberJack:   "FibreJack",
		ConnectorLc:          "LC",
		ConnectorMtRj:        "MT-RJ",
		ConnectorMu:          "MU",
		ConnectorSg:          "SG",
		ConnectorOptPtail:    "Optical pigtail",
		ConnectorMpo:         "MPO Parallel Optic",
		ConnectorMpo2:        "MPO Parallel Optic - 2x16",
		ConnectorHssdcII:     "HSSDC II",
		ConnectorCopperPtail: "Copper pigtail",
		ConnectorRj45:        "RJ45",
		ConnectorNoSeparable: "No separable connector",
		ConnectorMxc2x16:     "MXC 2x16",
	}

	str, found := mapping[c]

	if found {
		return str
	} else if c >= ConnectorVendorStart && c <= ConnectorVendorEnd {
		return "Vendor specific"
	} else {
		return "invalid or unknown"
	}
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (c ConnectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": c.String(),
		"hex":   fmt.Sprintf("%#02x", byte(c)),
	})
}
