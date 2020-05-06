package SFF8472

type PowerMeasurementType bool

const (
	OmaPowerMeasurementType     PowerMeasurementType = false
	AveragePowerMeasurementType PowerMeasurementType = true
)

func (p PowerMeasurementType) String() string {
	if p == OmaPowerMeasurementType {
		return "OMA power measurement"
	}
	return "Average power measurement"
}
