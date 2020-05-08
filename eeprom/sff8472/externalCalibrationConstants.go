package sff8472

// ExternalCalibrationConstants Diagnostic calibration constants for optional External Calibration as of SFF-8472 rev 12.3 table 9-6
type ExternalCalibrationConstants struct {
	RxPwr             [5]float64
	BiasSlope         float64
	BiasOffset        float64
	TxPowerSlope      float64
	TxPowerOffset     float64
	TemperatureSlope  float64
	TemperatureOffset float64
	VoltageSlope      float64
	VoltageOffset     float64
}

var externalCalibrationConstantsMemoryMap = map[uint]func(*ExternalCalibrationConstants, byte, byte){
	// 0x00-0x13 handled extra
	0x14: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.BiasSlope = parseUnsignedDecimal(msb, lsb)
	},
	0x16: func(e *ExternalCalibrationConstants, msb byte, lsb byte) { e.BiasOffset = parseSignedDecimal(msb, lsb) },
	0x18: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.TxPowerSlope = parseUnsignedDecimal(msb, lsb)
	},
	0x1A: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.TxPowerOffset = parseSignedDecimal(msb, lsb)
	},
	0x1C: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.TemperatureSlope = parseUnsignedDecimal(msb, lsb)
	},
	0x1E: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.TemperatureOffset = parseSignedDecimal(msb, lsb)
	},
	0x20: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.VoltageSlope = parseUnsignedDecimal(msb, lsb)
	},
	0x22: func(e *ExternalCalibrationConstants, msb byte, lsb byte) {
		e.VoltageOffset = parseSignedDecimal(msb, lsb)
	},
}

// NewExternalCalibrationConstants parses [36]byte into a new ExternalCalibrationConstants instance
func NewExternalCalibrationConstants(raw [36]byte) *ExternalCalibrationConstants {
	e := &ExternalCalibrationConstants{}
	for i := 0; i < 5; i++ {
		e.RxPwr[i] = parseFloatingPoint([4]byte{
			raw[(i*4)+0],
			raw[(i*4)+1],
			raw[(i*4)+2],
			raw[(i*4)+3],
		})
	}
	for byteOffset, callback := range externalCalibrationConstantsMemoryMap {
		callback(e, raw[byteOffset], raw[byteOffset+1])
	}
	return e
}
