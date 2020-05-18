package sff8472

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func getEEPROMfromHex(t *testing.T, hexRaw string) *EEPROM {
	rawData, err := hex.DecodeString(hexRaw)
	if err != nil {
		t.Errorf("Decode hex failed: %v", err)
	}

	eeprom, err := NewEEPROM(rawData)
	if err != nil {
		t.Errorf("EEPROM decode failed: %v", err)
	}

	return eeprom
}

func getEEPROM(t *testing.T) *EEPROM {
	hexRaw := "03040720000000000000000667000a6400000000464c45584f5054495820202020202020000002c9502e42313639362e31302e44412020204120202005320087001a000046373942354b48202020202020202020313931323138202068b0038b0000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff46373942354b4820202024f3bb1f3ab6709be265b13a5c2089d85a00e7005500ec008ca0753088b87724afc803e89c4005dc457707cb372d09d0312d00fb2bd4013c000000000000000000000000000000000000000000000000000000003f800000000000000100000001000000010000000100000000000070198b7f21367018d9158e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	return getEEPROMfromHex(t, hexRaw)
}

func getEEPROM1(t *testing.T) *EEPROM {
	hexRaw := "03040720000000000000000667000a64000000004f454d202020202020202020202020200000176a5346502d313047422d4c52202020202042342020051e004f001a000059414f59313036302020202020202020313530373033202068f003ee000006230d3855ec5278f2d03b4f771b81ee4000000000000000000009af5848ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5a00f6005500fb008dcc740487287a44c35003e89c4007d0349804f32bd406304e20009e2710013cffffffffffffffffffffffffffffffff0000000000000000000000003f8000000000000001000000010000000100000001000000ffffffb7228a807d464d1433000000000000020000400000004000000000000000000000434f554941523543414131302d323435372d30315630312001004600000000c9000000000000000000000000000000000000000000000000000000000000aaaa5346502d3130472d4c522020202020202020202031300000000000000000002a1e282e2e3134293600000000000000000000000000660000ffffffffffffffff"
	return getEEPROMfromHex(t, hexRaw)
}

func getEEPROM2(t *testing.T) *EEPROM {
	hexRaw := "0304000000000800000000010d0000000000640042524f434144452020202020202020200000051e35372d313030303034322d30322020202020202000000013001200004637385456554d2020202020202020203138313132312020000000a10000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4637385456554d202020202020202020a7c90f6053fefc039588e656bbe938ab"
	return getEEPROMfromHex(t, hexRaw)
}

func assertString(t *testing.T, got string, expected string, function string) {
	if got != expected {
		t.Errorf("%s returned '%s', but expected '%s'", function, got, expected)
	}
}
func assertFloat64(t *testing.T, got float64, expected float64, function string) {
	if got != expected {
		t.Errorf("%s returned %f, but expected %f", function, got, expected)
	}
}
func assertBool(t *testing.T, got bool, expected bool, function string) {
	if got != expected {
		t.Errorf("%s returned %t, but expected %t", function, got, expected)
	}
}
func assertInt(t *testing.T, got int, expected int, function string) {
	if got != expected {
		t.Errorf("%s returned %d, but expected %d", function, got, expected)
	}
}

func TestParseEEPROM(t *testing.T) {
	eeprom := getEEPROM(t)

	assertString(t, eeprom.GetIdentifier().String(), "SFP", "eeprom.GetIdentifier")
	assertString(t, eeprom.GetConnectorType().String(), "LC", "eeprom.GetConnectorType")
	assertString(t, eeprom.GetEncoding(), "64B/66B", "eeprom.GetEncoding")
	assertString(t, eeprom.GetPowerClass().String(), "Power Level 1 (max 1.50 W)", "eeprom.GetPowerClass")
	assertFloat64(t, eeprom.GetSignalingRate(), 10300000000.0, "eeprom.GetSignalingRate")
	assertString(t, eeprom.GetVendorName(), "FLEXOPTIX", "eeprom.GetVendorName")
	assertString(t, eeprom.GetVendorPN(), "P.B1696.10.DA", "eeprom.GetVendorPN")
	assertString(t, eeprom.GetVendorRev(), "A", "eeprom.GetVendorRev")
	assertString(t, eeprom.GetVendorSN(), "F79B5KH", "eeprom.GetVendorSN")
	assertString(t, eeprom.GetVendorOUI().String(), "20:00:00", "eeprom.GetVendorOUI")
	expectedDate, _ := time.Parse("060102", "191218")
	if dateCode := eeprom.GetDateCode(); dateCode != expectedDate {
		t.Errorf("eeprom.GetDateCode() returned %s, but expected %s", dateCode.Format("060102"), expectedDate.Format("060102"))
	}
	assertFloat64(t, eeprom.GetWavelength(), 1330, "eeprom.GetWavelength")
	lasers := eeprom.GetLasers()
	if len(lasers) != 1 {
		t.Errorf("Expected exactly one laser")
	}
	laser := lasers[0]
	assertBool(t, laser.SupportsMonitoring(), true, "laser.SupportsMonitoring")
	txPower, err := laser.GetTxPower()
	if err != nil {
		t.Errorf(err.Error())
	}
	assertBool(t, txPower.SupportsThresholds(), true, "txPower.SupportsThresholds")
	txPowerThresholds, err := txPower.GetAlarmThresholds()
	if err != nil {
		t.Errorf(err.Error())
	}
	assertFloat64(t, txPowerThresholds.GetHighAlarm(), 1.7783, "txPowerThresholds.GetHighAlarm")
}

func TestParseEEPROM1(t *testing.T) {
	eeprom := getEEPROM1(t)
	linkLengths := eeprom.GetSupportedLinkLengths()
	assertFloat64(t, linkLengths["SMF(km)"], 10, "linkLengths[\"SMF(km)\"]")
	assertFloat64(t, linkLengths["SMF(m)"], 10000, "linkLengths[\"SMF(m)\"]")

	assertBool(t, eeprom.SupportsMonitoring(), true, "eeprom.SupportsMonitoring")

	moduleTemp, err := eeprom.GetModuleTemperature()
	if err != nil {
		t.Error(err.Error())
	}
	assertString(t, fmt.Sprintf("%.2f", moduleTemp.GetValue()), "34.54", "moduleTemp.GetValue")

	moduleVoltage, err := eeprom.GetModuleVoltage()
	if err != nil {
		t.Error(err.Error())
	}
	assertString(t, fmt.Sprintf("%.2f", moduleVoltage.GetValue()), "3.29", "moduleVoltage.GetValue")

	lasers := eeprom.GetLasers()
	assertInt(t, len(lasers), 1, "len(lasers)")
	laser := lasers[0]
	bias, err := laser.GetBias()
	if err != nil {
		t.Error(err.Error())
	}
	assertFloat64(t, bias.GetValue(), 35.994, "bias.GetValue")

	rxPower, err := laser.GetRxPower()
	if err != nil {
		t.Error(err.Error())
	}
	assertFloat64(t, rxPower.GetValue(), 0.0, "rxPower.GetValue")

	thresholds, err := rxPower.GetAlarmThresholds()
	if err != nil {
		t.Error(err.Error())
	}

	assertFloat64(t, thresholds.GetHighWarning(), 1.0, "thresholds.GetHighWarning")
	assertFloat64(t, thresholds.GetLowAlarm(), 0.0158, "thresholds.GetLowAlarm")
	assertFloat64(t, thresholds.GetLowWarning(), 0.0316, "thresholds.GetLowWarning")
}

func TestParseEEPROMwithoutMonitoring(t *testing.T) {
	eeprom := getEEPROM2(t)

	assertFloat64(t, eeprom.GetWavelength(), 0, "eeprom.GetWavelength")
	lasers := eeprom.GetLasers()
	assertInt(t, len(lasers), 1, "len(lasers)")
	laser := lasers[0]
	_, err := laser.GetBias()
	assertString(t, err.Error(), "This module does not implement monitoring", "err.Error")
	_, err = laser.GetRxPower()
	assertString(t, err.Error(), "This module does not implement monitoring", "err.Error")
	_, err = laser.GetTxPower()
	assertString(t, err.Error(), "This module does not implement monitoring", "err.Error")
	_, err = eeprom.GetModuleTemperature()
	assertString(t, err.Error(), "Monitoring not implemented by module", "err.Error")
	_, err = eeprom.GetModuleVoltage()
	assertString(t, err.Error(), "Monitoring not implemented by module", "err.Error")
}

func TestParseEEPROMTooShort(t *testing.T) {
	eeprom, err := NewEEPROM(make([]byte, 42))
	assertString(t, err.Error(), "Required at least 256 bytes to comply with SFF8472", "err.Error()")
	if eeprom != nil {
		t.Error("Expected EEPROM to be nil")
	}
}

func TestCalibration(t *testing.T) {
	eeprom := getEEPROM(t)

	eeprom.Diagnostics.Temperature = 25.5
	eeprom.ExternalCalibrationConstants = &ExternalCalibrationConstants{
		TemperatureSlope:  2.0,
		TemperatureOffset: -25.5,
	}

	eeprom.calibrate()

	if eeprom.Diagnostics.Temperature != 25.5 {
		t.Errorf("Expected 25.5, got %f", eeprom.Diagnostics.Temperature)
	}
}

func TestParseEncoding(t *testing.T) {
	encoding := Encoding(0x04)

	if encoding.String() != "Manchester" {
		t.Errorf("Expected Manchester, got %v", encoding.String())
	}
}

func TestParseDiagnosticMonitoringType(t *testing.T) {
	diagnosticMonitoringType := NewDiagnosticMonitoringType(0b1011000)
	diagnosticMonitoringType2 := DiagnosticMonitoringType{
		DiagnosticMonitoringImplemented: true,
		ExternallyCalibrated:            true,
		ReceivedPowerMeasurementType:    ReceivedPowerMeasurementTypeAveragePower,
	}
	if *diagnosticMonitoringType != diagnosticMonitoringType2 {
		t.Errorf("Got %+v, expected %+v", diagnosticMonitoringType, diagnosticMonitoringType2)
	}
}

func TestParseEnhancedOptions(t *testing.T) {
	enhancedOptions := NewEnhancedOptions(0b00010000)
	enhancedOptions1 := EnhancedOptions{
		SoftRxLosImplemented: true,
	}
	if *enhancedOptions != enhancedOptions1 {
		t.Errorf("Got %+v, expected %+v", enhancedOptions, enhancedOptions1)
	}
}

func TestParseTemperature(t *testing.T) {
	temp := parseTemperature(0b01111111, 0b11111111)
	assertFloat64(t, temp, 127.0+(255.0/256.0), "parseTemperature(0b01111111, 0b11111111)")
	temp = parseTemperature(0b01111101, 0)
	assertFloat64(t, temp, 125, "parseTemperature(0b01111101, 0)")
	temp = parseTemperature(0, 0)
	assertFloat64(t, temp, 0, "parseTemperature(0, 0)")
	temp = parseTemperature(0b11111111, 0b11111111)
	assertFloat64(t, temp, -1.0/256.0, "parseTemperature(0b11111111, 0b11111111)")

}

func TestParseCompliance(t *testing.T) {
	compliance := Compliance(0x06)
	if compliance.String() != "11.3" {
		t.Errorf("Got %+v, expected '13.3'", compliance.String())
	}
}

func TestJSONserialize(t *testing.T) {
	eeprom := getEEPROM(t)

	_, err := json.Marshal(eeprom)
	if err != nil {
		t.Errorf("JSON serializing failed: %+v", err)
	}
}

func TestLaserAbstraction(t *testing.T) {
	lasers := getEEPROM(t).GetLasers()

	if len(lasers) != 1 {
		t.Errorf("Expected length 1, got %d", len(lasers))
	}

	laser := lasers[0]

	if !laser.SupportsMonitoring() {
		t.Errorf("Laser does support monitoring, got: Laser does not support monitoring")
	}
	// if _, err := laser.GetBias(); err != nil {
	//     t.Errorf("Cpi;")
	// }
}
