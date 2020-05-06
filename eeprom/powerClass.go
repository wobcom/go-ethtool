package eeprom

import (
	"fmt"
)

type PowerClass byte

const (
	_ PowerClass = iota
	PWR_CLASS_1
	PWR_CLASS_2
	PWR_CLASS_3
	PWR_CLASS_4
	PWR_CLASS_5
	PWR_CLASS_6
	PWR_CLASS_7
)

func (p PowerClass) GetMaxPower() float64 {
	return map[PowerClass]float64{
		PWR_CLASS_1: 1.5,
		PWR_CLASS_2: 2.0,
		PWR_CLASS_3: 2.5,
		PWR_CLASS_4: 3.5,
		PWR_CLASS_5: 4.0,
		PWR_CLASS_6: 4.5,
		PWR_CLASS_7: 5.0,
	}[p]
}

func (p PowerClass) String() string {
	return fmt.Sprintf("Power Level %d (max %.2f W)", byte(p), p.GetMaxPower())
}

func (p PowerClass) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"maxPowerWatts": p.GetMaxPower(),
		"powerLevel":    byte(p) + 1,
	}
}
