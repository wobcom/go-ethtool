package eeprom

import (
    "encoding/json"
	"fmt"
)

// PowerClass laser maximum power specification
type PowerClass byte

const (
	_ PowerClass = iota
    // PowerClass1 up to 1.5 Watts
	PowerClass1
    // PowerClass2 up to 2.0 Watts
	PowerClass2
    // PowerClass3 up to 2.5 Watts
	PowerClass3
    // PowerClass4 up to 3.5 Watts
	PowerClass4
    // PowerClass5 up to 4.0 Watts
	PowerClass5
    // PowerClass6 up to 4.5 Watts
	PowerClass6
    // PowerClass7 up to 5.0 Watts
	PowerClass7
)

// GetMaxPower returns the maximum power in watts for a given PowerClass
func (p PowerClass) GetMaxPower() float64 {
	return map[PowerClass]float64{
		PowerClass1: 1.5,
		PowerClass2: 2.0,
		PowerClass3: 2.5,
		PowerClass4: 3.5,
		PowerClass5: 4.0,
		PowerClass6: 4.5,
		PowerClass7: 5.0,
	}[p]
}

func (p PowerClass) String() string {
	return fmt.Sprintf("Power Level %d (max %.2f W)", byte(p), p.GetMaxPower())
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (p PowerClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"maxPowerWatts": p.GetMaxPower(),
		"powerLevel":    byte(p) + 1,
	})
}
