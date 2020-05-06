package SFF8636

import (
	"gitlab.com/wobcom/ethtool/eeprom"
	"gitlab.com/wobcom/ethtool/eeprom/sff8024"
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
	return e.ExtendedIdentifier.PowerClass
}

func (e *EEPROM) GetSignalingRate() float64 {
	return float64(e.SignalingRate)
}

func (e *EEPROM) GetSupportedLinkLengths() map[string]float64 {
	if e.SpecificationCompliance.IsNonOpticalImplementation() {
		return map[string]float64{
			"copperOrDAC": float64(e.LengthOM4ActiveOrPassiveCable),
		}
	}
	return map[string]float64{
		"SMF(km)": float64(e.LengthSMF),
		"OM1":     float64(e.LengthOM1),
		"OM2":     float64(e.LengthOM2),
		"OM3":     float64(e.LengthOM3),
		"OM4":     float64(e.LengthOM4ActiveOrPassiveCable),
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
	return true
}

func (e *EEPROM) GetModuleTemperature() (eeprom.Measurement, error) {
	m := &Measurement{
		Value:               e.FreeSideMonitors.Temperature,
		Unit:                "degrees celsius",
		ThresholdsSupported: e.Thresholds != nil,
	}

	if e.Thresholds != nil {
		m.Thresholds = &MeasurementThresholds{
			HighAlarm:   e.Thresholds.Temperature.HighAlarm,
			HighWarning: e.Thresholds.Temperature.HighWarning,
			LowAlarm:    e.Thresholds.Temperature.LowAlarm,
			LowWarning:  e.Thresholds.Temperature.LowWarning,
		}
	}
	return m, nil
}

func (e *EEPROM) GetModuleVoltage() (eeprom.Measurement, error) {
	m := &Measurement{
		Value:               e.FreeSideMonitors.SupplyVoltage,
		Unit:                "volts",
		ThresholdsSupported: e.Thresholds != nil,
	}

	if e.Thresholds != nil {
		m.Thresholds = &MeasurementThresholds{
			HighAlarm:   e.Thresholds.Voltage.HighAlarm,
			HighWarning: e.Thresholds.Voltage.HighWarning,
			LowAlarm:    e.Thresholds.Voltage.LowAlarm,
			LowWarning:  e.Thresholds.Voltage.LowWarning,
		}
	}
	return m, nil
}
