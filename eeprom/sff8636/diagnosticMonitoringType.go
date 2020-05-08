package sff8636

// DiagnosticMonitoringType indicators describing how diagnostic monitoring is implemented
type DiagnosticMonitoringType struct {
	TemperatureMonitoringImplemented     bool
	SupplyVoltageMonitoringImplemented   bool
	ReceivedPowerMeasurementsType        ReceivedPowerMeasurementsType
	TransmitterPowerMeasurementSupported bool
}

// ReceivedPowerMeasurementsType indicates how power measurements are implemented
type ReceivedPowerMeasurementsType bool

func (r ReceivedPowerMeasurementsType) String() string {
	if bool(r) {
		return "OMA"
	}
	return "Average Power"
}

// NewDiagnosticMonitoringType parses a bye into a new DiagnosticMonitoringType instance
func NewDiagnosticMonitoringType(raw byte) *DiagnosticMonitoringType {
	return &DiagnosticMonitoringType{
		TemperatureMonitoringImplemented:     raw&(1<<5) > 0,
		SupplyVoltageMonitoringImplemented:   raw&(1<<4) > 0,
		ReceivedPowerMeasurementsType:        ReceivedPowerMeasurementsType(raw&(1<<3) > 0),
		TransmitterPowerMeasurementSupported: raw&(1<<2) > 0,
	}
}
