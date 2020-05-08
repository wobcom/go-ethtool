package sff8472

// Thresholds Diagnostic Flag Alarm and Warning Thresholds as of SFF-84722 rev. 12.3 table 9-5
type Thresholds struct {
	Temperature *AlarmThresholds
	Voltage     *AlarmThresholds
	Bias        *AlarmThresholds
	TxPower     *AlarmThresholdsPower
	RxPower     *AlarmThresholdsPower
}

// AlarmThresholds provides high / low alarm / warning thresholds
type AlarmThresholds struct {
	HighAlarm   float64
	HighWarning float64
	LowAlarm    float64
	LowWarning  float64
}

// AlarmThresholdsPower provides high / low alarm / warning thresholds for the Power type
type AlarmThresholdsPower struct {
	HighAlarm   Power
	HighWarning Power
	LowAlarm    Power
	LowWarning  Power
}

var thresholdsMemoryMap = map[uint]func(*Thresholds, byte, byte){
	0x00: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.HighAlarm = parseTemperature(msb, lsb) },
	0x02: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.LowAlarm = parseTemperature(msb, lsb) },
	0x04: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.HighWarning = parseTemperature(msb, lsb) },
	0x06: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.LowWarning = parseTemperature(msb, lsb) },

	0x08: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.HighAlarm = parseVoltage(msb, lsb) },
	0x0A: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.LowAlarm = parseVoltage(msb, lsb) },
	0x0C: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.HighWarning = parseVoltage(msb, lsb) },
	0x0E: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.LowWarning = parseVoltage(msb, lsb) },

	0x10: func(t *Thresholds, msb byte, lsb byte) { t.Bias.HighAlarm = parseCurrent(msb, lsb) },
	0x12: func(t *Thresholds, msb byte, lsb byte) { t.Bias.LowAlarm = parseCurrent(msb, lsb) },
	0x14: func(t *Thresholds, msb byte, lsb byte) { t.Bias.HighWarning = parseCurrent(msb, lsb) },
	0x16: func(t *Thresholds, msb byte, lsb byte) { t.Bias.LowWarning = parseCurrent(msb, lsb) },

	0x18: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.HighAlarm = parsePower(msb, lsb) },
	0x1A: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.LowAlarm = parsePower(msb, lsb) },
	0x1C: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.HighWarning = parsePower(msb, lsb) },
	0x1E: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.LowWarning = parsePower(msb, lsb) },

	0x20: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.HighAlarm = parsePower(msb, lsb) },
	0x22: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.LowAlarm = parsePower(msb, lsb) },
	0x24: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.HighWarning = parsePower(msb, lsb) },
	0x26: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.LowWarning = parsePower(msb, lsb) },
}

// NewThresholds parses [40]byte into a new Thresholds instance
func NewThresholds(raw [40]byte) *Thresholds {
	t := &Thresholds{
		Temperature: &AlarmThresholds{},
		Voltage:     &AlarmThresholds{},
		Bias:        &AlarmThresholds{},
		TxPower:     &AlarmThresholdsPower{},
		RxPower:     &AlarmThresholdsPower{},
	}
	for byteOffset, callback := range thresholdsMemoryMap {
		callback(t, raw[byteOffset], raw[byteOffset+1])
	}
	return t
}
