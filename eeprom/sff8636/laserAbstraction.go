package sff8636

import (
	"errors"
	"github.com/wobcom/go-ethtool/eeprom"
)

// Laser a helper struct for implementing eeprom.Laser interface
type Laser struct {
	RxPower *Measurement
	TxPower *Measurement
	Bias    *Measurement
}

// Measurement a helper struct for implementing eeprom.Laser interface
type Measurement struct {
	Value               float64
	Unit                string
	ThresholdsSupported bool
	Thresholds          *MeasurementThresholds
}

// MeasurementThresholds a helper struct for implementing eeprom.Laser interface
type MeasurementThresholds struct {
	HighAlarm   float64
	HighWarning float64
	LowAlarm    float64
	LowWarning  float64
}

// SupportsMonitoring implements eeprom.Laser interface's SupportsMonitoring function
func (l *Laser) SupportsMonitoring() bool {
	return true
}

// GetBias implements eeprom.Laser interface's GetBias function
func (l *Laser) GetBias() (eeprom.Measurement, error) {
	return l.Bias, nil
}

// GetTxPower implements eeprom.Laser interface's GetTxPower function
func (l *Laser) GetTxPower() (eeprom.Measurement, error) {
	return l.TxPower, nil
}

// GetRxPower implements eeprom.Laser interface's GetRxPower function
func (l *Laser) GetRxPower() (eeprom.Measurement, error) {
	return l.RxPower, nil
}

// GetValue implements eeprom.Measurement interface's GetValue function
func (m *Measurement) GetValue() float64 {
	return m.Value
}

// GetUnit implements eeprom.Measurement interface's GetUnit function
func (m *Measurement) GetUnit() string {
	return m.Unit
}

// SupportsThresholds implements eeprom.Measurement interface's SupportsThresholds function
func (m *Measurement) SupportsThresholds() bool {
	return m.ThresholdsSupported
}

// GetAlarmThresholds implements eeprom.Measurement interface's GetAlarmThresholds function
func (m *Measurement) GetAlarmThresholds() (eeprom.AlarmThresholds, error) {
	if !m.SupportsThresholds() {
		return nil, errors.New("No thresholds implemented by this module")
	}
	return m.Thresholds, nil
}

// GetHighAlarm implements eeprom.Measurement interface's GetHighAlarm function
func (m *MeasurementThresholds) GetHighAlarm() float64 {
	return m.HighAlarm
}

// GetHighWarning implements eeprom.Measurement interface's GetHighWarning function
func (m *MeasurementThresholds) GetHighWarning() float64 {
	return m.HighWarning
}

// GetLowAlarm implements eeprom.Measurement interface's GetLowAlarm function
func (m *MeasurementThresholds) GetLowAlarm() float64 {
	return m.LowAlarm
}

// GetLowWarning implements eeprom.Measurement interface's GetLowWarning function
func (m *MeasurementThresholds) GetLowWarning() float64 {
	return m.LowWarning
}

// GetLasers implements eeprom.EEPROM interface's GetLasers function
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
