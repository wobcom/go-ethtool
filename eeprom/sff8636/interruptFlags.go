package sff8636

// InterruptFlags as defined in SFF-8636 rev 2.10a table 6-4
type InterruptFlags struct {
	ChannelInterrupt       [4]ChannelInterrupt
	FreeSideInterruptFlags FreeSideInterruptFlags
}

// FreeSideInterruptFlags freeside general interrupts
type FreeSideInterruptFlags struct {
	TemperatureAlarm       Alarm
	TCReadinessFlag        bool
	InitializationComplete bool
	VoltageAlarm           Alarm
}

// ChannelInterrupt channel specific interrupts
type ChannelInterrupt struct {
	TxLOS           bool
	RxLOS           bool
	AdaptiveEQFault bool
	TxFault         bool
	TxLOL           bool
	RxLOL           bool
	RxPowerAlarm    Alarm
	TxPowerAlarm    Alarm
	BiasAlarm       Alarm
}

// Alarm interrupt flags
type Alarm struct {
	HighAlarm   bool
	HighWarning bool
	LowAlarm    bool
	LowWarning  bool
}

var interruptFlagsMemoryMap = map[uint]map[uint]func(*InterruptFlags, bool){
	0x00: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxLOS = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxLOS = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxLOS = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxLOS = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].RxLOS = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].RxLOS = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].RxLOS = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].RxLOS = v },
	},
	0x01: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].AdaptiveEQFault = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].AdaptiveEQFault = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].AdaptiveEQFault = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].AdaptiveEQFault = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxFault = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxFault = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxFault = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxFault = v },
	},
	0x02: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxLOL = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxLOL = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxLOL = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxLOL = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].RxLOL = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].RxLOL = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].RxLOL = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].RxLOL = v },
	},
	0x03: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.TemperatureAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.TemperatureAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.TemperatureAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.TemperatureAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { /* reserved */ },
		0x02: func(i *InterruptFlags, v bool) { /* reserved */ },
		0x01: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.TCReadinessFlag = v },
		0x00: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.InitializationComplete = v },
	},
	0x04: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.VoltageAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.VoltageAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.VoltageAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.FreeSideInterruptFlags.VoltageAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { /* reserved */ },
		0x02: func(i *InterruptFlags, v bool) { /* reserved */ },
		0x01: func(i *InterruptFlags, v bool) { /* reserved */ },
		0x00: func(i *InterruptFlags, v bool) { /* reserved */ },
	},
	// 0x05: vendor specific
	0x06: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].RxPowerAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].RxPowerAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].RxPowerAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].RxPowerAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].RxPowerAlarm.HighAlarm = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].RxPowerAlarm.LowAlarm = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].RxPowerAlarm.HighWarning = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].RxPowerAlarm.LowWarning = v },
	},
	0x07: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].RxPowerAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].RxPowerAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].RxPowerAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].RxPowerAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].RxPowerAlarm.HighAlarm = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].RxPowerAlarm.LowAlarm = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].RxPowerAlarm.HighWarning = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].RxPowerAlarm.LowWarning = v },
	},
	0x08: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].BiasAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].BiasAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].BiasAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].BiasAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].BiasAlarm.HighAlarm = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].BiasAlarm.LowAlarm = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].BiasAlarm.HighWarning = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].BiasAlarm.LowWarning = v },
	},
	0x09: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].BiasAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].BiasAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].BiasAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].BiasAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].BiasAlarm.HighAlarm = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].BiasAlarm.LowAlarm = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].BiasAlarm.HighWarning = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].BiasAlarm.LowWarning = v },
	},
	0x0A: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxPowerAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxPowerAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxPowerAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[0].TxPowerAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxPowerAlarm.HighAlarm = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxPowerAlarm.LowAlarm = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxPowerAlarm.HighWarning = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[1].TxPowerAlarm.LowWarning = v },
	},
	0x0B: map[uint]func(*InterruptFlags, bool){
		0x07: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxPowerAlarm.HighAlarm = v },
		0x06: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxPowerAlarm.LowAlarm = v },
		0x05: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxPowerAlarm.HighWarning = v },
		0x04: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[2].TxPowerAlarm.LowWarning = v },
		0x03: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxPowerAlarm.HighAlarm = v },
		0x02: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxPowerAlarm.LowAlarm = v },
		0x01: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxPowerAlarm.HighWarning = v },
		0x00: func(i *InterruptFlags, v bool) { i.ChannelInterrupt[3].TxPowerAlarm.LowWarning = v },
	},
	// 0x0C-0x0D: Channel monitor flags, set 4
	// 0x0E-0x0F: Channel monitor flags, set 5
	// 0x10-0x12: Vendor specific
}

// NewInterruptFlags parses [19]byte into a new InterruptFlags instance
func NewInterruptFlags(raw [19]byte) *InterruptFlags {
	i := &InterruptFlags{}
	for byteIndex, bitmap := range interruptFlagsMemoryMap {
		for bitIndex, callback := range bitmap {
			value := raw[byteIndex]&(1<<bitIndex) > 0
			callback(i, value)
		}
	}
	return i
}
