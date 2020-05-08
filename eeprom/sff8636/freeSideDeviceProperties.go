package sff8636

// FreeSideDeviceProperties as defined in SFF-8636 rev 2.10a Table 6-13
type FreeSideDeviceProperties struct {
	MaxPowerConsumption   float64
	PropagationDelay      float64
	AdvancedLowPowerMode  AdvancedLowPowerMode
	FarSideManaged        bool
	MinOperatingVoltage   float64
	FarEndImplementation  FarEndImplementation
	NearEndImplementation NearEndImplementation
}

// FarEndImplementation as defined in SFF-8636 rev 2.10a Table 6-13
type FarEndImplementation byte

const (
    // FarEndUnspecified Far end is unspecified
	FarEndUnspecified                        FarEndImplementation = 0b000
    // FarEnd4ChannelsOrModule4ChannelConnector Cable with single far-end with 4 channels implemented, or separable module with a 4-channel connector
	FarEnd4ChannelsOrModule4ChannelConnector FarEndImplementation = 0b001
    // FarEnd2ChannelsOrModule2ChannelConnector Cable with single far-end with 2 channels implemented, or separable module with a 2-channel connector
	FarEnd2ChannelsOrModule2ChannelConnector FarEndImplementation = 0b010
    // FarEnd1ChannelOrModule1ChannelConnector Cable with single far-end with 1 channel implmented, or separable module with a 1-channel connector
	FarEnd1ChannelOrModule1ChannelConnector  FarEndImplementation = 0b011
    // FarEnd4x1BreakOut 4 far-ends with 1 channel implemented in each (i.e. 4x1 break out)
	FarEnd4x1BreakOut                        FarEndImplementation = 0b100
    // FarEnd2x2BreakOut 2 far-ends with 2 channels implemented in each (i.e. 2x2 break out)
	FarEnd2x2BreakOut                        FarEndImplementation = 0b101
    // FarEnd2x1BreakOut 2 far-ends with 1 channel implemented in each (i.e. 2x1 break out)
	FarEnd2x1BreakOut                        FarEndImplementation = 0b110
)

// NearEndImplementation as defined in SFF-8636 rev 2.10a Table 6-13
type NearEndImplementation struct {
	ChannelImplemented [4]bool `json:"channelImplemented"`
}

// AdvancedLowPowerMode as defined in SFF-8636 rev 2.10a Table 6-13
type AdvancedLowPowerMode byte

const (
    // AdvancedLowPowerMode1_5 1.5 W or higher
	AdvancedLowPowerMode1_5  AdvancedLowPowerMode = 0b0000
    // AdvancedLowPowerMode1_0 no more than 1 W
	AdvancedLowPowerMode1_0  AdvancedLowPowerMode = 0b0001
    // AdvancedLowPowerMode0_75 no more than 0.75W
	AdvancedLowPowerMode0_75 AdvancedLowPowerMode = 0b0010
    // AdvancedLowPowerMode0_5 no more than 0.5W
	AdvancedLowPowerMode0_5  AdvancedLowPowerMode = 0b0011
)

func (f FarEndImplementation) String() string {
	return map[FarEndImplementation]string{
		FarEndUnspecified:                        "Far end is unspecified",
		FarEnd4ChannelsOrModule4ChannelConnector: "Cable with single far-end with 4 channels implemented, or separable module with a 4-channel connector",
		FarEnd2ChannelsOrModule2ChannelConnector: "Cable with single far-end with 2 channels implemented, or separable module with a 2-channel connector",
		FarEnd4x1BreakOut:                        "4 far-ends with 1 channel implemented in each (i.e. 4x1 break out)",
		FarEnd2x2BreakOut:                        "2 far-ends with 2 channels implemented in each (i.e. 2x2 break out)",
		FarEnd2x1BreakOut:                        "2 far-ends with 1 channel implemented in each (i.e. 2x1 break out)",
	}[f]
}

func (a AdvancedLowPowerMode) String() string {
	return map[AdvancedLowPowerMode]string{
        AdvancedLowPowerMode1_5:  "1.5 W or higher",
		AdvancedLowPowerMode1_0:  "no more than 1 W",
		AdvancedLowPowerMode0_75: "no more than 0.75W",
		AdvancedLowPowerMode0_5:  "no more than 0.5W",
	}[a]
}

func parseDevicePowerConsumption(raw byte) float64 {
	return float64(raw) * 0.1
}

func parsePropagationDelay(msb byte, lsb byte) float64 {
	return float64((uint16(msb)<<8 | uint16(lsb))) * 10
}

func parseMinOperatingVoltage(raw byte) float64 {
	switch raw {
	case 0b000:
		return 3.3
	case 0b001:
		return 2.5
	case 0b010:
		return 1.8
	default:
		return 0
	}
}

func parseNearEndImplementation(raw byte) NearEndImplementation {
	n := NearEndImplementation{}
	for i := 0; i < 4; i++ {
		n.ChannelImplemented[i] = raw&(1<<i) > 0
	}
	return n
}

// NewFreeSideDeviceProperties parses [10]byte into a new FreeSideDeviceProperties instance
func NewFreeSideDeviceProperties(raw [10]byte) *FreeSideDeviceProperties {
	return &FreeSideDeviceProperties{
		MaxPowerConsumption:   parseDevicePowerConsumption(raw[0]),
		PropagationDelay:      parsePropagationDelay(raw[1], raw[2]),
		AdvancedLowPowerMode:  AdvancedLowPowerMode((raw[3] & 0b11110000) >> 4),
		FarSideManaged:        raw[3]&(1<<3) > 0,
		MinOperatingVoltage:   parseMinOperatingVoltage(raw[3] & 0b111),
		FarEndImplementation:  FarEndImplementation((raw[6] & 0b01110000) >> 4),
		NearEndImplementation: parseNearEndImplementation(raw[6] & 0b000001111),
	}
}
