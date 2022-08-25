package sff8636

import (
	"github.com/wobcom/go-ethtool/eeprom"
	"github.com/wobcom/go-ethtool/eeprom/sff8024"
	"strings"
	"time"
)

// GetIdentifier implements eeprom.EEPROM interface's GetIdentifier function
func (e *EEPROM) GetIdentifier() sff8024.Identifier {
	return e.Identifier
}

// GetConnectorType implements eeprom.EEPROM interface's GetConnectorType function
func (e *EEPROM) GetConnectorType() sff8024.ConnectorType {
	return e.ConnectorType
}

// GetEncoding implements eeprom.EEPROM interface's GetEncoding function
func (e *EEPROM) GetEncoding() string {
	return e.Encoding.String()
}

// GetPowerClass implements eeprom.EEPROM interface's GetPowerClass function
func (e *EEPROM) GetPowerClass() eeprom.PowerClass {
	return e.ExtendedIdentifier.PowerClass
}

// GetSignalingRate implements eeprom.EEPROM interface's GetSignalingRate function
func (e *EEPROM) GetSignalingRate() float64 {
	return float64(e.SignalingRate)
}

// GetSupportedLinkLengths implements eeprom.EEPROM interface's GetSupportedLinkLengths function
func (e *EEPROM) GetSupportedLinkLengths() map[string]float64 {
	if e.SpecificationCompliance.IsNonOpticalImplementation() {
		return map[string]float64{
			"copperOrDAC": float64(e.LengthOM4ActiveOrPassiveCable),
		}
	}
	return map[string]float64{
		"SMF": float64(e.LengthSMF),
		"OM1": float64(e.LengthOM1),
		"OM2": float64(e.LengthOM2),
		"OM3": float64(e.LengthOM3),
		"OM4": float64(e.LengthOM4ActiveOrPassiveCable),
	}
}

// GetVendorName implements eeprom.EEPROM interface's GetVendorName function
func (e *EEPROM) GetVendorName() string {
	return e.VendorName
}

// GetVendorPN implements eeprom.EEPROM interface's GetVendorPN function
func (e *EEPROM) GetVendorPN() string {
	return e.VendorPN
}

// GetVendorRev implements eeprom.EEPROM interface's GetVendorRev function
func (e *EEPROM) GetVendorRev() string {
	return e.VendorRev
}

// GetVendorSN implements eeprom.EEPROM interface's GetVendorSN function
func (e *EEPROM) GetVendorSN() string {
	return e.VendorSN
}

// GetVendorOUI implements eeprom.EEPROM interface's GetVendorOUI function
func (e *EEPROM) GetVendorOUI() eeprom.OUI {
	return e.VendorOUI
}

// GetDateCode implements eeprom.EEPROM interface's GetDateCode function
func (e *EEPROM) GetDateCode() time.Time {
	dateCode := strings.Trim(e.DateCode, " ")
	t, _ := time.Parse("060102", dateCode)
	return t
}

// GetWavelength implements eeprom.EEPROM interface's GetWavelength function
func (e *EEPROM) GetWavelength() float64 {
	return e.Wavelength
}

// SupportsMonitoring implements eeprom.EEPROM interface's SupportsMonitoring function
func (e *EEPROM) SupportsMonitoring() bool {
	return true
}

// GetModuleTemperature implements eeprom.EEPROM interface's GetModuleTemperature function
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

// GetModuleVoltage implements eeprom.EEPROM interface's GetModuleVoltage function
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
