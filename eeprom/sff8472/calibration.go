package SFF8472

func (e *EEPROM) calibrate() {
	e.Diagnostics = e.Diagnostics.calibrate(e.ExternalCalibrationConstants)
	e.Thresholds = e.Thresholds.calibrate(e.ExternalCalibrationConstants)
}

func (d *Diagnostics) calibrate(e *ExternalCalibrationConstants) *Diagnostics {
	return &Diagnostics{
		Temperature: calibrateValue(d.Temperature, e.TemperatureSlope, e.TemperatureOffset),
		Voltage:     calibrateValue(d.Voltage, e.VoltageSlope, e.VoltageOffset),
		Bias:        calibrateValue(d.Bias, e.BiasSlope, e.BiasOffset),
		TxPower:     calibrateTxPower(d.TxPower, e.TxPowerSlope, e.TxPowerOffset),
		RxPower:     calibrateRxPower(d.RxPower, e.RxPwr),
	}
}

func (t *Thresholds) calibrate(e *ExternalCalibrationConstants) *Thresholds {
	return &Thresholds{
		Temperature: t.Temperature.calibrate(e.TemperatureSlope, e.TemperatureOffset),
		Voltage:     t.Voltage.calibrate(e.TemperatureSlope, e.TemperatureOffset),
		Bias:        t.Bias.calibrate(e.TemperatureSlope, e.TemperatureOffset),
		TxPower:     t.TxPower.calibrateTxPower(e.TxPowerSlope, e.TxPowerOffset),
		RxPower:     t.RxPower.calibrateRxPower(e.RxPwr),
	}
}

func (a *AlarmThresholds) calibrate(slope float64, offset float64) *AlarmThresholds {
	return &AlarmThresholds{
		HighAlarm:   calibrateValue(a.HighAlarm, slope, offset),
		HighWarning: calibrateValue(a.HighWarning, slope, offset),
		LowAlarm:    calibrateValue(a.LowAlarm, slope, offset),
		LowWarning:  calibrateValue(a.LowWarning, slope, offset),
	}
}

func (a *AlarmThresholdsPower) calibrateTxPower(slope float64, offset float64) *AlarmThresholdsPower {
	return &AlarmThresholdsPower{
		HighAlarm:   calibrateTxPower(a.HighAlarm, slope, offset),
		HighWarning: calibrateTxPower(a.HighWarning, slope, offset),
		LowAlarm:    calibrateTxPower(a.LowAlarm, slope, offset),
		LowWarning:  calibrateTxPower(a.LowWarning, slope, offset),
	}
}

func (a *AlarmThresholdsPower) calibrateRxPower(rxPwr [5]float64) *AlarmThresholdsPower {
	return &AlarmThresholdsPower{
		HighAlarm:   calibrateRxPower(a.HighAlarm, rxPwr),
		HighWarning: calibrateRxPower(a.HighWarning, rxPwr),
		LowAlarm:    calibrateRxPower(a.LowAlarm, rxPwr),
		LowWarning:  calibrateRxPower(a.LowWarning, rxPwr),
	}
}

// implements formulas 1-4 given in SFF-8472 Rev 12.3, Section 9.3
// used for calibrating temperature, voltage, current and TxPower
func calibrateValue(uncalibratedValue float64, slope float64, offset float64) float64 {
	return uncalibratedValue*slope + offset
}

func calibrateTxPower(uncalibratedValue Power, slope float64, offset float64) Power {
	return Power(calibrateValue(float64(uncalibratedValue), slope, offset))
}

// implements formula 5 given in SFF-8472 Rev 12.3, Section 9.3
// used for calibrating RxPower
func calibrateRxPower(uncalibratedValue Power, rxPwr [5]float64) Power {
	rxPwrAD := float64(uncalibratedValue)
	return Power(rxPwr[4]*rxPwrAD + rxPwr[3]*rxPwrAD + rxPwr[2]*rxPwrAD + rxPwr[1]*rxPwrAD + rxPwr[0])
}
