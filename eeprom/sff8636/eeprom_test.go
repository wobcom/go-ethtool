package sff8636

import (
	"encoding/hex"
	"encoding/json"
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
	hexRaw := "110702000000000000000000000000000000000000001c97000081b9000000000000436a31c82b822ed64340414047c045402ccc2d91302c2f9d0000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000001f00000000000000000000000000000000000000000011cc07800000000000000005ff0002000000004446532020202020202020202020202020000002c95153465032382d4952342d3130304720413165bf00ce00fb0307ffde4331383132313535343631202020202031393031313020200c1068bc0000000000000000000000000000000000000000000000000000000014320000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000f6004b00fb000000000000000000908871708c7075480000000000000000000000000000000000000000000000006e18016357730232927c138888b81d4c6e170584577306f20000000000000000000000000000000000000000000000000000000000000000000000000077111100000000000000000000000000000000"
	return getEEPROMfromHex(t, hexRaw)
}

func getEEPROM1(t *testing.T) *EEPROM {
	hexRaw := "0d0500000000000000000000000000000000000000001bd800007dcb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000020004000000000000000000000000cab1efed000d0023880000000000000000ff000000000001a0465320202020202020202020202020201f0002c95132382d504378782020202020202020413206080a1000a10b00000043313930373235333435312d322020203139303830352020000067f03132383835353232333058420000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000f6004b00fb000000000000000000908871708c7075480000000000000000000000000000000000000000000000004e0001f545760277157c03e8138805dc4df009cf45760c5a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
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

	assertString(t, eeprom.GetIdentifier().String(), "QSFP28", "eeprom.GetIdentifier().String")
	assertString(t, eeprom.GetConnectorType().String(), "LC", "eeprom.GetConnectorType().String")
	assertString(t, eeprom.GetEncoding(), "64B/66B", "eeprom.GetEncoding")
	assertFloat64(t, eeprom.GetPowerClass().GetMaxPower(), 3.5, "eeprom.GetPowerClass().GetMaxPower")
	assertFloat64(t, eeprom.GetSignalingRate(), 26000000000, "eeprom.GetSignalingRate")
	linkLengths := eeprom.GetSupportedLinkLengths()
	assertFloat64(t, linkLengths["SMF"], 2000, "linkLengths[\"SMF(km)\"]")
	assertString(t, eeprom.GetVendorName(), "FS", "eeprom.GetVendorName")
	assertString(t, eeprom.GetVendorPN(), "QSFP28-IR4-100G", "eeprom.GetVendorPN")
	assertString(t, eeprom.GetVendorRev(), "A1", "eeprom.GetVendorRev")
	assertString(t, eeprom.GetVendorSN(), "C1812155461", "eeprom.GetVendorSN")
	assertString(t, eeprom.GetVendorOUI().String(), "00:02:C9", "eeprom.GetVendorOUI().String()")
	expectedDate, _ := time.Parse("060102", "190110")
	if dateCode := eeprom.GetDateCode(); dateCode != expectedDate {
		t.Errorf("eeprom.GetDateCode() returned %s, but expected %s", dateCode.Format("060102"), expectedDate.Format("060102"))
	}
	assertFloat64(t, eeprom.GetWavelength(), 1302.35, "eeprom.GetWavelength")
	lasers := eeprom.GetLasers()
	assertInt(t, len(lasers), 4, "len(lasers)")

	assertBool(t, eeprom.SupportsMonitoring(), true, "eeprom.SupportsMonitoring")

	laser0 := lasers[0]
	assertBool(t, laser0.SupportsMonitoring(), true, "laser0.SupportsMonitoring")

	bias, err := laser0.GetBias()
	if err != nil {
		t.Error(err.Error())
	}
	assertBool(t, bias.SupportsThresholds(), true, "bias.SupportsThresholds")
	assertFloat64(t, bias.GetValue(), 34.432, "bias.GetValue")
	thresholds, err := bias.GetAlarmThresholds()
	if err != nil {
		t.Error(err.Error())
	}
	assertFloat64(t, thresholds.GetHighAlarm(), 75, "thresholds.GetHighAlarm")
	assertFloat64(t, thresholds.GetLowAlarm(), 10, "thresholds.GetLowAlarm")
	assertFloat64(t, thresholds.GetHighWarning(), 70, "thresholds.GetHighWarning")
	assertFloat64(t, thresholds.GetLowWarning(), 15, "thresholds.GetLowWarning")

	assertString(t, eeprom.DeviceTechnology.TransmitterTechnology.String(), "1310 nm DFB", "eeprom.DeviceTechnology.String")
}

func TestParseEEPROM1(t *testing.T) {
	eeprom := getEEPROM1(t)

	assertBool(t, eeprom.SpecificationCompliance.IsNonOpticalImplementation(), true, "eeprom.SpecificationCompliance.IsNonOpticalImplementation")
	assertFloat64(t, eeprom.Wavelength, 0, "eeprom.Wavelength")
	assertInt(t, int(eeprom.CopperAttenuation2_5GHz), 6, "int(eeprom.CopperAttenuation2_5GHz)")
	assertInt(t, int(eeprom.CopperAttenuation5GHz), 8, "int(eeprom.CopperAttenuation5GHz)")
	assertInt(t, int(eeprom.CopperAttenuation7GHz), 10, "int(eeprom.CopperAttenuation7GHz)")
	assertInt(t, int(eeprom.CopperAttenuation12_9GHz), 16, "int(eeprom.CopperAttenuation12_9GHz)")
}

func TestJsonMarshal(t *testing.T) {
	eeprom := getEEPROM(t)

	_, err := json.Marshal(eeprom)
	if err != nil {
		t.Error(err.Error())
	}
}
