package sff8636

import (
	"bytes"
	"encoding/json"
	"math"
)

// Power type for power measurements, provides conversion to dBm when JSON serialized
type Power float64

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (p Power) MarshalJSON() ([]byte, error) {
	dbm := 10 * math.Log10(float64(p))

	if dbm == math.Inf(-1) {
		// JSON does not support -Inf, so give a rough approximation
		dbm = -math.MaxFloat64
	}
	return json.Marshal(
		map[string]float64{
			"Milliwatts":        float64(p),
			"DecibelMilliwatts": dbm,
		},
	)
}

func parseTemperature(msb byte, lsb byte) float64 {
	return float64(parseInt16(msb, lsb)) / 256.0
}

func parseVoltage(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb)) / 10000
}

func parsePower(msb byte, lsb byte) Power {
	return Power(float64(parseUint16(msb, lsb)) / 10000)
}

func parseCurrent(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb)) * 0.002
}

func parseString(raw []byte) string {
	return string(bytes.Trim(raw, "\x00"))
}

func parseWavelength(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb)) / 20
}

func parseWavelengthTolerance(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb)) / 200
}

func parseUint16(msb byte, lsb byte) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}

func parseInt16(msb byte, lsb byte) int16 {
	return int16(int16(msb)<<8) | int16(lsb)
}
