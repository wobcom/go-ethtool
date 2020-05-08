package sff8636

// InterruptMasks as defined in SFF-8636 rev 2.10a table 6-12
type InterruptMasks struct {
	ChannelInterruptMasks  [4]ChannelInterruptMasks `json:"channelInterruptMasks"`
	FreeSideInterruptMasks FreeSideInterruptMasks   `json:"freeSideInterruptMasks"`
}

// FreeSideInterruptMasks general free side interrupt masks
type FreeSideInterruptMasks struct {
	TemperatureAlarmMask AlarmMask
	TCReadinessFlagMask  bool
	VoltageAlarmMask     AlarmMask
}

// ChannelInterruptMasks channel specific interrupt masks
type ChannelInterruptMasks struct {
	TxLOSMask           bool `json:"txLOSMask"`
	RxLOSMask           bool `json:"rxLOSMask"`
	TxFaultMask         bool `json:"txFaultMask"`
	AdaptiveEQFaultMask bool `json:"adaptiveEQFaultMask"`
	TxLOLMask           bool `json:"txLOLMask"`
	RxLOLMask           bool `json:"rxLOLMask"`
}

// AlarmMask alarm interrupt masks
type AlarmMask struct {
	HighAlarmMask   bool `json:"highAlarmMask"`
	HighWarningMask bool `json:"highWarningMask"`
	LowAlarmMask    bool `json:"lowAlarmMask"`
	LowWarningMask  bool `json:"lowWarningMask"`
}

var interruptMasksMemoryMap = map[uint]map[uint]func(*InterruptMasks, bool){
	0x00: map[uint]func(*InterruptMasks, bool){
		0x07: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[3].TxLOSMask = b },
		0x06: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[2].TxLOSMask = b },
		0x05: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[1].TxLOSMask = b },
		0x04: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[0].TxLOSMask = b },
		0x03: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[3].RxLOSMask = b },
		0x02: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[2].RxLOSMask = b },
		0x01: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[1].RxLOSMask = b },
		0x00: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[0].RxLOSMask = b },
	},
	0x01: map[uint]func(*InterruptMasks, bool){
		0x07: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[3].AdaptiveEQFaultMask = b },
		0x06: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[2].AdaptiveEQFaultMask = b },
		0x05: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[1].AdaptiveEQFaultMask = b },
		0x04: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[0].AdaptiveEQFaultMask = b },
		0x03: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[3].TxFaultMask = b },
		0x02: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[2].TxFaultMask = b },
		0x01: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[1].TxFaultMask = b },
		0x00: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[0].TxFaultMask = b },
	},
	0x02: map[uint]func(*InterruptMasks, bool){
		0x07: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[3].TxLOLMask = b },
		0x06: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[2].TxLOLMask = b },
		0x05: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[1].TxLOLMask = b },
		0x04: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[0].TxLOLMask = b },
		0x03: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[3].RxLOLMask = b },
		0x02: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[2].RxLOLMask = b },
		0x01: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[1].RxLOLMask = b },
		0x00: func(i *InterruptMasks, b bool) { i.ChannelInterruptMasks[0].RxLOLMask = b },
	},
	0x03: map[uint]func(*InterruptMasks, bool){
		0x07: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.TemperatureAlarmMask.HighAlarmMask = b },
		0x06: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.TemperatureAlarmMask.LowAlarmMask = b },
		0x05: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.TemperatureAlarmMask.HighWarningMask = b },
		0x04: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.TemperatureAlarmMask.LowWarningMask = b },
		0x03: func(i *InterruptMasks, b bool) { /* reserved */ },
		0x02: func(i *InterruptMasks, b bool) { /* reserved */ },
		0x01: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.TCReadinessFlagMask = b },
		0x00: func(i *InterruptMasks, b bool) { /* reserved */ },
	},
	0x04: map[uint]func(*InterruptMasks, bool){
		0x07: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.VoltageAlarmMask.HighAlarmMask = b },
		0x06: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.VoltageAlarmMask.LowAlarmMask = b },
		0x05: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.VoltageAlarmMask.HighWarningMask = b },
		0x04: func(i *InterruptMasks, b bool) { i.FreeSideInterruptMasks.VoltageAlarmMask.LowWarningMask = b },
		0x03: func(i *InterruptMasks, b bool) { /* reserved */ },
		0x02: func(i *InterruptMasks, b bool) { /* reserved */ },
		0x01: func(i *InterruptMasks, b bool) { /* reserved */ },
		0x00: func(i *InterruptMasks, b bool) { /* reserved */ },
	},
	// 0x05-0x06 vendor specific
}

// NewInterruptMasks parses [6]byte into a new InterruptMasks instance
func NewInterruptMasks(raw [6]byte) *InterruptMasks {
	i := &InterruptMasks{}

	for byteOffset, bitMap := range interruptMasksMemoryMap {
		for bitOffset, callback := range bitMap {
			value := raw[byteOffset]&(1<<bitOffset) > 0
			callback(i, value)
		}
	}
	return i
}
