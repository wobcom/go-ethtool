package SFF8636

import (
	"errors"
	"gitlab.com/wobcom/ethtool/eeprom"
)

type Laser struct {
	RxPower *Measurement
	TxPower *Measurement
	Bias    *Measurement
}

type Measurement struct {
	Value               float64
	Unit                string
	ThresholdsSupported bool
	Thresholds          *MeasurementThresholds
}

type MeasurementThresholds struct {
	HighAlarm   float64
	HighWarning float64
	LowAlarm    float64
	LowWarning  float64
}

func (l *Laser) SupportsMonitoring() bool {
	return true
}

func (l *Laser) GetBias() (eeprom.Measurement, error) {
	return l.Bias, nil
}

func (l *Laser) GetTxPower() (eeprom.Measurement, error) {
	return l.TxPower, nil
}

func (l *Laser) GetRxPower() (eeprom.Measurement, error) {
	return l.RxPower, nil
}

func (m *Measurement) GetValue() float64 {
	return m.Value
}

func (m *Measurement) GetUnit() string {
	return m.Unit
}

func (m *Measurement) SupportsThresholds() bool {
	return m.ThresholdsSupported
}

func (m *Measurement) GetAlarmThresholds() (eeprom.AlarmThresholds, error) {
	if !m.SupportsThresholds() {
		return nil, errors.New("No thresholds implemented by this module")
	}
	return m.Thresholds, nil
}

func (m *MeasurementThresholds) GetHighAlarm() float64 {
	return m.HighAlarm
}

func (m *MeasurementThresholds) GetHighWarning() float64 {
	return m.HighWarning
}

func (m *MeasurementThresholds) GetLowAlarm() float64 {
	return m.LowAlarm
}

func (m *MeasurementThresholds) GetLowWarning() float64 {
	return m.LowWarning
}

func (e *EEPROM) GetLasers() []eeprom.Laser {
	if e.SpecificationCompliance.IsNonOpticalImplementation() {
		return []eeprom.Laser{}
	}
	ret := []eeprom.Laser{}

	for i := 0; i < 4; i++ {
		laser := &Laser{
			RxPower: &Measurement{
				Value:               float64(e.ChannelMonitors[i].RxPower),
				ThresholdsSupported: e.Thresholds != nil,
				Unit:                "milliwatts",
			},
			TxPower: &Measurement{
				Value:               float64(e.ChannelMonitors[i].TxPower),
				ThresholdsSupported: e.Thresholds != nil,
				Unit:                "miliwatts",
			},
			Bias: &Measurement{
				Value:               e.ChannelMonitors[i].Bias,
				ThresholdsSupported: e.Thresholds != nil,
				Unit:                "milliamperes",
			},
		}

		if e.Thresholds != nil {
			laser.RxPower.Thresholds = &MeasurementThresholds{
				HighAlarm:   float64(e.Thresholds.RxPower.HighAlarm),
				HighWarning: float64(e.Thresholds.RxPower.HighWarning),
				LowAlarm:    float64(e.Thresholds.RxPower.LowAlarm),
				LowWarning:  float64(e.Thresholds.RxPower.LowWarning),
			}
			laser.TxPower.Thresholds = &MeasurementThresholds{
				HighAlarm:   float64(e.Thresholds.TxPower.HighAlarm),
				HighWarning: float64(e.Thresholds.TxPower.HighWarning),
				LowAlarm:    float64(e.Thresholds.TxPower.LowAlarm),
				LowWarning:  float64(e.Thresholds.TxPower.LowWarning),
			}
			laser.Bias.Thresholds = &MeasurementThresholds{
				HighAlarm:   e.Thresholds.TxBias.HighAlarm,
				HighWarning: e.Thresholds.TxBias.HighWarning,
				LowAlarm:    e.Thresholds.TxBias.LowAlarm,
				LowWarning:  e.Thresholds.TxBias.LowWarning,
			}
		}

		ret = append(ret, laser)
	}
	return ret
}
