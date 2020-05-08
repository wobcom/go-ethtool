package sff8472

import (
    "testing"
    "encoding/hex"
    "encoding/json"
)

func getEEPROM(t *testing.T) *EEPROM {
    hexRaw := "03040720000000000000000667000a6400000000464c45584f5054495820202020202020000002c9502e42313639362e31302e44412020204120202005320087001a000046373942354b48202020202020202020313931323138202068b0038b0000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff46373942354b4820202024f3bb1f3ab6709be265b13a5c2089d85a00e7005500ec008ca0753088b87724afc803e89c4005dc457707cb372d09d0312d00fb2bd4013c000000000000000000000000000000000000000000000000000000003f800000000000000100000001000000010000000100000000000070198b7f21367018d9158e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
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

func TestParseEEPROM(t *testing.T) {
    eeprom := getEEPROM(t)

    if eeprom.VendorName != "FLEXOPTIX" {
        t.Errorf("Expected FLEXOPTIX, got %s", eeprom.VendorName)
    }
}

func TestCalibration(t *testing.T) {
    eeprom := getEEPROM(t)

    eeprom.Diagnostics.Temperature = 25.5
    eeprom.ExternalCalibrationConstants = &ExternalCalibrationConstants{
        TemperatureSlope: 2.0,
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
        ExternallyCalibrated: true,
        ReceivedPowerMeasurementType: ReceivedPowerMeasurementTypeAveragePower,
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
