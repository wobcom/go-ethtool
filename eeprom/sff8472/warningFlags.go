package sff8472

// WarningFlags Diagnostic Warning Flag Status Bits  as of SFF-8472 rev 12.3 table 9-12
type WarningFlags struct {
	Temperature      WarningFlagStatus
	Voltage          WarningFlagStatus
	Bias             WarningFlagStatus
	TxPower          WarningFlagStatus
	RxPower          WarningFlagStatus
	LaserTemperature WarningFlagStatus
	TecCurrent       WarningFlagStatus
}

// WarningFlagStatus provides high / low warning flag. A true represent that an alarm condition is present
type WarningFlagStatus struct {
	HighWarning bool
	LowWarning  bool
}

var warningFlagsMemoryMap = map[uint]map[uint]func(*WarningFlags, bool){
	0x00: map[uint]func(*WarningFlags, bool){
		0x07: func(a *WarningFlags, b bool) { a.Temperature.HighWarning = b },
		0x06: func(a *WarningFlags, b bool) { a.Temperature.LowWarning = b },
		0x05: func(a *WarningFlags, b bool) { a.Voltage.HighWarning = b },
		0x04: func(a *WarningFlags, b bool) { a.Voltage.LowWarning = b },
		0x03: func(a *WarningFlags, b bool) { a.Bias.HighWarning = b },
		0x02: func(a *WarningFlags, b bool) { a.Bias.LowWarning = b },
		0x01: func(a *WarningFlags, b bool) { a.TxPower.HighWarning = b },
		0x00: func(a *WarningFlags, b bool) { a.TxPower.LowWarning = b },
	},
	0x01: map[uint]func(*WarningFlags, bool){
		0x07: func(a *WarningFlags, b bool) { a.RxPower.HighWarning = b },
		0x06: func(a *WarningFlags, b bool) { a.RxPower.LowWarning = b },
		0x05: func(a *WarningFlags, b bool) { a.LaserTemperature.HighWarning = b },
		0x04: func(a *WarningFlags, b bool) { a.LaserTemperature.LowWarning = b },
		0x03: func(a *WarningFlags, b bool) { a.TecCurrent.HighWarning = b },
		0x02: func(a *WarningFlags, b bool) { a.TecCurrent.LowWarning = b },
	},
}

// NewWarningFlags parses [2]byte into a new WarningFlags instance
func NewWarningFlags(raw [2]byte) *WarningFlags {
	a := &WarningFlags{}
	for byteOffset, bitMap := range warningFlagsMemoryMap {
		for bitOffset, callback := range bitMap {
			callback(a, raw[byteOffset]&(1<<bitOffset) > 0)
		}
	}
	return a
}
