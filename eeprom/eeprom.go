package eeprom

import (
	"gitlab.com/wobcom/ethtool/eeprom/sff8024"
	"time"
)

type EEPROMType uint32

/* EEPROM Standards for plug in modules */
const (
	ETH_MODULE_SFF_8079 EEPROMType = 0x01
	ETH_MODULE_SFF_8472 EEPROMType = 0x02
	ETH_MODULE_SFF_8636 EEPROMType = 0x03
	ETH_MODULE_SFF_8436 EEPROMType = 0x04
)

func (e EEPROMType) String() string {
	return map[EEPROMType]string{
		ETH_MODULE_SFF_8079: "SFF-8079",
		ETH_MODULE_SFF_8472: "SFF-8472",
		ETH_MODULE_SFF_8636: "SFF-8636",
		ETH_MODULE_SFF_8436: "SFF-8436",
	}[e]
}

type EEPROM interface {
	GetIdentifier() SFF8024.Identifier
	GetConnectorType() SFF8024.ConnectorType
	GetEncoding() Encoding
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

type Encoding interface {
	String() string
}

type Laser interface {
	SupportsMonitoring() bool
	GetBias() (Measurement, error)
	GetTxPower() (Measurement, error)
	GetRxPower() (Measurement, error)
}

type Measurement interface {
	GetValue() float64
	GetUnit() string
	SupportsThresholds() bool
	GetAlarmThresholds() (AlarmThresholds, error)
}

type AlarmThresholds interface {
	GetHighAlarm() float64
	GetHighWarning() float64
	GetLowAlarm() float64
	GetLowWarning() float64
}
