package SFF8472

import (
	"errors"
	"gitlab.com/wobcom/ethtool/eeprom"
	"gitlab.com/wobcom/ethtool/eeprom/sff8024"
	"gitlab.com/wobcom/ethtool/eeprom/sff8079"
	"strings"
)

/* Memory offsets */
const (
	/* Page A0h */
	/* Base ID Fields */
	SFF8472_IDENTIFIER                  = 0x00 /* Type of transceiver */
	SFF8472_EXTENDED_IDENTIFIER         = 0x01 /* Extended identifier of type of transceiver */
	SFF8472_CONNECTOR                   = 0x02 /* Code for connector type */
	SFF8472_TRANSCEIVER_COMPLIANCE      = 0x03 /* Code for electronic or optical compatibility */
	SFF8472_ENCODING                    = 0x0B /* Code for high speed serial encoding algorithm */
	SFF8472_BAUDRATE_NOMINAL            = 0x0C /* Nominal signalling rate, units of 100 MBd */
	SFF8472_RATE_IDENTIFIER             = 0x0D /* Type of rate select functionality */
	SFF8472_LENGTH_SMF_KM               = 0x0E /* Link length supported for single mode fiber, units of km */
	SFF8472_LENGTH_SMF_100M             = 0x0F /*  Link length supported for single mode fiber, units of 100 m */
	SFF8472_LENGTH_OM2                  = 0x10 /*  Link length supported for 50 um OM2 fiber, units of 10 m */
	SFF8472_LENGTH_OM1                  = 0x11 /* Link length supported for 62.5 um OM1 fiber, units of 10 m */
	SFF8472_LENGTH_OM4_COPEPR_DAC_CABLE = 0x12 /* Link length supported for 50um OM4 fiber, units of 10m. Alternatively copper or direct attach cable, units of m */
	SFF8472_LENGTH_OM3                  = 0x13 /* Link length supported for 50 um OM3 fiber, units of 10 m */
	SFF8472_VENDOR_START                = 0x14 /* SFP vendor name (ASCII) */
	SFF8472_VENDOR_END                  = 0x23
	SFF8472_TRANSCEIVER_COMPLIANCE1     = 0x24 /* Code for electronic or optical compatibility */
	SFF8472_VENDOR_OUI                  = 0x23
	SFF8472_VENDOR_PN_START             = 0x28 /* Part number provided by SFP vendor (ASCII) */
	SFF8472_VENDOR_PN_END               = 0x37
	SFF8472_VENDOR_REV_START            = 0x38 /* Revision level for part number provided by vendor (ASCII) */
	SFF8472_VENDOR_REV_END              = 0x3B
	SFF8472_WAVELENGTH                  = 0x3C /* Laser wavelength (Passive/Active Cable Specification Compliance) */
	/* Extended ID Fields */
	SFF8472_OPTIONS                    = 0x40 /* Indicates which optional transceiver signals are implemented */
	SFF8472_UPPER_BITRATE_MARGIN       = 0x42 /* Upper bit rate margin, units of % */
	SFF8472_LOWER_BITRATE_MARGIN       = 0x43 /* Lower bit rate margin, units of % */
	SFF8472_VENDOR_SN_START            = 0x44 /* Serial number provided by vendor (ASCII) */
	SFF8472_VENDOR_SN_END              = 0x53
	SFF8472_DATE_CODE_START            = 0x54 /* Vendor's manufacturing date code */
	SFF8472_DATE_CODE_END              = 0x5B
	SFF8472_DIAGNOSITC_MONITORING_TYPE = 0x5C /* Indicates which type of diagnostic monitoring is implemented (if any) in the transceiver (see Table 8-5) */
	SFF8472_ENHANCED_OPTIONS           = 0x5D /*  Indicates which optional enhanced features are implemented (if any) in the transceiver (see Table 8-6) */
	SFF8472_COMPLIANCE                 = 0x5E /* Indicates which revision of SFF-8472 the transceiver complies with. (see Table 8-8). */

	/* Page A2h */
	/* Diagnostic and control/status fields */
	SFF8472_THRESHOLDS          = 0x100 /* Diagnostic Flag Alarm and Warning Thresholds (see Table 9-5) */
	SFF8472_OPTIONAL_THRESHOLDS = 0x128 /* Thresholds for optional Laser Temperature and TEC Current alarms
	and warnings (see Table 9-5) */
	SFF8472_EXTERNAL_CALIBRATION_CONSTANTS = 0x138 /* Diagnostic calibration constants for optional External Calibration (see Table 9-6) */
	SFF8472_DIAGNOSTICS                    = 0x160 /* Diagnostic Monitor Data (internally or externally calibrated) (see Table 9-11) */
	SFF8472_OPTIONAL_DIAGNOSTICS           = 0x16A /*  Monitor Data for Optional Laser temperature and TEC current (see Table 9-11) */
	SFF8472_STATUS_CONTROL                 = 0x16E /* Optional Status and Control Bits (see Table 9-11) */
	SFF8472_ALARM_FLAGS                    = 0x170 /* Diagnostic Alarm Flag Status Bits (see Table 9-12) */
	SFF8472_UNALLOCATED                    = 0x172 /* SIC! Standard contradicts itself, see Table 9-12 this byte is in use! */
	SFF8472_CDR_UNLOCKED                   = 0x173 /* SIC! Table 9-12 and Table 4-2 contradict, going for specification in 9-12 because this is precise enough to write an implementation */
	SFF8472_WARNING_FLAGS                  = 0x174 /* Diagnostic Warning Flag Status Bits (see Table 9-12) */
	SFF8472_EXTENDED_STATUS_CONTROL        = 0x176 /* Extended module control and status bytes (see Table 10-1) */
	SFF8472_USER_EEPROM_START              = 0x180 /* User writeable EEPROM */
	SFF8472_USER_EEPROM_END                = 0x1F7
)

// EEPROM implementation is based on SFF-8462 Rev. 12.3
type EEPROM struct {
	/* Page A0h */
	/* Base ID Fields */
	Identifier                SFF8024.Identifier
	ExtendedIdentifier        SFF8079.ExtendedIdentifier
	ConnectorType             SFF8024.ConnectorType
	Encoding                  Encoding
	SignalingRate             float64
	RateIdentifier            SFF8079.RateIdentifier
	LengthSMFKm               float64
	LengthSMF                 float64
	LengthOM2                 float64
	LengthOM1                 float64
	LengthOM4OrDAC            float64
	LengthOM3                 float64
	VendorName                string
	TransceiverCompliance     SFF8079.Compliance
	VendorOUI                 eeprom.OUI
	VendorPN                  string
	VendorRev                 string
	Wavelength                float64
	PassiveCableSpecification SFF8079.PassiveCableSpecifications /* used if 0x08 Bit 2 set */
	ActiveCableSpecification  SFF8079.ActiveCableSpecifications  /* used if 0x08 Bit 3 set */
	/* Extended ID Fields */
	Options                  *SFF8079.Options
	UpperBitrateMargin       uint8
	LowerBitrateMargin       uint8
	VendorSN                 string
	DateCode                 string
	DiagnosticMonitoringType *DiagnosticMonitoringType
	EnhancedOptions          *EnhancedOptions
	Compliance               Compliance
	/* Page A2h */
	/* Diagnostic and control/status fields */
	Thresholds                   *Thresholds
	OptionalThresholds           *OptionalThresholds
	ExternalCalibrationConstants *ExternalCalibrationConstants
	Diagnostics                  *Diagnostics
	OptionalDiagnostics          *OptionalDiagnostics
	StatusControl                *StatusControl
	AlarmFlags                   *AlarmFlags
	InputEqualizationControl     *InputEqualizationControl
	OutputEmphasisControl        *OutputEmphasisControl
	WarningFlags                 *WarningFlags
	ExtendedStatusControl        *ExtendedStatusControl
	UserData                     []byte
}

func NewEEPROM(raw []byte) (*EEPROM, error) {
	if len(raw) < 256 {
		return nil, errors.New("Required at least 256 bytes to comply with SFF8472")
	}

	e := &EEPROM{
		/* Page A0h */
		/* Base ID Fields */
		Identifier:         SFF8024.Identifier(raw[SFF8472_IDENTIFIER]),
		ExtendedIdentifier: SFF8079.ExtendedIdentifier(raw[SFF8472_EXTENDED_IDENTIFIER]),
		ConnectorType:      SFF8024.ConnectorType(raw[SFF8472_CONNECTOR]),
		Encoding:           Encoding(raw[SFF8472_ENCODING]),
		SignalingRate: func() float64 {
			if raw[SFF8472_BAUDRATE_NOMINAL] == 0xFF {
				return float64(uint16(raw[SFF8472_UPPER_BITRATE_MARGIN])<<8|uint16(raw[SFF8472_LOWER_BITRATE_MARGIN])) * 100 * 1000000
			}
			return float64(raw[SFF8472_BAUDRATE_NOMINAL]) * 100 * 1000000
		}(),
		RateIdentifier: SFF8079.RateIdentifier(raw[SFF8472_RATE_IDENTIFIER]),
		LengthSMFKm:    float64(raw[SFF8472_LENGTH_SMF_KM]),
		LengthSMF:      float64(raw[SFF8472_LENGTH_SMF_100M]) * 100,
		LengthOM2:      float64(raw[SFF8472_LENGTH_OM2]) * 10,
		LengthOM1:      float64(raw[SFF8472_LENGTH_OM1]) * 10,
		LengthOM4OrDAC: float64(raw[SFF8472_LENGTH_OM4_COPEPR_DAC_CABLE]),
		LengthOM3:      float64(raw[SFF8472_LENGTH_OM3]) * 10,
		VendorName:     strings.Trim(parseString(raw[SFF8472_VENDOR_START:SFF8472_VENDOR_END+1]), " "),
		TransceiverCompliance: SFF8079.NewCompliance([8]byte{
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+0],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+1],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+2],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+3],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+4],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+5],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+6],
			raw[SFF8472_TRANSCEIVER_COMPLIANCE+7],
		}),
		VendorOUI: eeprom.NewOUI([3]byte{
			raw[SFF8472_VENDOR_OUI+0],
			raw[SFF8472_VENDOR_OUI+1],
			raw[SFF8472_VENDOR_OUI+2],
		}),
		VendorPN:  strings.Trim(parseString(raw[SFF8472_VENDOR_PN_START:SFF8472_VENDOR_PN_END+1]), " "),
		VendorRev: strings.Trim(parseString(raw[SFF8472_VENDOR_REV_START:SFF8472_VENDOR_REV_END+1]), " "),
		Wavelength: parseWavelength(
			raw[SFF8472_WAVELENGTH+0],
			raw[SFF8472_WAVELENGTH+1],
		),
		PassiveCableSpecification: SFF8079.NewPassiveCableSpecifications([2]byte{
			raw[SFF8472_WAVELENGTH+0],
			raw[SFF8472_WAVELENGTH+1],
		}),
		ActiveCableSpecification: SFF8079.NewActiveCableSpecifications([2]byte{
			raw[SFF8472_WAVELENGTH+0],
			raw[SFF8472_WAVELENGTH+1],
		}),
		Options: SFF8079.NewOptions([2]byte{
			raw[SFF8472_OPTIONS+0],
			raw[SFF8472_OPTIONS+1],
		}),
		UpperBitrateMargin:       uint8(raw[SFF8472_UPPER_BITRATE_MARGIN]),
		LowerBitrateMargin:       uint8(raw[SFF8472_LOWER_BITRATE_MARGIN]),
		VendorSN:                 strings.Trim(parseString(raw[SFF8472_VENDOR_SN_START:SFF8472_VENDOR_SN_END+1]), " "),
		DateCode:                 strings.Trim(parseString(raw[SFF8472_DATE_CODE_START:SFF8472_DATE_CODE_END+1]), " "),
		DiagnosticMonitoringType: NewDiagnosticMonitoringType(raw[SFF8472_DIAGNOSITC_MONITORING_TYPE]),
		EnhancedOptions:          NewEnhancedOptions(raw[SFF8472_ENHANCED_OPTIONS]),
		Compliance:               Compliance(raw[SFF8472_COMPLIANCE]),
	}

	if len(raw) >= 512 {
		e.Thresholds = NewThresholds([40]byte{
			raw[SFF8472_THRESHOLDS+0],
			raw[SFF8472_THRESHOLDS+1],
			raw[SFF8472_THRESHOLDS+2],
			raw[SFF8472_THRESHOLDS+3],
			raw[SFF8472_THRESHOLDS+4],
			raw[SFF8472_THRESHOLDS+5],
			raw[SFF8472_THRESHOLDS+6],
			raw[SFF8472_THRESHOLDS+7],
			raw[SFF8472_THRESHOLDS+8],
			raw[SFF8472_THRESHOLDS+9],
			raw[SFF8472_THRESHOLDS+10],
			raw[SFF8472_THRESHOLDS+11],
			raw[SFF8472_THRESHOLDS+12],
			raw[SFF8472_THRESHOLDS+13],
			raw[SFF8472_THRESHOLDS+14],
			raw[SFF8472_THRESHOLDS+15],
			raw[SFF8472_THRESHOLDS+16],
			raw[SFF8472_THRESHOLDS+17],
			raw[SFF8472_THRESHOLDS+18],
			raw[SFF8472_THRESHOLDS+19],
			raw[SFF8472_THRESHOLDS+20],
			raw[SFF8472_THRESHOLDS+21],
			raw[SFF8472_THRESHOLDS+22],
			raw[SFF8472_THRESHOLDS+23],
			raw[SFF8472_THRESHOLDS+24],
			raw[SFF8472_THRESHOLDS+25],
			raw[SFF8472_THRESHOLDS+26],
			raw[SFF8472_THRESHOLDS+27],
			raw[SFF8472_THRESHOLDS+28],
			raw[SFF8472_THRESHOLDS+29],
			raw[SFF8472_THRESHOLDS+30],
			raw[SFF8472_THRESHOLDS+31],
			raw[SFF8472_THRESHOLDS+32],
			raw[SFF8472_THRESHOLDS+33],
			raw[SFF8472_THRESHOLDS+34],
			raw[SFF8472_THRESHOLDS+35],
			raw[SFF8472_THRESHOLDS+36],
			raw[SFF8472_THRESHOLDS+37],
			raw[SFF8472_THRESHOLDS+38],
			raw[SFF8472_THRESHOLDS+39],
		})
		e.OptionalThresholds = NewOptionalThresholds([16]byte{
			raw[SFF8472_OPTIONAL_THRESHOLDS+0],
			raw[SFF8472_OPTIONAL_THRESHOLDS+1],
			raw[SFF8472_OPTIONAL_THRESHOLDS+2],
			raw[SFF8472_OPTIONAL_THRESHOLDS+3],
			raw[SFF8472_OPTIONAL_THRESHOLDS+4],
			raw[SFF8472_OPTIONAL_THRESHOLDS+5],
			raw[SFF8472_OPTIONAL_THRESHOLDS+6],
			raw[SFF8472_OPTIONAL_THRESHOLDS+7],
			raw[SFF8472_OPTIONAL_THRESHOLDS+8],
			raw[SFF8472_OPTIONAL_THRESHOLDS+9],
			raw[SFF8472_OPTIONAL_THRESHOLDS+10],
			raw[SFF8472_OPTIONAL_THRESHOLDS+11],
			raw[SFF8472_OPTIONAL_THRESHOLDS+12],
			raw[SFF8472_OPTIONAL_THRESHOLDS+13],
			raw[SFF8472_OPTIONAL_THRESHOLDS+14],
			raw[SFF8472_OPTIONAL_THRESHOLDS+15],
		})
		e.ExternalCalibrationConstants = NewExternalCalibrationConstants([36]byte{
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+0],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+1],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+2],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+3],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+4],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+5],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+6],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+7],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+8],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+9],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+10],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+11],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+12],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+13],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+14],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+15],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+16],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+17],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+18],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+19],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+20],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+21],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+22],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+23],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+24],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+25],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+26],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+27],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+28],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+29],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+30],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+31],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+32],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+33],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+34],
			raw[SFF8472_EXTERNAL_CALIBRATION_CONSTANTS+35],
		})
		e.Diagnostics = NewDiagnostics([10]byte{
			raw[SFF8472_DIAGNOSTICS+0],
			raw[SFF8472_DIAGNOSTICS+1],
			raw[SFF8472_DIAGNOSTICS+2],
			raw[SFF8472_DIAGNOSTICS+3],
			raw[SFF8472_DIAGNOSTICS+4],
			raw[SFF8472_DIAGNOSTICS+5],
			raw[SFF8472_DIAGNOSTICS+6],
			raw[SFF8472_DIAGNOSTICS+7],
			raw[SFF8472_DIAGNOSTICS+8],
			raw[SFF8472_DIAGNOSTICS+9],
		})
		e.OptionalDiagnostics = NewOptionalDiagnostics([4]byte{
			raw[SFF8472_OPTIONAL_DIAGNOSTICS+0],
			raw[SFF8472_OPTIONAL_DIAGNOSTICS+1],
			raw[SFF8472_OPTIONAL_DIAGNOSTICS+2],
			raw[SFF8472_OPTIONAL_DIAGNOSTICS+3],
		})
		e.StatusControl = NewStatusControl(raw[SFF8472_STATUS_CONTROL])
		e.AlarmFlags = NewAlarmFlags([2]byte{
			raw[SFF8472_ALARM_FLAGS+0],
			raw[SFF8472_ALARM_FLAGS+1],
		})
		e.InputEqualizationControl = NewInputEqualizationControl(raw[SFF8472_UNALLOCATED]) // Table 4-2 and Table 9-12 contradict, going for Table 9-12 which allows for an implementation
		e.OutputEmphasisControl = NewOutputEmphasisControl(raw[SFF8472_CDR_UNLOCKED])
		e.WarningFlags = NewWarningFlags([2]byte{
			raw[SFF8472_WARNING_FLAGS+0],
			raw[SFF8472_WARNING_FLAGS+1],
		})
		e.ExtendedStatusControl = NewExtendedStatusControl([2]byte{
			raw[SFF8472_EXTENDED_STATUS_CONTROL+0],
			raw[SFF8472_EXTENDED_STATUS_CONTROL+1],
		})
		e.UserData = raw[SFF8472_USER_EEPROM_START:SFF8472_USER_EEPROM_END]
	}

	if e.TransceiverCompliance.IsSFPCableImplementation() {
		e.Wavelength = 0
		e.LengthOM4OrDAC /= 10
		if e.TransceiverCompliance[SFF8079.ComplianceFlagActiveCable] {
			e.PassiveCableSpecification = SFF8079.PassiveCableSpecifications{}
		} else {
			e.ActiveCableSpecification = SFF8079.ActiveCableSpecifications{}
		}
	} else {
		e.PassiveCableSpecification = SFF8079.PassiveCableSpecifications{}
		e.ActiveCableSpecification = SFF8079.ActiveCableSpecifications{}
	}

	// apply calibration data if necessary
	if e.DiagnosticMonitoringType.ExternallyCalibrated {
		e.calibrate()
	}

	return e, nil
}
