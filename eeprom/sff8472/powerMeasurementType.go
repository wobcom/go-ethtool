package sff8472

// PowerMeasurementType how to interpret power measurements reported by the transceiver
type PowerMeasurementType bool

const (
	// OmaPowerMeasurementType OMA
	OmaPowerMeasurementType PowerMeasurementType = false
	// AveragePowerMeasurementType average power
	AveragePowerMeasurementType PowerMeasurementType = true
)

func (p PowerMeasurementType) String() string {
	if p == OmaPowerMeasurementType {
		return "OMA power measurement"
	}
	return "Average power measurement"
}
