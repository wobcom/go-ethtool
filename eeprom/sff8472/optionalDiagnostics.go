package SFF8472

type OptionalDiagnostics struct {
	LaserTemperature float64
	TecCurrent       float64
}

func NewOptionalDiagnostics(raw [4]byte) *OptionalDiagnostics {
	return &OptionalDiagnostics{
		LaserTemperature: parseTemperature(raw[0], raw[1]), // TODO this may also mean wavelength
		TecCurrent:       parseCurrent(raw[2], raw[3]),
	}
}
