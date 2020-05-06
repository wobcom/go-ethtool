package SFF8636

type DiagnosticMonitoringType struct {
	TemperatureMonitoringImplemented     bool
	SupplyVoltageMonitoringImplemented   bool
	ReceivedPowerMeasurementsType        ReceivedPowerMeasurementsType
	TransmitterPowerMeasurementSupported bool
}

type ReceivedPowerMeasurementsType bool

func (r ReceivedPowerMeasurementsType) String() string {
	if bool(r) {
		return "OMA"
	}
	return "Average Power"
}

func NewDiagnosticMonitoringType(raw byte) *DiagnosticMonitoringType {
	return &DiagnosticMonitoringType{
		TemperatureMonitoringImplemented:     raw&(1<<5) > 0,
		SupplyVoltageMonitoringImplemented:   raw&(1<<4) > 0,
		ReceivedPowerMeasurementsType:        ReceivedPowerMeasurementsType(raw&(1<<3) > 0),
		TransmitterPowerMeasurementSupported: raw&(1<<2) > 0,
	}
}
