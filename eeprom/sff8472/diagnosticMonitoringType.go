package sff8472

import (
	"encoding/json"
)

// DiagnosticMonitoringType as of SFF-8472
type DiagnosticMonitoringType struct {
	DiagnosticMonitoringImplemented bool
	InternallyCalibrated            bool
	ExternallyCalibrated            bool
	ReceivedPowerMeasurementType    ReceivedPowerMeasurementType
}

// ReceivedPowerMeasurementType how to interpret the rx power readings
type ReceivedPowerMeasurementType bool

const (
    // ReceivedPowerMeasurementTypeOMA OMA
	ReceivedPowerMeasurementTypeOMA          ReceivedPowerMeasurementType = false
    // ReceivedPowerMeasurementTypeAveragePower Average power
	ReceivedPowerMeasurementTypeAveragePower ReceivedPowerMeasurementType = true
)

func (r ReceivedPowerMeasurementType) String() string {
	if r == ReceivedPowerMeasurementTypeOMA {
		return "OMA"
	}
	return "Average power"
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (r ReceivedPowerMeasurementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// NewDiagnosticMonitoringType parses a byte into a new instance of DiagnosticMonitoringType
func NewDiagnosticMonitoringType(raw byte) *DiagnosticMonitoringType {
	return &DiagnosticMonitoringType{
		DiagnosticMonitoringImplemented: raw&(1<<6) > 0,
		InternallyCalibrated:            raw&(1<<5) > 0,
		ExternallyCalibrated:            raw&(1<<4) > 0,
		ReceivedPowerMeasurementType:    ReceivedPowerMeasurementType(raw&(1<<3) > 0),
	}
}
