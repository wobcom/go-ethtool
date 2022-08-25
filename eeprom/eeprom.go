package eeprom

import (
	"github.com/wobcom/go-ethtool/eeprom/sff8024"
	"time"
)

// Type Standard the interface's EEPROM complies with
type Type uint32

// EEPROM Standards for plug in modules
const (
	TypeSFF8079 Type = 0x01
	TypeSFF8472 Type = 0x02
	TypeSFF8636 Type = 0x03
	TypeSFF8436 Type = 0x04
)

func (e Type) String() string {
	return map[Type]string{
		TypeSFF8079: "SFF-8079",
		TypeSFF8472: "SFF-8472",
		TypeSFF8636: "SFF-8636",
		TypeSFF8436: "SFF-8436",
	}[e]
}

// EEPROM is a unified interface for eeproms complying with different standards
type EEPROM interface {
	GetIdentifier() sff8024.Identifier
	GetConnectorType() sff8024.ConnectorType
	GetEncoding() string
	GetPowerClass() PowerClass
	GetSignalingRate() float64
	GetSupportedLinkLengths() map[string]float64
	GetVendorName() string
	GetVendorPN() string
	GetVendorRev() string
	GetVendorSN() string
	GetVendorOUI() OUI
	GetDateCode() time.Time
	GetWavelength() float64
	GetLasers() []Laser
	SupportsMonitoring() bool
	GetModuleTemperature() (Measurement, error)
	GetModuleVoltage() (Measurement, error)
}

// Laser may provide realtime monitoring information
type Laser interface {
	SupportsMonitoring() bool
	GetBias() (Measurement, error)
	GetTxPower() (Measurement, error)
	GetRxPower() (Measurement, error)
}

// Measurement a value read from a sensor, may provide alarm thresholds
type Measurement interface {
	GetValue() float64
	GetUnit() string
	SupportsThresholds() bool
	GetAlarmThresholds() (AlarmThresholds, error)
}

// AlarmThresholds warning / alarm thresholds for a given reading
type AlarmThresholds interface {
	GetHighAlarm() float64
	GetHighWarning() float64
	GetLowAlarm() float64
	GetLowWarning() float64
}
