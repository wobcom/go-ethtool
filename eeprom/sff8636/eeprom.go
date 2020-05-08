package sff8636

import (
	"errors"
	"gitlab.com/wobcom/ethtool/eeprom"
	"gitlab.com/wobcom/ethtool/eeprom/sff8024"
	"strings"
)

/* Memory offsets as defined in SFF-8636 Rev 2.10a (September 24, 2019) */
const (
	/* Lower Page 00h (Table 6-1) */

	// Identifier (See SFF-8024 Transceiver Management)
	identifierOffset = 0x00
	// Status
	statusIndicatorsOffset = 0x01
	// Interrupt Flags
	interruptFlagsOffset = 0x03
	// Free Side Device Monitors
	freeSideDeviceMonitorsOffset = 0x16
	// Channel Monitors
	channelMonitorsOffset = 0x22

	// 0x52 - 0x55 reserved

	// Control
	controlOffset = 0x56
	// Free Side Device and Channel Masks
	freeSideInterruptMasksOffset = 0x64
	// Free Side Device Properties
	freeSideDevicePropertiesOffset = 0x6B

	/* Upper Page 00h (Table 6-14) */
	// Identifier Type of free side device (See SFF-8024 Transceiver Management)
	// Note: Should read the same as identifierOffset
	identifierOffset1 = 0x80

	// Extended Identifier of free side device. Includes
	// power classes, CLEI codes, CDR capability (See
	// Table 6-15)
	extendedIdentifierOffset = 0x81

	// Code for media connector type (See SFF-8024
	// Transceiver Management)
	connectorTypeOffset = 0x82

	// Code for electronic or optical compatibility (See
	// Table 6-16)
	specificationCompliance = 0x83

	// Code for serial encoding algorithm. (See SFF-8024
	// Transceiver Management)
	encodingOffset = 0x8B

	// Nominal signaling rate, units of 100 MBd. For rate
	// > 25.4 GBd, set this to FFh and use Byte 222.
	signalingRateOffset = 0x8C

	//  Nominal baud rate per channel, units of 250 MBd.
	//  Complements Byte 140. See Table 6-25.
	signalingRateExtendedOffset = 0xDE

	// Tags for extended rate select compliance. See
	// Table 6-17.
	extendedRateSelectOffset = 0x8D

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for SMF fiber in km. A value
	// of 1 shall be used for reaches from 0 to 1 km.
	lengthSmfOffset = 0x8E

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for EBW 50/125 um fiber (OM3),
	// units of 2 m
	lengthOM3Offset = 0x8F

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for 50/125 um fiber (OM2),
	// units of 1 m
	lengthOM2Offset = 0x90

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for 62.5/125 um fiber (OM1),
	// units of 1 m, or copper cable attenuation in dB at
	// 25.78 GHz.
	lengthOM1offset = 0x91

	// Length of passive or active cable assembly (units of
	// 1 m) or link length supported at the bit rate in byte
	// 140 or page 00h byte 222, for OM4 50/125 um fiber
	// (units of 2 m) as indicated by Byte 147. See 6.3.12.
	lengthOM4orActivePassiveCableOffset = 0x92

	// Device technology (Table 6-18 and Table 6-19).
	deviceTechnologyOffset = 0x93

	// Free side device vendor name (ASCII)
	vendorNameStartOffset     = 0x94
	vendorNameEndOffset = 0xA3

	// Extended Module codes for InfiniBand (See Table
	// 6-20 )
	extendedModuleOffset = 0xA4

	// Free side device vendor IEEE company ID
	vendorOuiOffset = 0xA5

	// Part number provided by free side device
	// vendor(ASCII)
	vendorPnStartOffset     = 0xA8
	vendorPnEndOffset = 0xB7

	// Revision level for part number provided by the
	// vendor(ASCII)
	vendorRevStartOffset     = 0xB8
	vendorRevEndOffset = 0xB9

	// Nominal laser wavelength (wavelength=value/20 in
	// nm) or copper cable attenuation in dB at 2.5 GHz
	// (Byte 186) and 5.0 GHz (Byte 187)
	wavelengthOffset = 0xBA

	// Nominal laser wavelength (wavelength=value/20 in
	// nm) or copper cable attenuation in dB at 2.5 GHz
	// (Byte 186) and 5.0 GHz (Byte 187)
	copperAttenuation2dot5GHzOffset = 0xBA

	// Nominal laser wavelength (wavelength=value/20 in
	// nm) or copper cable attenuation in dB at 2.5 GHz
	// (Byte 186) and 5.0 GHz (Byte 187)
	copperAttenuation5GHzOffset = 0xBB

	// The range of laser wavelength (+/- value) from
	// nominal wavelength. (wavelength Tol. =value/200
	// in nm) or copper cable attenuation in dB at 7.0 GHz
	// (Byte 188) and 12.9 GHz (Byte 189)
	wavelengthToleranceOffset = 0xBC

	// The range of laser wavelength (+/- value) from
	// nominal wavelength. (wavelength Tol. =value/200
	// in nm) or copper cable attenuation in dB at 7.0 GHz
	// (Byte 188) and 12.9 GHz (Byte 189)
	copperAttenuation7GHzOffset = 0xBC

	// The range of laser wavelength (+/- value) from
	// nominal wavelength. (wavelength Tol. =value/200
	// in nm) or copper cable attenuation in dB at 7.0 GHz
	// (Byte 188) and 12.9 GHz (Byte 189)
	copperAttenuation12dot9GHzOffset = 0xBD

	// Maximum case temperature
	// Note: The standard does not specify how to parse this field
	maxCaseTemperatureOffset = 0xBE

	// Extended Specification Compliance Codes (See SFF-
	// 8024)
	linkCodesOffset = 0xC0

	// Optional features implemented. See Table 6-21.
	optionsOffset = 0xC1

	// Serial number provided by vendor (ASCII)
	vendorSnStartOffset     = 0xC4
	vendorSnEndOffset = 0xD3

	// Vendor's manufacturing date code
	vendorDateCodeStartOffset     = 0xD4
	vendorDateCodeEndOffset = 0xDB

	// Indicates which type of diagnostic monitoring is
	// implemented (if any) in the free side device. Bit 1,0
	// Reserved. See Table 6-23.
	diagnosticMonitoringTypeOffset = 0xDC

	// Indicates which optional enhanced features are
	// implemented in the free side device. See Table
	// 6-24.
	enhancedOptionsOffset = 0xDD

	/* Upper Page 03h (Optional) */
	// Free Side Device Thresholds
	thresholdsOffset = 0x200
)

// EEPROM implementation is based on SFF-8636 Rev 2.10a
type EEPROM struct {
	/* Lower Page */
	Identifier               sff8024.Identifier
	StatusIndicators         *StatusIndicators
	InterruptFlags           *InterruptFlags
	FreeSideMonitors         *FreeSideMonitors
	ChannelMonitors          *ChannelMonitors
	Control                  *Control
	InterruptMasks           *InterruptMasks
	FreeSideDeviceProperties *FreeSideDeviceProperties

	/* Upper Page 00h */
	// Identifier "shall" be the same as Identifier
	Identifier1                     sff8024.Identifier
	ExtendedIdentifier              *ExtendedIdentifier
	ConnectorType                   sff8024.ConnectorType
	SpecificationCompliance         SpecificationCompliance
	Encoding                        Encoding
	SignalingRate                   int
	ExtendedRateSelectCompliance    ExtendedRateSelectCompliance
	LengthSMF                       int
	LengthOM3                       int
	LengthOM2                       int
	LengthOM1                       int
	LengthOM4ActiveOrPassiveCable   int
	DeviceTechnology                *DeviceTechnology
	VendorName                      string
	ExtendedModuleCodeValues        *ExtendedModuleCodeValues
	VendorOUI                       eeprom.OUI
	VendorPN                        string
	VendorRev                       string
	Wavelength                      float64
	WavelengthTolerance             float64
	CopperAttenuation2_5GHz         byte
	CopperAttenuation5GHz           byte
	CopperAttenuation7GHz           byte
	CopperAttenuation12_9GHz        byte
	MaxCaseTemperature              byte
	ExtendedSpecificationCompliance sff8024.ExtendedSpecificationCompliance
	Options                         *Options
	VendorSN                        string
	DateCode                        string
	DiagnosticMonitoringType        *DiagnosticMonitoringType
	EnhancedOptions                 *EnhancedOptions

	/* Upper Page 03h (optional) */
	Thresholds *Thresholds
}

// NewEEPROM parses a byte slice of at least length 512 into a new EEPROM instance 
func NewEEPROM(raw []byte) (*EEPROM, error) {
	if len(raw) < 512 {
		return nil, errors.New("SFF-8636 requires EEPROM to be at least of 512 bytes length")
	}

	e := &EEPROM{
		/* Lower Page */
		Identifier: sff8024.Identifier(raw[identifierOffset]),
		StatusIndicators: NewStatusIndicators([2]byte{
			raw[statusIndicatorsOffset],
			raw[statusIndicatorsOffset+1]}),
		InterruptFlags: NewInterruptFlags([19]byte{
			raw[interruptFlagsOffset+0],
			raw[interruptFlagsOffset+1],
			raw[interruptFlagsOffset+2],
			raw[interruptFlagsOffset+3],
			raw[interruptFlagsOffset+4],
			raw[interruptFlagsOffset+5],
			raw[interruptFlagsOffset+6],
			raw[interruptFlagsOffset+7],
			raw[interruptFlagsOffset+8],
			raw[interruptFlagsOffset+9],
			raw[interruptFlagsOffset+10],
			raw[interruptFlagsOffset+11],
			raw[interruptFlagsOffset+12],
			raw[interruptFlagsOffset+13],
			raw[interruptFlagsOffset+14],
			raw[interruptFlagsOffset+15],
			raw[interruptFlagsOffset+16],
			raw[interruptFlagsOffset+17],
			raw[interruptFlagsOffset+18],
		}),
		FreeSideMonitors: NewFreeSideMonitors([12]byte{
			raw[freeSideDeviceMonitorsOffset+0],
			raw[freeSideDeviceMonitorsOffset+1],
			raw[freeSideDeviceMonitorsOffset+2],
			raw[freeSideDeviceMonitorsOffset+3],
			raw[freeSideDeviceMonitorsOffset+4],
			raw[freeSideDeviceMonitorsOffset+5],
			raw[freeSideDeviceMonitorsOffset+6],
			raw[freeSideDeviceMonitorsOffset+7],
			raw[freeSideDeviceMonitorsOffset+8],
			raw[freeSideDeviceMonitorsOffset+9],
			raw[freeSideDeviceMonitorsOffset+10],
			raw[freeSideDeviceMonitorsOffset+11],
		}),
		ChannelMonitors: NewChannelMonitors([48]byte{
			raw[channelMonitorsOffset+0],
			raw[channelMonitorsOffset+1],
			raw[channelMonitorsOffset+2],
			raw[channelMonitorsOffset+3],
			raw[channelMonitorsOffset+4],
			raw[channelMonitorsOffset+5],
			raw[channelMonitorsOffset+6],
			raw[channelMonitorsOffset+7],
			raw[channelMonitorsOffset+8],
			raw[channelMonitorsOffset+9],
			raw[channelMonitorsOffset+10],
			raw[channelMonitorsOffset+11],
			raw[channelMonitorsOffset+12],
			raw[channelMonitorsOffset+13],
			raw[channelMonitorsOffset+14],
			raw[channelMonitorsOffset+15],
			raw[channelMonitorsOffset+16],
			raw[channelMonitorsOffset+17],
			raw[channelMonitorsOffset+18],
			raw[channelMonitorsOffset+19],
			raw[channelMonitorsOffset+20],
			raw[channelMonitorsOffset+21],
			raw[channelMonitorsOffset+22],
			raw[channelMonitorsOffset+23],
			raw[channelMonitorsOffset+24],
			raw[channelMonitorsOffset+25],
			raw[channelMonitorsOffset+26],
			raw[channelMonitorsOffset+27],
			raw[channelMonitorsOffset+28],
			raw[channelMonitorsOffset+29],
			raw[channelMonitorsOffset+30],
			raw[channelMonitorsOffset+31],
			raw[channelMonitorsOffset+32],
			raw[channelMonitorsOffset+33],
			raw[channelMonitorsOffset+34],
			raw[channelMonitorsOffset+35],
			raw[channelMonitorsOffset+36],
			raw[channelMonitorsOffset+37],
			raw[channelMonitorsOffset+38],
			raw[channelMonitorsOffset+39],
			raw[channelMonitorsOffset+40],
			raw[channelMonitorsOffset+41],
			raw[channelMonitorsOffset+42],
			raw[channelMonitorsOffset+43],
			raw[channelMonitorsOffset+44],
			raw[channelMonitorsOffset+45],
			raw[channelMonitorsOffset+46],
			raw[channelMonitorsOffset+47],
		}),
		Control: NewControl([14]byte{
			raw[controlOffset+0],
			raw[controlOffset+1],
			raw[controlOffset+2],
			raw[controlOffset+3],
			raw[controlOffset+4],
			raw[controlOffset+5],
			raw[controlOffset+6],
			raw[controlOffset+7],
			raw[controlOffset+8],
			raw[controlOffset+9],
			raw[controlOffset+10],
			raw[controlOffset+11],
			raw[controlOffset+12],
			raw[controlOffset+13],
		}),
		InterruptMasks: NewInterruptMasks([6]byte{
			raw[freeSideInterruptMasksOffset+0],
			raw[freeSideInterruptMasksOffset+1],
			raw[freeSideInterruptMasksOffset+2],
			raw[freeSideInterruptMasksOffset+3],
			raw[freeSideInterruptMasksOffset+4],
			raw[freeSideInterruptMasksOffset+5],
		}),
		FreeSideDeviceProperties: NewFreeSideDeviceProperties([10]byte{
			raw[freeSideDevicePropertiesOffset+0],
			raw[freeSideDevicePropertiesOffset+1],
			raw[freeSideDevicePropertiesOffset+2],
			raw[freeSideDevicePropertiesOffset+3],
			raw[freeSideDevicePropertiesOffset+4],
			raw[freeSideDevicePropertiesOffset+5],
			raw[freeSideDevicePropertiesOffset+6],
			raw[freeSideDevicePropertiesOffset+7],
			raw[freeSideDevicePropertiesOffset+8],
			raw[freeSideDevicePropertiesOffset+9],
		}),
		/* Upper Page 00h */
		Identifier1:        sff8024.Identifier(raw[identifierOffset1]),
		ExtendedIdentifier: NewExtendedIdentifier(raw[extendedIdentifierOffset]),
		ConnectorType:      sff8024.ConnectorType(raw[connectorTypeOffset]),
		SpecificationCompliance: NewSpecificationCompliance([8]byte{
			raw[specificationCompliance+0],
			raw[specificationCompliance+1],
			raw[specificationCompliance+2],
			raw[specificationCompliance+3],
			raw[specificationCompliance+4],
			raw[specificationCompliance+5],
			raw[specificationCompliance+6],
			raw[specificationCompliance+7],
		}),
		Encoding: Encoding(raw[encodingOffset]),
		SignalingRate: func() int {
			if raw[signalingRateOffset] == 255 {
				return int(raw[signalingRateExtendedOffset]) * 250 * 1000000
			}
			return int(raw[signalingRateOffset]) * 100 * 1000000
		}(),
		ExtendedRateSelectCompliance:  ExtendedRateSelectCompliance(raw[extendedRateSelectOffset]),
		LengthSMF:                     int(raw[lengthSmfOffset]) * 1000,
		LengthOM3:                     int(raw[lengthOM3Offset]) * 2,
		LengthOM2:                     int(raw[lengthOM2Offset]),
		LengthOM1:                     int(raw[lengthOM1offset]),
		LengthOM4ActiveOrPassiveCable: int(raw[lengthOM4orActivePassiveCableOffset]),
		DeviceTechnology:              NewDeviceTechnology(raw[deviceTechnologyOffset]),
		VendorName:                    strings.Trim(parseString(raw[vendorNameStartOffset:vendorNameEndOffset+1]), " "),
		ExtendedModuleCodeValues:      NewExtendedModuleCodeValues(raw[extendedModuleOffset]),
		VendorOUI: eeprom.NewOUI([3]byte{
			raw[vendorOuiOffset+0],
			raw[vendorOuiOffset+1],
			raw[vendorOuiOffset+2],
		}),
		VendorPN:  strings.Trim(parseString(raw[vendorPnStartOffset:vendorPnEndOffset+1]), " "),
		VendorRev: strings.Trim(parseString(raw[vendorRevStartOffset:vendorRevEndOffset+1]), " "),
		Wavelength: parseWavelength(
			raw[wavelengthOffset+0],
			raw[wavelengthOffset+1],
		),
		WavelengthTolerance: parseWavelengthTolerance(
			raw[wavelengthToleranceOffset+0],
			raw[wavelengthToleranceOffset+1],
		),
		CopperAttenuation2_5GHz:         raw[copperAttenuation2dot5GHzOffset],
		CopperAttenuation5GHz:           raw[copperAttenuation5GHzOffset],
		CopperAttenuation7GHz:           raw[copperAttenuation7GHzOffset],
		CopperAttenuation12_9GHz:        raw[copperAttenuation12dot9GHzOffset],
		MaxCaseTemperature:              raw[maxCaseTemperatureOffset],
		ExtendedSpecificationCompliance: sff8024.ExtendedSpecificationCompliance(raw[linkCodesOffset]),
		Options: NewOptions([3]byte{
			raw[optionsOffset+0],
			raw[optionsOffset+1],
			raw[optionsOffset+2],
		}),
		VendorSN:                 strings.Trim(parseString(raw[vendorSnStartOffset:vendorSnEndOffset+1]), " "),
		DateCode:                 parseString(raw[vendorDateCodeStartOffset : vendorDateCodeEndOffset+1]),
		DiagnosticMonitoringType: NewDiagnosticMonitoringType(raw[diagnosticMonitoringTypeOffset]),
		EnhancedOptions:          NewEnhancedOptions(raw[enhancedOptionsOffset]),
	}
	/* Upper Page 03h (Optional) */
	if len(raw) >= 0x248 {
		e.Thresholds = NewThresholds([72]byte{
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
			raw[thresholdsOffset+40],
			raw[thresholdsOffset+41],
			raw[thresholdsOffset+42],
			raw[thresholdsOffset+43],
			raw[thresholdsOffset+44],
			raw[thresholdsOffset+45],
			raw[thresholdsOffset+46],
			raw[thresholdsOffset+47],
			raw[thresholdsOffset+48],
			raw[thresholdsOffset+49],
			raw[thresholdsOffset+50],
			raw[thresholdsOffset+51],
			raw[thresholdsOffset+52],
			raw[thresholdsOffset+53],
			raw[thresholdsOffset+54],
			raw[thresholdsOffset+55],
			raw[thresholdsOffset+56],
			raw[thresholdsOffset+57],
			raw[thresholdsOffset+58],
			raw[thresholdsOffset+59],
			raw[thresholdsOffset+60],
			raw[thresholdsOffset+61],
			raw[thresholdsOffset+62],
			raw[thresholdsOffset+63],
			raw[thresholdsOffset+64],
			raw[thresholdsOffset+65],
			raw[thresholdsOffset+66],
			raw[thresholdsOffset+67],
			raw[thresholdsOffset+68],
			raw[thresholdsOffset+69],
			raw[thresholdsOffset+70],
			raw[thresholdsOffset+71],
		})
	}
	if e.SpecificationCompliance.IsNonOpticalImplementation() {
		e.Wavelength = 0
		e.WavelengthTolerance = 0
	} else {
		e.CopperAttenuation2_5GHz = 0
		e.CopperAttenuation5GHz = 0
		e.CopperAttenuation7GHz = 0
		e.CopperAttenuation12_9GHz = 0
	}

	return e, nil
}
