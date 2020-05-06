package SFF8636

import (
	"errors"
	"gitlab.com/wobcom/golang-ethtool/eeprom"
	"gitlab.com/wobcom/golang-ethtool/eeprom/sff8024"
	"strings"
)

/* Memory offsets as defined in SFF-8636 Rev 2.10a (September 24, 2019) */
const (
	/* Lower Page 00h (Table 6-1) */

	// Identifier (See SFF-8024 Transceiver Management)
	SFF8636_IDENTIFIER = 0x00
	// Status
	SFF8636_STATUS_INDICATORS = 0x01
	// Interrupt Flags
	SFF8636_INTERRUPT_FLAGS = 0x03
	// Free Side Device Monitors
	SFF8636_FREE_SIDE_DEVICE_MONITORS = 0x16
	// Channel Monitors
	SFF8636_CHANNEL_MONITORS = 0x22

	// 0x52 - 0x55 reserved

	// Control
	SFF8636_CONTROL = 0x56
	// Free Side Device and Channel Masks
	SFF8636_FREE_SIDE_INTERRUPT_MASKS = 0x64
	// Free Side Device Properties
	SFF8636_FREE_SIDE_DEVICE_PROPERTIES = 0x6B

	/* Upper Page 00h (Table 6-14) */
	// Identifier Type of free side device (See SFF-8024 Transceiver Management)
	// Note: Should read the same as SFF8636_IDENTIFIER
	SFF8636_IDENTIFIER1 = 0x80

	// Extended Identifier of free side device. Includes
	// power classes, CLEI codes, CDR capability (See
	// Table 6-15)
	SFF8636_EXTENDED_IDENTIFIER = 0x81

	// Code for media connector type (See SFF-8024
	// Transceiver Management)
	SFF8636_CONNECTOR_TYPE = 0x82

	// Code for electronic or optical compatibility (See
	// Table 6-16)
	SFF8636_SPECIFICATION_COMPLIANCE = 0x83

	// Code for serial encoding algorithm. (See SFF-8024
	// Transceiver Management)
	SFF8636_ENCODING = 0x8B

	// Nominal signaling rate, units of 100 MBd. For rate
	// > 25.4 GBd, set this to FFh and use Byte 222.
	SFF8636_SIGNALLING_RATE = 0x8C

	//  Nominal baud rate per channel, units of 250 MBd.
	//  Complements Byte 140. See Table 6-25.
	SFF8636_SIGNALLING_RATE_EXTENDED = 0xDE

	// Tags for extended rate select compliance. See
	// Table 6-17.
	SFF8636_EXTENDED_RATE_SELECT = 0x8D

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for SMF fiber in km. A value
	// of 1 shall be used for reaches from 0 to 1 km.
	SFF8636_LENGTH_SMF = 0x8E

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for EBW 50/125 um fiber (OM3),
	// units of 2 m
	SFF8636_LENGTH_OM3 = 0x8F

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for 50/125 um fiber (OM2),
	// units of 1 m
	SFF8636_LENGTH_OM2 = 0x90

	// Link length supported at the bit rate in byte 140 or
	// page 00h byte 222, for 62.5/125 um fiber (OM1),
	// units of 1 m, or copper cable attenuation in dB at
	// 25.78 GHz.
	SFF8636_LENGTH_OM1 = 0x91

	// Length of passive or active cable assembly (units of
	// 1 m) or link length supported at the bit rate in byte
	// 140 or page 00h byte 222, for OM4 50/125 um fiber
	// (units of 2 m) as indicated by Byte 147. See 6.3.12.
	SFF8636_LENGTH_OM4_ACTIVE_PASSIVE_CABLE = 0x92

	// Device technology (Table 6-18 and Table 6-19).
	SFF8636_DEVICE_TECHNOLOGY = 0x93

	// Free side device vendor name (ASCII)
	SFF8636_VENDOR_NAME     = 0x94
	SFF8636_VENDOR_NAME_END = 0xA3

	// Extended Module codes for InfiniBand (See Table
	// 6-20 )
	SFF8636_EXTENDED_MODULE = 0xA4

	// Free side device vendor IEEE company ID
	SFF8636_VENDOR_OUI = 0xA5

	// Part number provided by free side device
	// vendor(ASCII)
	SFF8636_VENDOR_PN     = 0xA8
	SFF8636_VENDOR_PN_END = 0xB7

	// Revision level for part number provided by the
	// vendor(ASCII)
	SFF8636_VENDOR_REV     = 0xB8
	SFF8636_VENDOR_REV_END = 0xB9

	// Nominal laser wavelength (wavelength=value/20 in
	// nm) or copper cable attenuation in dB at 2.5 GHz
	// (Byte 186) and 5.0 GHz (Byte 187)
	SFF8636_WAVELENGTH = 0xBA

	// Nominal laser wavelength (wavelength=value/20 in
	// nm) or copper cable attenuation in dB at 2.5 GHz
	// (Byte 186) and 5.0 GHz (Byte 187)
	SFF8636_COPPER_ATTENUATION_2_5_GHZ = 0xBA

	// Nominal laser wavelength (wavelength=value/20 in
	// nm) or copper cable attenuation in dB at 2.5 GHz
	// (Byte 186) and 5.0 GHz (Byte 187)
	SFF8636_COPPER_ATTENUATION_5_GHZ = 0xBB

	// The range of laser wavelength (+/- value) from
	// nominal wavelength. (wavelength Tol. =value/200
	// in nm) or copper cable attenuation in dB at 7.0 GHz
	// (Byte 188) and 12.9 GHz (Byte 189)
	SFF8636_WAVELENGTH_TOLERANCE = 0xBC

	// The range of laser wavelength (+/- value) from
	// nominal wavelength. (wavelength Tol. =value/200
	// in nm) or copper cable attenuation in dB at 7.0 GHz
	// (Byte 188) and 12.9 GHz (Byte 189)
	SFF8636_COPPER_ATTENUATION_7_GHZ = 0xBC

	// The range of laser wavelength (+/- value) from
	// nominal wavelength. (wavelength Tol. =value/200
	// in nm) or copper cable attenuation in dB at 7.0 GHz
	// (Byte 188) and 12.9 GHz (Byte 189)
	SFF8636_COPPER_ATTENUATION_12_9_GHZ = 0xBD

	// Maximum case temperature
	// Note: The standard does not specify how to parse this field
	SFF8636_MAX_CASE_TEMPERATURE = 0xBE

	// Extended Specification Compliance Codes (See SFF-
	// 8024)
	SFF8636_LINK_CODES = 0xC0

	// Optional features implemented. See Table 6-21.
	SFF8636_OPTIONS = 0xC1

	// Serial number provided by vendor (ASCII)
	SFF8636_VENDOR_SN     = 0xC4
	SFF8636_VENDOR_SN_END = 0xD3

	// Vendor's manufacturing date code
	SFF8636_VENDOR_DATE_CODE     = 0xD4
	SFF8636_VENDOR_DATE_CODE_END = 0xDB

	// Indicates which type of diagnostic monitoring is
	// implemented (if any) in the free side device. Bit 1,0
	// Reserved. See Table 6-23.
	SFF8636_DIAGNOSTIC_MONITORING_TYPE = 0xDC

	// Indicates which optional enhanced features are
	// implemented in the free side device. See Table
	// 6-24.
	SFF8636_ENHANCED_OPTIONS = 0xDD

	/* Upper Page 03h (Optional) */
	// Free Side Device Thresholds
	SFF8636_THRESHOLDS = 0x200
)

// EEPROM implementation is based on SFF-8636 Rev 2.10a
type EEPROM struct {
	/* Lower Page */
	Identifier               SFF8024.Identifier
	StatusIndicators         *StatusIndicators
	InterruptFlags           *InterruptFlags
	FreeSideMonitors         *FreeSideMonitors
	ChannelMonitors          *ChannelMonitors
	Control                  *Control
	InterruptMasks           *InterruptMasks
	FreeSideDeviceProperties *FreeSideDeviceProperties

	/* Upper Page 00h */
	// Identifier "shall" be the same as Identifier
	Identifier1                     SFF8024.Identifier
	ExtendedIdentifier              *ExtendedIdentifier
	ConnectorType                   SFF8024.ConnectorType
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
	ExtendedSpecificationCompliance SFF8024.ExtendedSpecificationCompliance
	Options                         *Options
	VendorSN                        string
	DateCode                        string
	DiagnosticMonitoringType        *DiagnosticMonitoringType
	EnhancedOptions                 *EnhancedOptions

	/* Upper Page 03h (optional) */
	Thresholds *Thresholds
}

func NewEEPROM(raw []byte) (*EEPROM, error) {
	if len(raw) < 512 {
		return nil, errors.New("SFF-8636 requires EEPROM to be at least of 512 bytes length")
	}

	e := &EEPROM{
		/* Lower Page */
		Identifier: SFF8024.Identifier(raw[SFF8636_IDENTIFIER]),
		StatusIndicators: NewStatusIndicators([2]byte{
			raw[SFF8636_STATUS_INDICATORS],
			raw[SFF8636_STATUS_INDICATORS+1]}),
		InterruptFlags: NewInterruptFlags([19]byte{
			raw[SFF8636_INTERRUPT_FLAGS+0],
			raw[SFF8636_INTERRUPT_FLAGS+1],
			raw[SFF8636_INTERRUPT_FLAGS+2],
			raw[SFF8636_INTERRUPT_FLAGS+3],
			raw[SFF8636_INTERRUPT_FLAGS+4],
			raw[SFF8636_INTERRUPT_FLAGS+5],
			raw[SFF8636_INTERRUPT_FLAGS+6],
			raw[SFF8636_INTERRUPT_FLAGS+7],
			raw[SFF8636_INTERRUPT_FLAGS+8],
			raw[SFF8636_INTERRUPT_FLAGS+9],
			raw[SFF8636_INTERRUPT_FLAGS+10],
			raw[SFF8636_INTERRUPT_FLAGS+11],
			raw[SFF8636_INTERRUPT_FLAGS+12],
			raw[SFF8636_INTERRUPT_FLAGS+13],
			raw[SFF8636_INTERRUPT_FLAGS+14],
			raw[SFF8636_INTERRUPT_FLAGS+15],
			raw[SFF8636_INTERRUPT_FLAGS+16],
			raw[SFF8636_INTERRUPT_FLAGS+17],
			raw[SFF8636_INTERRUPT_FLAGS+18],
		}),
		FreeSideMonitors: NewFreeSideMonitors([12]byte{
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+0],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+1],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+2],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+3],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+4],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+5],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+6],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+7],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+8],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+9],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+10],
			raw[SFF8636_FREE_SIDE_DEVICE_MONITORS+11],
		}),
		ChannelMonitors: NewChannelMonitors([48]byte{
			raw[SFF8636_CHANNEL_MONITORS+0],
			raw[SFF8636_CHANNEL_MONITORS+1],
			raw[SFF8636_CHANNEL_MONITORS+2],
			raw[SFF8636_CHANNEL_MONITORS+3],
			raw[SFF8636_CHANNEL_MONITORS+4],
			raw[SFF8636_CHANNEL_MONITORS+5],
			raw[SFF8636_CHANNEL_MONITORS+6],
			raw[SFF8636_CHANNEL_MONITORS+7],
			raw[SFF8636_CHANNEL_MONITORS+8],
			raw[SFF8636_CHANNEL_MONITORS+9],
			raw[SFF8636_CHANNEL_MONITORS+10],
			raw[SFF8636_CHANNEL_MONITORS+11],
			raw[SFF8636_CHANNEL_MONITORS+12],
			raw[SFF8636_CHANNEL_MONITORS+13],
			raw[SFF8636_CHANNEL_MONITORS+14],
			raw[SFF8636_CHANNEL_MONITORS+15],
			raw[SFF8636_CHANNEL_MONITORS+16],
			raw[SFF8636_CHANNEL_MONITORS+17],
			raw[SFF8636_CHANNEL_MONITORS+18],
			raw[SFF8636_CHANNEL_MONITORS+19],
			raw[SFF8636_CHANNEL_MONITORS+20],
			raw[SFF8636_CHANNEL_MONITORS+21],
			raw[SFF8636_CHANNEL_MONITORS+22],
			raw[SFF8636_CHANNEL_MONITORS+23],
			raw[SFF8636_CHANNEL_MONITORS+24],
			raw[SFF8636_CHANNEL_MONITORS+25],
			raw[SFF8636_CHANNEL_MONITORS+26],
			raw[SFF8636_CHANNEL_MONITORS+27],
			raw[SFF8636_CHANNEL_MONITORS+28],
			raw[SFF8636_CHANNEL_MONITORS+29],
			raw[SFF8636_CHANNEL_MONITORS+30],
			raw[SFF8636_CHANNEL_MONITORS+31],
			raw[SFF8636_CHANNEL_MONITORS+32],
			raw[SFF8636_CHANNEL_MONITORS+33],
			raw[SFF8636_CHANNEL_MONITORS+34],
			raw[SFF8636_CHANNEL_MONITORS+35],
			raw[SFF8636_CHANNEL_MONITORS+36],
			raw[SFF8636_CHANNEL_MONITORS+37],
			raw[SFF8636_CHANNEL_MONITORS+38],
			raw[SFF8636_CHANNEL_MONITORS+39],
			raw[SFF8636_CHANNEL_MONITORS+40],
			raw[SFF8636_CHANNEL_MONITORS+41],
			raw[SFF8636_CHANNEL_MONITORS+42],
			raw[SFF8636_CHANNEL_MONITORS+43],
			raw[SFF8636_CHANNEL_MONITORS+44],
			raw[SFF8636_CHANNEL_MONITORS+45],
			raw[SFF8636_CHANNEL_MONITORS+46],
			raw[SFF8636_CHANNEL_MONITORS+47],
		}),
		Control: NewControl([14]byte{
			raw[SFF8636_CONTROL+0],
			raw[SFF8636_CONTROL+1],
			raw[SFF8636_CONTROL+2],
			raw[SFF8636_CONTROL+3],
			raw[SFF8636_CONTROL+4],
			raw[SFF8636_CONTROL+5],
			raw[SFF8636_CONTROL+6],
			raw[SFF8636_CONTROL+7],
			raw[SFF8636_CONTROL+8],
			raw[SFF8636_CONTROL+9],
			raw[SFF8636_CONTROL+10],
			raw[SFF8636_CONTROL+11],
			raw[SFF8636_CONTROL+12],
			raw[SFF8636_CONTROL+13],
		}),
		InterruptMasks: NewInterruptMasks([6]byte{
			raw[SFF8636_FREE_SIDE_INTERRUPT_MASKS+0],
			raw[SFF8636_FREE_SIDE_INTERRUPT_MASKS+1],
			raw[SFF8636_FREE_SIDE_INTERRUPT_MASKS+2],
			raw[SFF8636_FREE_SIDE_INTERRUPT_MASKS+3],
			raw[SFF8636_FREE_SIDE_INTERRUPT_MASKS+4],
			raw[SFF8636_FREE_SIDE_INTERRUPT_MASKS+5],
		}),
		FreeSideDeviceProperties: NewFreeSideDeviceProperties([10]byte{
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+0],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+1],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+2],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+3],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+4],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+5],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+6],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+7],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+8],
			raw[SFF8636_FREE_SIDE_DEVICE_PROPERTIES+9],
		}),
		/* Upper Page 00h */
		Identifier1:        SFF8024.Identifier(raw[SFF8636_IDENTIFIER1]),
		ExtendedIdentifier: NewExtendedIdentifier(raw[SFF8636_EXTENDED_IDENTIFIER]),
		ConnectorType:      SFF8024.ConnectorType(raw[SFF8636_CONNECTOR_TYPE]),
		SpecificationCompliance: NewSpecificationCompliance([8]byte{
			raw[SFF8636_SPECIFICATION_COMPLIANCE+0],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+1],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+2],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+3],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+4],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+5],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+6],
			raw[SFF8636_SPECIFICATION_COMPLIANCE+7],
		}),
		Encoding: Encoding(raw[SFF8636_ENCODING]),
		SignalingRate: func() int {
			if raw[SFF8636_SIGNALLING_RATE] == 255 {
				return int(raw[SFF8636_SIGNALLING_RATE_EXTENDED]) * 250 * 1000000
			}
			return int(raw[SFF8636_SIGNALLING_RATE]) * 100 * 1000000
		}(),
		ExtendedRateSelectCompliance:  ExtendedRateSelectCompliance(raw[SFF8636_EXTENDED_RATE_SELECT]),
		LengthSMF:                     int(raw[SFF8636_LENGTH_SMF]) * 1000,
		LengthOM3:                     int(raw[SFF8636_LENGTH_OM3]) * 2,
		LengthOM2:                     int(raw[SFF8636_LENGTH_OM2]),
		LengthOM1:                     int(raw[SFF8636_LENGTH_OM1]),
		LengthOM4ActiveOrPassiveCable: int(raw[SFF8636_LENGTH_OM4_ACTIVE_PASSIVE_CABLE]),
		DeviceTechnology:              NewDeviceTechnology(raw[SFF8636_DEVICE_TECHNOLOGY]),
		VendorName:                    strings.Trim(parseString(raw[SFF8636_VENDOR_NAME:SFF8636_VENDOR_NAME_END+1]), " "),
		ExtendedModuleCodeValues:      NewExtendedModuleCodeValues(raw[SFF8636_EXTENDED_MODULE]),
		VendorOUI: eeprom.NewOUI([3]byte{
			raw[SFF8636_VENDOR_OUI+0],
			raw[SFF8636_VENDOR_OUI+1],
			raw[SFF8636_VENDOR_OUI+2],
		}),
		VendorPN:  strings.Trim(parseString(raw[SFF8636_VENDOR_PN:SFF8636_VENDOR_PN_END+1]), " "),
		VendorRev: strings.Trim(parseString(raw[SFF8636_VENDOR_REV:SFF8636_VENDOR_REV_END+1]), " "),
		Wavelength: parseWavelength(
			raw[SFF8636_WAVELENGTH+0],
			raw[SFF8636_WAVELENGTH+1],
		),
		WavelengthTolerance: parseWavelengthTolerance(
			raw[SFF8636_WAVELENGTH_TOLERANCE+0],
			raw[SFF8636_WAVELENGTH_TOLERANCE+1],
		),
		CopperAttenuation2_5GHz:         raw[SFF8636_COPPER_ATTENUATION_2_5_GHZ],
		CopperAttenuation5GHz:           raw[SFF8636_COPPER_ATTENUATION_5_GHZ],
		CopperAttenuation7GHz:           raw[SFF8636_COPPER_ATTENUATION_7_GHZ],
		CopperAttenuation12_9GHz:        raw[SFF8636_COPPER_ATTENUATION_12_9_GHZ],
		MaxCaseTemperature:              raw[SFF8636_MAX_CASE_TEMPERATURE],
		ExtendedSpecificationCompliance: SFF8024.ExtendedSpecificationCompliance(raw[SFF8636_LINK_CODES]),
		Options: NewOptions([3]byte{
			raw[SFF8636_OPTIONS+0],
			raw[SFF8636_OPTIONS+1],
			raw[SFF8636_OPTIONS+2],
		}),
		VendorSN:                 strings.Trim(parseString(raw[SFF8636_VENDOR_SN:SFF8636_VENDOR_SN_END+1]), " "),
		DateCode:                 parseString(raw[SFF8636_VENDOR_DATE_CODE : SFF8636_VENDOR_DATE_CODE_END+1]),
		DiagnosticMonitoringType: NewDiagnosticMonitoringType(raw[SFF8636_DIAGNOSTIC_MONITORING_TYPE]),
		EnhancedOptions:          NewEnhancedOptions(raw[SFF8636_ENHANCED_OPTIONS]),
	}
	/* Upper Page 03h (Optional) */
	if len(raw) >= 0x248 {
		e.Thresholds = NewThresholds([72]byte{
			raw[SFF8636_THRESHOLDS+0],
			raw[SFF8636_THRESHOLDS+1],
			raw[SFF8636_THRESHOLDS+2],
			raw[SFF8636_THRESHOLDS+3],
			raw[SFF8636_THRESHOLDS+4],
			raw[SFF8636_THRESHOLDS+5],
			raw[SFF8636_THRESHOLDS+6],
			raw[SFF8636_THRESHOLDS+7],
			raw[SFF8636_THRESHOLDS+8],
			raw[SFF8636_THRESHOLDS+9],
			raw[SFF8636_THRESHOLDS+10],
			raw[SFF8636_THRESHOLDS+11],
			raw[SFF8636_THRESHOLDS+12],
			raw[SFF8636_THRESHOLDS+13],
			raw[SFF8636_THRESHOLDS+14],
			raw[SFF8636_THRESHOLDS+15],
			raw[SFF8636_THRESHOLDS+16],
			raw[SFF8636_THRESHOLDS+17],
			raw[SFF8636_THRESHOLDS+18],
			raw[SFF8636_THRESHOLDS+19],
			raw[SFF8636_THRESHOLDS+20],
			raw[SFF8636_THRESHOLDS+21],
			raw[SFF8636_THRESHOLDS+22],
			raw[SFF8636_THRESHOLDS+23],
			raw[SFF8636_THRESHOLDS+24],
			raw[SFF8636_THRESHOLDS+25],
			raw[SFF8636_THRESHOLDS+26],
			raw[SFF8636_THRESHOLDS+27],
			raw[SFF8636_THRESHOLDS+28],
			raw[SFF8636_THRESHOLDS+29],
			raw[SFF8636_THRESHOLDS+30],
			raw[SFF8636_THRESHOLDS+31],
			raw[SFF8636_THRESHOLDS+32],
			raw[SFF8636_THRESHOLDS+33],
			raw[SFF8636_THRESHOLDS+34],
			raw[SFF8636_THRESHOLDS+35],
			raw[SFF8636_THRESHOLDS+36],
			raw[SFF8636_THRESHOLDS+37],
			raw[SFF8636_THRESHOLDS+38],
			raw[SFF8636_THRESHOLDS+39],
			raw[SFF8636_THRESHOLDS+40],
			raw[SFF8636_THRESHOLDS+41],
			raw[SFF8636_THRESHOLDS+42],
			raw[SFF8636_THRESHOLDS+43],
			raw[SFF8636_THRESHOLDS+44],
			raw[SFF8636_THRESHOLDS+45],
			raw[SFF8636_THRESHOLDS+46],
			raw[SFF8636_THRESHOLDS+47],
			raw[SFF8636_THRESHOLDS+48],
			raw[SFF8636_THRESHOLDS+49],
			raw[SFF8636_THRESHOLDS+50],
			raw[SFF8636_THRESHOLDS+51],
			raw[SFF8636_THRESHOLDS+52],
			raw[SFF8636_THRESHOLDS+53],
			raw[SFF8636_THRESHOLDS+54],
			raw[SFF8636_THRESHOLDS+55],
			raw[SFF8636_THRESHOLDS+56],
			raw[SFF8636_THRESHOLDS+57],
			raw[SFF8636_THRESHOLDS+58],
			raw[SFF8636_THRESHOLDS+59],
			raw[SFF8636_THRESHOLDS+60],
			raw[SFF8636_THRESHOLDS+61],
			raw[SFF8636_THRESHOLDS+62],
			raw[SFF8636_THRESHOLDS+63],
			raw[SFF8636_THRESHOLDS+64],
			raw[SFF8636_THRESHOLDS+65],
			raw[SFF8636_THRESHOLDS+66],
			raw[SFF8636_THRESHOLDS+67],
			raw[SFF8636_THRESHOLDS+68],
			raw[SFF8636_THRESHOLDS+69],
			raw[SFF8636_THRESHOLDS+70],
			raw[SFF8636_THRESHOLDS+71],
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
