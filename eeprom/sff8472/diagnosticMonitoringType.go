package SFF8472

import (
	"encoding/json"
)

type DiagnosticMonitoringType struct {
	DiagnosticMonitoringImplemented bool
	InternallyCalibrated            bool
	ExternallyCalibrated            bool
	ReceivedPowerMeasurementType    ReceivedPowerMeasurementType
}

type ReceivedPowerMeasurementType bool

const (
	ReceivedPowerMeasurementTypeOMA          ReceivedPowerMeasurementType = false
	ReceivedPowerMeasurementTypeAveragePower ReceivedPowerMeasurementType = true
)

func (r ReceivedPowerMeasurementType) String() string {
	if r == ReceivedPowerMeasurementTypeOMA {
		return "OMA"
	}
	return "Average power"
}

func (r ReceivedPowerMeasurementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func NewDiagnosticMonitoringType(raw byte) *DiagnosticMonitoringType {
	return &DiagnosticMonitoringType{
		DiagnosticMonitoringImplemented: raw&(1<<6) > 0,
		InternallyCalibrated:            raw&(1<<5) > 0,
		ExternallyCalibrated:            raw&(1<<4) > 0,
		ReceivedPowerMeasurementType:    ReceivedPowerMeasurementType(raw&(1<<3) > 0),
	}
}
