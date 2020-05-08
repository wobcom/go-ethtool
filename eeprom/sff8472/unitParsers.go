package sff8472

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math"
)

// Power a power measurement
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

func parseString(raw []byte) string {
	return string(bytes.Trim(raw, "\x00"))
}

func parseWavelength(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb))
}

func parseTemperature(msb byte, lsb byte) float64 {
	return parseSignedDecimal(msb, lsb)
}

func parseVoltage(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb)) / 10000.0
}

func parseCurrent(msb byte, lsb byte) float64 {
	return float64(parseUint16(msb, lsb)) / 500.0
}

func parsePower(msb byte, lsb byte) Power {
	return Power(float64(parseUint16(msb, lsb)) / 10000.0)
}

func parseUint16(msb byte, lsb byte) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}

func parseFloatingPoint(raw [4]byte) float64 {
	return float64(math.Float32frombits(binary.BigEndian.Uint32(raw[:])))
}

func parseSignedDecimal(msb byte, lsb byte) float64 {
	e := float64(int8(msb))
	if e > 0 {
		return e + float64(lsb)/256.0
	}
	return e - float64(lsb)/256.0
}

func parseUnsignedDecimal(msb byte, lsb byte) float64 {
	return float64(msb) + float64(lsb)/256.0
}
