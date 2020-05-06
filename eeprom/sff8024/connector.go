package SFF8024

import (
	"fmt"
)

type ConnectorType byte

const (
	ConnectorUnknown     ConnectorType = 0x00
	ConnectorSc          ConnectorType = 0x01
	ConnectorFcStyle1    ConnectorType = 0x02
	ConnectorFcStyle2    ConnectorType = 0x03
	ConnectorBncTnc      ConnectorType = 0x04
	ConnectorFcCoax      ConnectorType = 0x05
	ConnectorFiberJack   ConnectorType = 0x06
	ConnectorLc          ConnectorType = 0x07
	ConnectorMtRj        ConnectorType = 0x08
	ConnectorMu          ConnectorType = 0x09
	ConnectorSg          ConnectorType = 0x0A
	ConnectorOptPtail    ConnectorType = 0x0B
	ConnectorMpo         ConnectorType = 0x0C
	ConnectorMpo2        ConnectorType = 0x0D
	ConnectorHssdcII     ConnectorType = 0x20
	ConnectorCopperPtail ConnectorType = 0x21
	ConnectorRj45        ConnectorType = 0x22
	ConnectorNoSeparable ConnectorType = 0x23
	ConnectorMxc2x16     ConnectorType = 0x24
	ConnectorLast        ConnectorType = ConnectorMxc2x16
	ConnectorVendorStart ConnectorType = 0x80
	ConnectorVendorEnd   ConnectorType = 0xFF
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

func (c ConnectorType) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"ascii": c.String(),
		"hex":   fmt.Sprintf("%#02x", byte(c)),
	}
}
