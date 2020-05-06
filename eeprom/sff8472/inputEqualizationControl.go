package SFF8472

import (
	"encoding/json"
	"fmt"
)

type InputEqualizationControl struct {
	HighRate InputEqualization
	LowRate  InputEqualization
}
type InputEqualization byte

const (
	InputEqualization10dB InputEqualization = 10
	InputEqualization9dB  InputEqualization = 9
	InputEqualization8dB  InputEqualization = 8
	InputEqualization7dB  InputEqualization = 7
	InputEqualization6dB  InputEqualization = 6
	InputEqualization5dB  InputEqualization = 5
	InputEqualization4dB  InputEqualization = 4
	InputEqualization3dB  InputEqualization = 3
	InputEqualization2dB  InputEqualization = 2
	InputEqualization1dB  InputEqualization = 1
	InputEqualizationNoEQ InputEqualization = 0
)

func (i InputEqualization) String() string {
	if byte(i) > 0 && byte(i) <= 10 {
		return fmt.Sprintf("%d dB", byte(i))
	} else if byte(i) == 0 {
		return "No EQ"
	}
	return "Reserved"
}

func (i InputEqualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": i.String(),
		"hex":   fmt.Sprintf("%#02X", byte(i)),
	})
}

func NewInputEqualizationControl(raw byte) *InputEqualizationControl {
	return &InputEqualizationControl{
		HighRate: InputEqualization((raw & 0b11110000) >> 4),
		LowRate:  InputEqualization(raw & 0b00001111),
	}
}
