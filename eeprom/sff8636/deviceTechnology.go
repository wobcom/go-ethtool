package sff8636

import (
	"encoding/json"
	"fmt"
)

// DeviceTechnology aspects of the device or cable technology used
type DeviceTechnology struct {
	WavelengthControl     bool
	CooledTransmitter     bool
	APDDetector           bool
	TransmitterTunable    bool
	TransmitterTechnology TransmitterTechnology
}

// TransmitterTechnology transmitter technology
type TransmitterTechnology byte

const (
	// TransmitterTechnology850nmVCSEL 850 nm VCSEL
	TransmitterTechnology850nmVCSEL TransmitterTechnology = 0b0000
	// TransmitterTechnology1310nmVCSEL 1310 nm VCSEL
	TransmitterTechnology1310nmVCSEL TransmitterTechnology = 0b0001
	// TransmitterTechnology1550nmVCSEL 1550 nm VCSEL
	TransmitterTechnology1550nmVCSEL TransmitterTechnology = 0b0010
	// TransmitterTechnology1310nmFP 1310 nm FP
	TransmitterTechnology1310nmFP TransmitterTechnology = 0b0011
	// TransmitterTechnology1310nmDFB 1310 nm DFB
	TransmitterTechnology1310nmDFB TransmitterTechnology = 0b0100
	// TransmitterTechnology1550nmDFB 1550 nm DFB
	TransmitterTechnology1550nmDFB TransmitterTechnology = 0b0101
	// TransmitterTechnology1310nmEML 1310 nm EML
	TransmitterTechnology1310nmEML TransmitterTechnology = 0b0110
	// TransmitterTechnology1550nmEML 1550 nm EML
	TransmitterTechnology1550nmEML TransmitterTechnology = 0b0111
	// TransmitterTechnologyOther Other / Undefined
	TransmitterTechnologyOther TransmitterTechnology = 0b1000
	// TransmitterTechnology1490nmDFB 1490 nm DFB
	TransmitterTechnology1490nmDFB TransmitterTechnology = 0b1001
	// TransmitterTechnologyCopperCableUnequalized Copper cable unequalized
	TransmitterTechnologyCopperCableUnequalized TransmitterTechnology = 0b1010
	// TransmitterTechnologyCopperCablePassiveEqualized Copper cable passive equalized
	TransmitterTechnologyCopperCablePassiveEqualized TransmitterTechnology = 0b1011
	// TransmitterTechnologyCopperCableNearAndFarEndLimitingActiveEqualizers Copper cable, near and far end limiting active equalizers
	TransmitterTechnologyCopperCableNearAndFarEndLimitingActiveEqualizers TransmitterTechnology = 0b1100
	// TransmitterTechnologyCopperCableFarEndLimitingActiveEqualizers Copper cable, far end limiting active equalizers
	TransmitterTechnologyCopperCableFarEndLimitingActiveEqualizers TransmitterTechnology = 0b1101
	// TransmitterTechnologyCopperCableNearEndLimitingActiveEqualizers Copper cable, near end limiting active equalizers
	TransmitterTechnologyCopperCableNearEndLimitingActiveEqualizers TransmitterTechnology = 0b1110
	// TransmitterTechnologyCopperCableLinearActiveEqualizers Copper cable, linear active equalizers
	TransmitterTechnologyCopperCableLinearActiveEqualizers TransmitterTechnology = 0b1111
)

func (t TransmitterTechnology) String() string {
	return map[TransmitterTechnology]string{
		TransmitterTechnology850nmVCSEL:                                       "850 nm VCSEL",
		TransmitterTechnology1310nmVCSEL:                                      "1310 nm VCSEL",
		TransmitterTechnology1550nmVCSEL:                                      "1550 nm VCSEL",
		TransmitterTechnology1310nmFP:                                         "1310 nm FP",
		TransmitterTechnology1310nmDFB:                                        "1310 nm DFB",
		TransmitterTechnology1550nmDFB:                                        "1550 nm DFB",
		TransmitterTechnology1310nmEML:                                        "1310 nm EML",
		TransmitterTechnology1550nmEML:                                        "1550 nm EML",
		TransmitterTechnologyOther:                                            "Other / Undefined",
		TransmitterTechnology1490nmDFB:                                        "1490 nm DFB",
		TransmitterTechnologyCopperCableUnequalized:                           "Copper cable unequalized",
		TransmitterTechnologyCopperCablePassiveEqualized:                      "Copper cable passive equalized",
		TransmitterTechnologyCopperCableNearAndFarEndLimitingActiveEqualizers: "Copper cable, near and far end limiting active equalizers",
		TransmitterTechnologyCopperCableFarEndLimitingActiveEqualizers:        "Copper cable, far end limiting active equalizers",
		TransmitterTechnologyCopperCableNearEndLimitingActiveEqualizers:       "Copper cable, near end limiting active equalizers",
		TransmitterTechnologyCopperCableLinearActiveEqualizers:                "Copper cable, linear active equalizers",
	}[t]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (t TransmitterTechnology) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"raw":   fmt.Sprintf("%#02X", byte(t)),
		"ascii": t.String(),
	})
}

// NewDeviceTechnology parses a byte into a new DeviceTechnology instance
func NewDeviceTechnology(raw byte) *DeviceTechnology {
	return &DeviceTechnology{
		WavelengthControl:     raw&(1<<3) > 0,
		CooledTransmitter:     raw&(1<<2) > 0,
		APDDetector:           raw&(1<<1) > 0,
		TransmitterTunable:    raw&(1<<0) > 0,
		TransmitterTechnology: TransmitterTechnology((raw & 0b11110000) >> 4),
	}
}
