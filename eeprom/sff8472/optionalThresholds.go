package sff8472

// OptionalThresholds Thresholds for optional Laser Temperature and TEC Current alarms and warnings as of SFF-8472 rev 12.3 Table 9-5
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

// NewOptionalThresholds parses [16]byte into a new OptionalThresholds instance
func NewOptionalThresholds(raw [16]byte) *OptionalThresholds {
	o := &OptionalThresholds{}
	for byteOffset, callback := range optionalThresholdsMemoryMap {
		callback(o, raw[byteOffset], raw[byteOffset+1])
	}
	return o
}
