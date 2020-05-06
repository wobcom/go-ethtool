package SFF8636

type Thresholds struct {
	Temperature AlarmThresholds
	Voltage     AlarmThresholds
	RxPower     AlarmPowerThresholds
	TxBias      AlarmThresholds
	TxPower     AlarmPowerThresholds
}

type AlarmThresholds struct {
	HighAlarm   float64
	HighWarning float64
	LowAlarm    float64
	LowWarning  float64
}

type AlarmPowerThresholds struct {
	HighAlarm   Power
	HighWarning Power
	LowAlarm    Power
	LowWarning  Power
}

var thresholdMemoryMap = map[uint]func(*Thresholds, byte, byte){
	0x00: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.HighAlarm = parseTemperature(msb, lsb) },
	0x02: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.LowAlarm = parseTemperature(msb, lsb) },
	0x04: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.HighWarning = parseTemperature(msb, lsb) },
	0x06: func(t *Thresholds, msb byte, lsb byte) { t.Temperature.LowWarning = parseTemperature(msb, lsb) },
	// 0x07-0x0F reserved
	0x10: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.HighAlarm = parseVoltage(msb, lsb) },
	0x12: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.LowAlarm = parseVoltage(msb, lsb) },
	0x14: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.HighWarning = parseVoltage(msb, lsb) },
	0x16: func(t *Thresholds, msb byte, lsb byte) { t.Voltage.LowWarning = parseVoltage(msb, lsb) },
	// 0x18-0x1F reserved
	// 0x20-0x2F vendor specific
	0x30: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.HighAlarm = parsePower(msb, lsb) },
	0x32: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.LowAlarm = parsePower(msb, lsb) },
	0x34: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.HighWarning = parsePower(msb, lsb) },
	0x36: func(t *Thresholds, msb byte, lsb byte) { t.RxPower.LowWarning = parsePower(msb, lsb) },
	0x38: func(t *Thresholds, msb byte, lsb byte) { t.TxBias.HighAlarm = parseCurrent(msb, lsb) },
	0x3A: func(t *Thresholds, msb byte, lsb byte) { t.TxBias.LowAlarm = parseCurrent(msb, lsb) },
	0x3C: func(t *Thresholds, msb byte, lsb byte) { t.TxBias.HighWarning = parseCurrent(msb, lsb) },
	0x3E: func(t *Thresholds, msb byte, lsb byte) { t.TxBias.LowWarning = parseCurrent(msb, lsb) },
	0x40: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.HighAlarm = parsePower(msb, lsb) },
	0x42: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.LowAlarm = parsePower(msb, lsb) },
	0x44: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.HighWarning = parsePower(msb, lsb) },
	0x46: func(t *Thresholds, msb byte, lsb byte) { t.TxPower.LowWarning = parsePower(msb, lsb) },
}

func NewThresholds(raw [72]byte) *Thresholds {
	t := &Thresholds{}

	for byteOffset, callback := range thresholdMemoryMap {
		callback(t, raw[byteOffset], raw[byteOffset+1])
	}

	return t
}
