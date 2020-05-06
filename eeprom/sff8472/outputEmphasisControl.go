package SFF8472

import (
	"encoding/json"
	"fmt"
)

type OutputEmphasisControl struct {
	HighRate OutputEmphasis
	LowRate  OutputEmphasis
}

type OutputEmphasis byte

const (
	OutputEmphasis7dB        OutputEmphasis = 7
	OutputEmphasis6dB        OutputEmphasis = 6
	OutputEmphasis5dB        OutputEmphasis = 5
	OutputEmphasis4dB        OutputEmphasis = 4
	OutputEmphasis3dB        OutputEmphasis = 3
	OutputEmphasis2dB        OutputEmphasis = 2
	OutputEmphasis1dB        OutputEmphasis = 1
	OutputEmphasisNoEmphasis OutputEmphasis = 0
)

func (o OutputEmphasis) String() string {
	if byte(o) > 0 && byte(o) <= 7 {
		return fmt.Sprintf("%d dB", byte(o))
	} else if byte(o) == 0 {
		return "No Emphasis"
	}
	return "Vendor specific"
}

func (o OutputEmphasis) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": o.String(),
		"hex":   fmt.Sprintf("%#02X", byte(o)),
	})
}

func NewOutputEmphasisControl(raw byte) *OutputEmphasisControl {
	return &OutputEmphasisControl{
		HighRate: OutputEmphasis((raw & 0b11110000) >> 4),
		LowRate:  OutputEmphasis(raw & 0b00001111),
	}
}
