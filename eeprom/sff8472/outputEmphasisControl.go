package sff8472

import (
	"encoding/json"
	"fmt"
)

// OutputEmphasisControl as of SFF-8472 rev 12.3 table 9-12
type OutputEmphasisControl struct {
	HighRate OutputEmphasis
	LowRate  OutputEmphasis
}

// OutputEmphasis Output emphasis level
type OutputEmphasis byte

const (
	// OutputEmphasis7dB 7dB
	OutputEmphasis7dB OutputEmphasis = 7
	// OutputEmphasis6dB 6dB
	OutputEmphasis6dB OutputEmphasis = 6
	// OutputEmphasis5dB 5dB
	OutputEmphasis5dB OutputEmphasis = 5
	// OutputEmphasis4dB 4dB
	OutputEmphasis4dB OutputEmphasis = 4
	// OutputEmphasis3dB 3dB
	OutputEmphasis3dB OutputEmphasis = 3
	// OutputEmphasis2dB 2dB
	OutputEmphasis2dB OutputEmphasis = 2
	// OutputEmphasis1dB 1dB
	OutputEmphasis1dB OutputEmphasis = 1
	// OutputEmphasisNoEmphasis No Emphasis
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

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (o OutputEmphasis) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": o.String(),
		"hex":   fmt.Sprintf("%#02X", byte(o)),
	})
}

// NewOutputEmphasisControl parses a byte into a new OutputEmphasisControl instance
func NewOutputEmphasisControl(raw byte) *OutputEmphasisControl {
	return &OutputEmphasisControl{
		HighRate: OutputEmphasis((raw & 0b11110000) >> 4),
		LowRate:  OutputEmphasis(raw & 0b00001111),
	}
}
