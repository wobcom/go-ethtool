package SFF8472

type OptionalThresholds struct {
	LaserTemperature AlarmThresholds
	TecCurrent       AlarmThresholds
}

var optionalThresholdsMemoryMap = map[uint]func(*OptionalThresholds, byte, byte){
	0x00: func(o *OptionalThresholds, msb byte, lsb byte) {
		o.LaserTemperature.HighAlarm = parseTemperature(msb, lsb)
	},
	0x02: func(o *OptionalThresholds, msb byte, lsb byte) {
		o.LaserTemperature.LowAlarm = parseTemperature(msb, lsb)
	},
	0x04: func(o *OptionalThresholds, msb byte, lsb byte) {
		o.LaserTemperature.HighWarning = parseTemperature(msb, lsb)
	},
	0x06: func(o *OptionalThresholds, msb byte, lsb byte) {
		o.LaserTemperature.HighAlarm = parseTemperature(msb, lsb)
	},

	0x08: func(o *OptionalThresholds, msb byte, lsb byte) { o.TecCurrent.HighAlarm = parseCurrent(msb, lsb) },
	0x0A: func(o *OptionalThresholds, msb byte, lsb byte) { o.TecCurrent.LowAlarm = parseCurrent(msb, lsb) },
	0x0C: func(o *OptionalThresholds, msb byte, lsb byte) { o.TecCurrent.HighWarning = parseCurrent(msb, lsb) },
	0x0E: func(o *OptionalThresholds, msb byte, lsb byte) { o.TecCurrent.HighAlarm = parseCurrent(msb, lsb) },
}

func NewOptionalThresholds(raw [16]byte) *OptionalThresholds {
	o := &OptionalThresholds{}
	for byteOffset, callback := range optionalThresholdsMemoryMap {
		callback(o, raw[byteOffset], raw[byteOffset+1])
	}
	return o
}
