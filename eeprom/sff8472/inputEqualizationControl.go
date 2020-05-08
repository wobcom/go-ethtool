package sff8472

import (
	"encoding/json"
	"fmt"
)

// InputEqualizationControl as of SFF-8472 rev 12.3 table 9-12
type InputEqualizationControl struct {
	HighRate InputEqualization
	LowRate  InputEqualization
}
// InputEqualization as of SFF-8472 rev 12.3 table 9-12
type InputEqualization byte

const (
    // InputEqualization10dB 10dB
	InputEqualization10dB InputEqualization = 10
    // InputEqualization9dB 9dB
	InputEqualization9dB  InputEqualization = 9
    // InputEqualization8dB 8dB
	InputEqualization8dB  InputEqualization = 8
    // InputEqualization7dB 7dB
	InputEqualization7dB  InputEqualization = 7
    // InputEqualization6dB 6dB
	InputEqualization6dB  InputEqualization = 6
    // InputEqualization5dB 5dB
	InputEqualization5dB  InputEqualization = 5
    // InputEqualization4dB 4dB
	InputEqualization4dB  InputEqualization = 4
    // InputEqualization3dB 3dB
	InputEqualization3dB  InputEqualization = 3
    // InputEqualization2dB 2dB
	InputEqualization2dB  InputEqualization = 2
    // InputEqualization1dB 1dB
	InputEqualization1dB  InputEqualization = 1
    // InputEqualizationNoEQ No EQ
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

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (i InputEqualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": i.String(),
		"hex":   fmt.Sprintf("%#02X", byte(i)),
	})
}

// NewInputEqualizationControl parses a byte into a new InputEqualizationControl instance
func NewInputEqualizationControl(raw byte) *InputEqualizationControl {
	return &InputEqualizationControl{
		HighRate: InputEqualization((raw & 0b11110000) >> 4),
		LowRate:  InputEqualization(raw & 0b00001111),
	}
}
