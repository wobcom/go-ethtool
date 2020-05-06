package SFF8472

type AlarmFlags struct {
	Temperature      AlarmFlagStatus
	Voltage          AlarmFlagStatus
	Bias             AlarmFlagStatus
	TxPower          AlarmFlagStatus
	RxPower          AlarmFlagStatus
	LaserTemperature AlarmFlagStatus
	TecCurrent       AlarmFlagStatus
}

type AlarmFlagStatus struct {
	HighAlarm bool
	LowAlarm  bool
}

var alarmFlagsMemoryMap = map[uint]map[uint]func(*AlarmFlags, bool){
	0x00: map[uint]func(*AlarmFlags, bool){
		0x07: func(a *AlarmFlags, b bool) { a.Temperature.HighAlarm = b },
		0x06: func(a *AlarmFlags, b bool) { a.Temperature.LowAlarm = b },
		0x05: func(a *AlarmFlags, b bool) { a.Voltage.HighAlarm = b },
		0x04: func(a *AlarmFlags, b bool) { a.Voltage.LowAlarm = b },
		0x03: func(a *AlarmFlags, b bool) { a.Bias.HighAlarm = b },
		0x02: func(a *AlarmFlags, b bool) { a.Bias.LowAlarm = b },
		0x01: func(a *AlarmFlags, b bool) { a.TxPower.HighAlarm = b },
		0x00: func(a *AlarmFlags, b bool) { a.TxPower.LowAlarm = b },
	},
	0x01: map[uint]func(*AlarmFlags, bool){
		0x07: func(a *AlarmFlags, b bool) { a.RxPower.HighAlarm = b },
		0x06: func(a *AlarmFlags, b bool) { a.RxPower.LowAlarm = b },
		0x05: func(a *AlarmFlags, b bool) { a.LaserTemperature.HighAlarm = b },
		0x04: func(a *AlarmFlags, b bool) { a.LaserTemperature.LowAlarm = b },
		0x03: func(a *AlarmFlags, b bool) { a.TecCurrent.HighAlarm = b },
		0x02: func(a *AlarmFlags, b bool) { a.TecCurrent.LowAlarm = b },
	},
}

func NewAlarmFlags(raw [2]byte) *AlarmFlags {
	a := &AlarmFlags{}
	for byteOffset, bitMap := range alarmFlagsMemoryMap {
		for bitOffset, callback := range bitMap {
			callback(a, raw[byteOffset]&(1<<bitOffset) > 0)
		}
	}
	return a
}
