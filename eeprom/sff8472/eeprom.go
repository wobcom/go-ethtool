package sff8472

import (
	"errors"
	"fmt"
	"github.com/wobcom/go-ethtool/eeprom"
	"github.com/wobcom/go-ethtool/eeprom/sff8024"
	"github.com/wobcom/go-ethtool/eeprom/sff8079"
	"strings"
)

// ChecksumError occurs when the data read from the EEPROM does not a have a valid checksum as defined in SFF8472
type ChecksumError struct{}

func (c *ChecksumError) Error() string { return "Invalid checksum" }

// NewChecksumError returns a new ChecksumError instance
func NewChecksumError() *ChecksumError { return &ChecksumError{} }

/* Memory offsets */
const (
	/* Page A0h */
	/* Base ID Fields */
	identifierOffset             = 0x00 /* Type of transceiver */
	extendedIdentifierOffset     = 0x01 /* Extended identifier of type of transceiver */
	connectorOffset              = 0x02 /* Code for connector type */
	transceiverComplianceOffset  = 0x03 /* Code for electronic or optical compatibility */
	encodingOffset               = 0x0B /* Code for high speed serial encoding algorithm */
	baurateNominalOffset         = 0x0C /* Nominal signalling rate, units of 100 MBd */
	rateIdentifierOffset         = 0x0D /* Type of rate select functionality */
	lengthSMFkmOffset            = 0x0E /* Link length supported for single mode fiber, units of km */
	lengthSMF100mOffset          = 0x0F /*  Link length supported for single mode fiber, units of 100 m */
	lengthOM2Offset              = 0x10 /*  Link length supported for 50 um OM2 fiber, units of 10 m */
	lengthOM1Offset              = 0x11 /* Link length supported for 62.5 um OM1 fiber, units of 10 m */
	lengthOM4orCopperOffset      = 0x12 /* Link length supported for 50um OM4 fiber, units of 10m. Alternatively copper or direct attach cable, units of m */
	lengthOM3Offset              = 0x13 /* Link length supported for 50 um OM3 fiber, units of 10 m */
	vendorStartOffset            = 0x14 /* SFP vendor name (ASCII) */
	vendorEndOffset              = 0x23
	transceiverComplianceOffset1 = 0x24 /* Code for electronic or optical compatibility */
	vendorOuiOffset              = 0x23
	vendorPnStartOffset          = 0x28 /* Part number provided by SFP vendor (ASCII) */
	vendorPnEndOfffset           = 0x37
	vendorRevStartOffset         = 0x38 /* Revision level for part number provided by vendor (ASCII) */
	vendorRevEndOffset           = 0x3B
	wavelengthOffset             = 0x3C /* Laser wavelength (Passive/Active Cable Specification Compliance) */
	/* Extended ID Fields */
	optionsOffset                  = 0x40 /* Indicates which optional transceiver signals are implemented */
	uppertBitrateMarginOffset      = 0x42 /* Upper bit rate margin, units of % */
	lowerBitrateMarginOffset       = 0x43 /* Lower bit rate margin, units of % */
	vendorSnStartOffset            = 0x44 /* Serial number provided by vendor (ASCII) */
	vendorSnEndOffset              = 0x53
	dateCodeStartOffset            = 0x54 /* Vendor's manufacturing date code */
	dateCodeEndOffset              = 0x5B
	diagnosticMonitoringTypeOffset = 0x5C /* Indicates which type of diagnostic monitoring is implemented (if any) in the transceiver (see Table 8-5) */
	enhancedOptionsOffset          = 0x5D /*  Indicates which optional enhanced features are implemented (if any) in the transceiver (see Table 8-6) */
	complianceOffset               = 0x5E /* Indicates which revision of SFF-8472 the transceiver complies with. (see Table 8-8). */
	checksumOffset                 = 0x5F /* Byte 95 contains the low order 8 bits of the sum of bytes 0-94 */

	/* Page A2h */
	/* Diagnostic and control/status fields */
	thresholdsOffset   = 0x100 /* Diagnostic Flag Alarm and Warning Thresholds (see Table 9-5) */
	optionalThresholds = 0x128 /* Thresholds for optional Laser Temperature and TEC Current alarms
	and warnings (see Table 9-5) */
	externalCalibrationConstantsOffset = 0x138 /* Diagnostic calibration constants for optional External Calibration (see Table 9-6) */
	diagnosticsOffset                  = 0x160 /* Diagnostic Monitor Data (internally or externally calibrated) (see Table 9-11) */
	optionalDiagnosticsOffset          = 0x16A /*  Monitor Data for Optional Laser temperature and TEC current (see Table 9-11) */
	statusControlOffset                = 0x16E /* Optional Status and Control Bits (see Table 9-11) */
	alarmFlagsOffset                   = 0x170 /* Diagnostic Alarm Flag Status Bits (see Table 9-12) */
	unallocatedOffset                  = 0x172 /* SIC! Standard contradicts itself, see Table 9-12 this byte is in use! */
	cdrUnlockedOffset                  = 0x173 /* SIC! Table 9-12 and Table 4-2 contradict, going for specification in 9-12 because this is precise enough to write an implementation */
	warningFlagsOffset                 = 0x174 /* Diagnostic Warning Flag Status Bits (see Table 9-12) */
	extendedStatusControlOffset        = 0x176 /* Extended module control and status bytes (see Table 10-1) */
	userEepromStartOffset              = 0x180 /* User writeable EEPROM */
	userEepromEndOffset                = 0x1F7
)

// EEPROM implementation is based on SFF-8462 Rev. 12.3
type EEPROM struct {
	Raw []byte
	/* Page A0h */
	/* Base ID Fields */
	Identifier                sff8024.Identifier
	ExtendedIdentifier        sff8079.ExtendedIdentifier
	ConnectorType             sff8024.ConnectorType
	Encoding                  Encoding
	SignalingRate             float64
	RateIdentifier            sff8079.RateIdentifier
	LengthSMFKm               float64
	LengthSMF                 float64
	LengthOM2                 float64
	LengthOM1                 float64
	LengthOM4OrDAC            float64
	LengthOM3                 float64
	VendorName                string
	TransceiverCompliance     sff8079.Compliance
	VendorOUI                 eeprom.OUI
	VendorPN                  string
	VendorRev                 string
	Wavelength                float64
	PassiveCableSpecification sff8079.PassiveCableSpecifications /* used if 0x08 Bit 2 set */
	ActiveCableSpecification  sff8079.ActiveCableSpecifications  /* used if 0x08 Bit 3 set */
	/* Extended ID Fields */
	Options                  *sff8079.Options
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

// NewEEPROM parses a byte slice of at least 256 size into a new sff8472.EERPOM instance
func NewEEPROM(raw []byte) (*EEPROM, error) {
	if len(raw) < 256 {
		return nil, errors.New("Required at least 256 bytes to comply with SFF8472")
	}

	e := &EEPROM{
		Raw: raw,
		/* Page A0h */
		/* Base ID Fields */
		Identifier:         sff8024.Identifier(raw[identifierOffset]),
		ExtendedIdentifier: sff8079.ExtendedIdentifier(raw[extendedIdentifierOffset]),
		ConnectorType:      sff8024.ConnectorType(raw[connectorOffset]),
		Encoding:           Encoding(raw[encodingOffset]),
		SignalingRate: func() float64 {
			if raw[baurateNominalOffset] == 0xFF {
				return float64(uint16(raw[uppertBitrateMarginOffset])<<8|uint16(raw[lowerBitrateMarginOffset])) * 100 * 1000000
			}
			return float64(raw[baurateNominalOffset]) * 100 * 1000000
		}(),
		RateIdentifier: sff8079.RateIdentifier(raw[rateIdentifierOffset]),
		LengthSMFKm:    float64(raw[lengthSMFkmOffset]),
		LengthSMF:      float64(raw[lengthSMF100mOffset]) * 100,
		LengthOM2:      float64(raw[lengthOM2Offset]) * 10,
		LengthOM1:      float64(raw[lengthOM1Offset]) * 10,
		LengthOM4OrDAC: float64(raw[lengthOM4orCopperOffset]),
		LengthOM3:      float64(raw[lengthOM3Offset]) * 10,
		VendorName:     strings.Trim(parseString(raw[vendorStartOffset:vendorEndOffset+1]), " "),
		TransceiverCompliance: sff8079.NewCompliance([8]byte{
			raw[transceiverComplianceOffset+0],
			raw[transceiverComplianceOffset+1],
			raw[transceiverComplianceOffset+2],
			raw[transceiverComplianceOffset+3],
			raw[transceiverComplianceOffset+4],
			raw[transceiverComplianceOffset+5],
			raw[transceiverComplianceOffset+6],
			raw[transceiverComplianceOffset+7],
		}),
		VendorOUI: eeprom.NewOUI([3]byte{
			raw[vendorOuiOffset+0],
			raw[vendorOuiOffset+1],
			raw[vendorOuiOffset+2],
		}),
		VendorPN:  strings.Trim(parseString(raw[vendorPnStartOffset:vendorPnEndOfffset+1]), " "),
		VendorRev: strings.Trim(parseString(raw[vendorRevStartOffset:vendorRevEndOffset+1]), " "),
		Wavelength: parseWavelength(
			raw[wavelengthOffset+0],
			raw[wavelengthOffset+1],
		),
		PassiveCableSpecification: sff8079.NewPassiveCableSpecifications([2]byte{
			raw[wavelengthOffset+0],
			raw[wavelengthOffset+1],
		}),
		ActiveCableSpecification: sff8079.NewActiveCableSpecifications([2]byte{
			raw[wavelengthOffset+0],
			raw[wavelengthOffset+1],
		}),
		Options: sff8079.NewOptions([2]byte{
			raw[optionsOffset+0],
			raw[optionsOffset+1],
		}),
		UpperBitrateMargin:       uint8(raw[uppertBitrateMarginOffset]),
		LowerBitrateMargin:       uint8(raw[lowerBitrateMarginOffset]),
		VendorSN:                 strings.Trim(parseString(raw[vendorSnStartOffset:vendorSnEndOffset+1]), " "),
		DateCode:                 strings.Trim(parseString(raw[dateCodeStartOffset:dateCodeEndOffset+1]), " "),
		DiagnosticMonitoringType: NewDiagnosticMonitoringType(raw[diagnosticMonitoringTypeOffset]),
		EnhancedOptions:          NewEnhancedOptions(raw[enhancedOptionsOffset]),
		Compliance:               Compliance(raw[complianceOffset]),
	}

	if len(raw) >= 512 {
		sum := uint(0)
		for i := 0x100; i < 0x15F; i++ {
			sum += uint(raw[i])
		}
		if byte(sum&0xFF) != raw[0x15F] {
			fmt.Printf("Expected checksum %02X, got %X\n", raw[checksumOffset], sum)
			return nil, NewChecksumError()
		}
		e.Thresholds = NewThresholds([40]byte{
			raw[thresholdsOffset+0],
			raw[thresholdsOffset+1],
			raw[thresholdsOffset+2],
			raw[thresholdsOffset+3],
			raw[thresholdsOffset+4],
			raw[thresholdsOffset+5],
			raw[thresholdsOffset+6],
			raw[thresholdsOffset+7],
			raw[thresholdsOffset+8],
			raw[thresholdsOffset+9],
			raw[thresholdsOffset+10],
			raw[thresholdsOffset+11],
			raw[thresholdsOffset+12],
			raw[thresholdsOffset+13],
			raw[thresholdsOffset+14],
			raw[thresholdsOffset+15],
			raw[thresholdsOffset+16],
			raw[thresholdsOffset+17],
			raw[thresholdsOffset+18],
			raw[thresholdsOffset+19],
			raw[thresholdsOffset+20],
			raw[thresholdsOffset+21],
			raw[thresholdsOffset+22],
			raw[thresholdsOffset+23],
			raw[thresholdsOffset+24],
			raw[thresholdsOffset+25],
			raw[thresholdsOffset+26],
			raw[thresholdsOffset+27],
			raw[thresholdsOffset+28],
			raw[thresholdsOffset+29],
			raw[thresholdsOffset+30],
			raw[thresholdsOffset+31],
			raw[thresholdsOffset+32],
			raw[thresholdsOffset+33],
			raw[thresholdsOffset+34],
			raw[thresholdsOffset+35],
			raw[thresholdsOffset+36],
			raw[thresholdsOffset+37],
			raw[thresholdsOffset+38],
			raw[thresholdsOffset+39],
		})
		e.OptionalThresholds = NewOptionalThresholds([16]byte{
			raw[optionalThresholds+0],
			raw[optionalThresholds+1],
			raw[optionalThresholds+2],
			raw[optionalThresholds+3],
			raw[optionalThresholds+4],
			raw[optionalThresholds+5],
			raw[optionalThresholds+6],
			raw[optionalThresholds+7],
			raw[optionalThresholds+8],
			raw[optionalThresholds+9],
			raw[optionalThresholds+10],
			raw[optionalThresholds+11],
			raw[optionalThresholds+12],
			raw[optionalThresholds+13],
			raw[optionalThresholds+14],
			raw[optionalThresholds+15],
		})
		e.ExternalCalibrationConstants = NewExternalCalibrationConstants([36]byte{
			raw[externalCalibrationConstantsOffset+0],
			raw[externalCalibrationConstantsOffset+1],
			raw[externalCalibrationConstantsOffset+2],
			raw[externalCalibrationConstantsOffset+3],
			raw[externalCalibrationConstantsOffset+4],
			raw[externalCalibrationConstantsOffset+5],
			raw[externalCalibrationConstantsOffset+6],
			raw[externalCalibrationConstantsOffset+7],
			raw[externalCalibrationConstantsOffset+8],
			raw[externalCalibrationConstantsOffset+9],
			raw[externalCalibrationConstantsOffset+10],
			raw[externalCalibrationConstantsOffset+11],
			raw[externalCalibrationConstantsOffset+12],
			raw[externalCalibrationConstantsOffset+13],
			raw[externalCalibrationConstantsOffset+14],
			raw[externalCalibrationConstantsOffset+15],
			raw[externalCalibrationConstantsOffset+16],
			raw[externalCalibrationConstantsOffset+17],
			raw[externalCalibrationConstantsOffset+18],
			raw[externalCalibrationConstantsOffset+19],
			raw[externalCalibrationConstantsOffset+20],
			raw[externalCalibrationConstantsOffset+21],
			raw[externalCalibrationConstantsOffset+22],
			raw[externalCalibrationConstantsOffset+23],
			raw[externalCalibrationConstantsOffset+24],
			raw[externalCalibrationConstantsOffset+25],
			raw[externalCalibrationConstantsOffset+26],
			raw[externalCalibrationConstantsOffset+27],
			raw[externalCalibrationConstantsOffset+28],
			raw[externalCalibrationConstantsOffset+29],
			raw[externalCalibrationConstantsOffset+30],
			raw[externalCalibrationConstantsOffset+31],
			raw[externalCalibrationConstantsOffset+32],
			raw[externalCalibrationConstantsOffset+33],
			raw[externalCalibrationConstantsOffset+34],
			raw[externalCalibrationConstantsOffset+35],
		})
		e.Diagnostics = NewDiagnostics([10]byte{
			raw[diagnosticsOffset+0],
			raw[diagnosticsOffset+1],
			raw[diagnosticsOffset+2],
			raw[diagnosticsOffset+3],
			raw[diagnosticsOffset+4],
			raw[diagnosticsOffset+5],
			raw[diagnosticsOffset+6],
			raw[diagnosticsOffset+7],
			raw[diagnosticsOffset+8],
			raw[diagnosticsOffset+9],
		})
		e.OptionalDiagnostics = NewOptionalDiagnostics([4]byte{
			raw[optionalDiagnosticsOffset+0],
			raw[optionalDiagnosticsOffset+1],
			raw[optionalDiagnosticsOffset+2],
			raw[optionalDiagnosticsOffset+3],
		})
		e.StatusControl = NewStatusControl(raw[statusControlOffset])
		e.AlarmFlags = NewAlarmFlags([2]byte{
			raw[alarmFlagsOffset+0],
			raw[alarmFlagsOffset+1],
		})
		e.InputEqualizationControl = NewInputEqualizationControl(raw[unallocatedOffset]) // Table 4-2 and Table 9-12 contradict, going for Table 9-12 which allows for an implementation
		e.OutputEmphasisControl = NewOutputEmphasisControl(raw[cdrUnlockedOffset])
		e.WarningFlags = NewWarningFlags([2]byte{
			raw[warningFlagsOffset+0],
			raw[warningFlagsOffset+1],
		})
		e.ExtendedStatusControl = NewExtendedStatusControl([2]byte{
			raw[extendedStatusControlOffset+0],
			raw[extendedStatusControlOffset+1],
		})
		e.UserData = raw[userEepromStartOffset:userEepromEndOffset]
	}

	if e.TransceiverCompliance.IsSFPCableImplementation() {
		e.Wavelength = 0
		e.LengthOM4OrDAC /= 10
		if e.TransceiverCompliance[sff8079.ComplianceFlagActiveCable] {
			e.PassiveCableSpecification = sff8079.PassiveCableSpecifications{}
		} else {
			e.ActiveCableSpecification = sff8079.ActiveCableSpecifications{}
		}
	} else {
		e.PassiveCableSpecification = sff8079.PassiveCableSpecifications{}
		e.ActiveCableSpecification = sff8079.ActiveCableSpecifications{}
	}

	// apply calibration data if necessary
	if e.DiagnosticMonitoringType.ExternallyCalibrated && len(raw) >= 512 {
		e.calibrate()
	}

	return e, nil
}
