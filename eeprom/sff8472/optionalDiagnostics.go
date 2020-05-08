package sff8472

// OptionalDiagnostics Monitor Data for Optional Laser temperature and TEC current as of SFF-8472 rev 12.3 table 9-11
type OptionalDiagnostics struct {
	LaserTemperature float64
	TecCurrent       float64
}

// NewOptionalDiagnostics parses [4]byte into a new OptionalDiagnostics instance
func NewOptionalDiagnostics(raw [4]byte) *OptionalDiagnostics {
	return &OptionalDiagnostics{
		LaserTemperature: parseTemperature(raw[0], raw[1]), // TODO this may also mean wavelength
		TecCurrent:       parseCurrent(raw[2], raw[3]),
	}
}
