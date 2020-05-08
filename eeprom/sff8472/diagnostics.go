package sff8472

// Diagnostics as of SFF-8472
type Diagnostics struct {
	Temperature float64
	Voltage     float64
	Bias        float64
	TxPower     Power
	RxPower     Power
}

var diagnosticsMemoryMap = map[uint]func(*Diagnostics, byte, byte){
	0x00: func(d *Diagnostics, msb byte, lsb byte) { d.Temperature = parseTemperature(msb, lsb) },
	0x02: func(d *Diagnostics, msb byte, lsb byte) { d.Voltage = parseVoltage(msb, lsb) },
	0x04: func(d *Diagnostics, msb byte, lsb byte) { d.Bias = parseCurrent(msb, lsb) },
	0x06: func(d *Diagnostics, msb byte, lsb byte) { d.TxPower = parsePower(msb, lsb) },
	0x08: func(d *Diagnostics, msb byte, lsb byte) { d.RxPower = parsePower(msb, lsb) },
}

// NewDiagnostics parses [10]byte into a new instance of Diagnostics
func NewDiagnostics(raw [10]byte) *Diagnostics {
	d := &Diagnostics{}
	for byteOffset, callback := range diagnosticsMemoryMap {
		callback(d, raw[byteOffset], raw[byteOffset+1])
	}
	return d
}
