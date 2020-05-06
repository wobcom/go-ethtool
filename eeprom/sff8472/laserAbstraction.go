package SFF8472

import (
	"errors"
	"gitlab.com/wobcom/ethtool/eeprom"
)

type Laser struct {
	RxPower             *Measurement
	TxPower             *Measurement
	Bias                *Measurement
	MonitoringSupported bool
}

type Measurement struct {
	Value      float64
	Unit       string
	Thresholds *MeasurementThresholds
}

type MeasurementThresholds struct {
	HighAlarm   float64
	HighWarning float64
	LowAlarm    float64
	LowWarning  float64
}

func (l *Laser) SupportsMonitoring() bool {
	return l.MonitoringSupported
}

func (l *Laser) GetBias() (eeprom.Measurement, error) {
	if !l.SupportsMonitoring() {
		return nil, errors.New("This module does not implement monitoring")
	}
	return l.Bias, nil
}

func (l *Laser) GetTxPower() (eeprom.Measurement, error) {
	if !l.SupportsMonitoring() {
		return nil, errors.New("This module does not implement monitoring")
	}
	return l.TxPower, nil
}

func (l *Laser) GetRxPower() (eeprom.Measurement, error) {
	if !l.SupportsMonitoring() {
		return nil, errors.New("This module does not implement monitoring")
	}
	return l.RxPower, nil
}

func (m *Measurement) GetValue() float64 {
	return m.Value
}

func (m *Measurement) GetUnit() string {
	return m.Unit
}

func (m *Measurement) SupportsThresholds() bool {
	return true
}

func (m *Measurement) GetAlarmThresholds() (eeprom.AlarmThresholds, error) {
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
	if e.TransceiverCompliance.IsSFPCableImplementation() {
		return []eeprom.Laser{}
	}
	return []eeprom.Laser{
		&Laser{
			RxPower: &Measurement{
				Value: float64(e.Diagnostics.RxPower),
				Unit:  "milliwatts",
				Thresholds: &MeasurementThresholds{
					HighAlarm:   float64(e.Thresholds.RxPower.HighAlarm),
					HighWarning: float64(e.Thresholds.RxPower.HighWarning),
					LowAlarm:    float64(e.Thresholds.RxPower.LowAlarm),
					LowWarning:  float64(e.Thresholds.RxPower.LowWarning),
				},
			},
			TxPower: &Measurement{
				Value: float64(e.Diagnostics.TxPower),
				Unit:  "milliwatts",
				Thresholds: &MeasurementThresholds{
					HighAlarm:   float64(e.Thresholds.TxPower.HighAlarm),
					HighWarning: float64(e.Thresholds.TxPower.HighWarning),
					LowAlarm:    float64(e.Thresholds.TxPower.LowAlarm),
					LowWarning:  float64(e.Thresholds.TxPower.LowWarning),
				},
			},
			Bias: &Measurement{
				Value: float64(e.Diagnostics.Bias),
				Unit:  "milliwatts",
				Thresholds: &MeasurementThresholds{
					HighAlarm:   e.Thresholds.Bias.HighAlarm,
					HighWarning: e.Thresholds.Bias.HighWarning,
					LowAlarm:    e.Thresholds.Bias.LowAlarm,
					LowWarning:  e.Thresholds.Bias.LowWarning,
				},
			},
		},
	}
}
