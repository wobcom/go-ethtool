package SFF8472

import (
	"errors"
	"gitlab.com/wobcom/golang-ethtool/eeprom"
	"gitlab.com/wobcom/golang-ethtool/eeprom/sff8024"
	"strings"
	"time"
)

func (e *EEPROM) GetIdentifier() SFF8024.Identifier {
	return e.Identifier
}

func (e *EEPROM) GetConnectorType() SFF8024.ConnectorType {
	return e.ConnectorType
}

func (e *EEPROM) GetEncoding() eeprom.Encoding {
	return e.Encoding
}

func (e *EEPROM) GetPowerClass() eeprom.PowerClass {
	return e.Options.GetPowerClass()
}

func (e *EEPROM) GetSignalingRate() float64 {
	return float64(e.SignalingRate)
}

func (e *EEPROM) GetSupportedLinkLengths() map[string]float64 {
	if e.TransceiverCompliance.IsSFPCableImplementation() {
		return map[string]float64{
			"copperOrDAC": e.LengthOM4OrDAC,
		}
	}
	return map[string]float64{
		"SMF(km)": e.LengthSMFKm,
		"SMF(m)":  e.LengthSMF,
		"OM1":     e.LengthOM1,
		"OM2":     e.LengthOM2,
		"OM3":     e.LengthOM3,
		"OM4":     e.LengthOM4OrDAC,
	}
}

func (e *EEPROM) GetVendorName() string {
	return e.VendorName
}

func (e *EEPROM) GetVendorPN() string {
	return e.VendorPN
}

func (e *EEPROM) GetVendorRev() string {
	return e.VendorRev
}

func (e *EEPROM) GetVendorSN() string {
	return e.VendorSN
}

func (e *EEPROM) GetVendorOUI() eeprom.OUI {
	return e.VendorOUI
}

func (e *EEPROM) GetDateCode() time.Time {
	dateCode := strings.Trim(e.DateCode, " ")
	t, _ := time.Parse("060102", dateCode)
	return t
}

func (e *EEPROM) GetWavelength() float64 {
	return e.Wavelength
}

func (e *EEPROM) SupportsMonitoring() bool {
	return e.DiagnosticMonitoringType.DiagnosticMonitoringImplemented
}

func (e *EEPROM) GetModuleTemperature() (eeprom.Measurement, error) {
	if !e.SupportsMonitoring() {
		return nil, errors.New("Monitoring not implemented by module")
	}
	return &Measurement{
		Value: e.Diagnostics.Temperature,
		Unit:  "degrees celsius",
		Thresholds: &MeasurementThresholds{
			HighAlarm:   e.Thresholds.Temperature.HighAlarm,
			HighWarning: e.Thresholds.Temperature.HighWarning,
			LowAlarm:    e.Thresholds.Temperature.LowAlarm,
			LowWarning:  e.Thresholds.Temperature.LowWarning,
		},
	}, nil
}

func (e *EEPROM) GetModuleVoltage() (eeprom.Measurement, error) {
	if !e.SupportsMonitoring() {
		return nil, errors.New("Monitoring not implemented by module")
	}
	return &Measurement{
		Value: e.Diagnostics.Voltage,
		Unit:  "degrees celsius",
		Thresholds: &MeasurementThresholds{
			HighAlarm:   e.Thresholds.Voltage.HighAlarm,
			HighWarning: e.Thresholds.Voltage.HighWarning,
			LowAlarm:    e.Thresholds.Voltage.LowAlarm,
			LowWarning:  e.Thresholds.Voltage.LowWarning,
		},
	}, nil
}
