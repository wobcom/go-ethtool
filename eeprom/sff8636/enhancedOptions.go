package SFF8636

type EnhancedOptions struct {
	InitializationCompleteFlag bool
	RateSelectImplemented      bool
	TCReadinessImplemented     bool
	SoftwareResetImplemented   bool
}

func NewEnhancedOptions(raw byte) *EnhancedOptions {
	return &EnhancedOptions{
		InitializationCompleteFlag: raw&(1<<4) > 0,
		RateSelectImplemented:      raw&(1<<3) > 0,
		TCReadinessImplemented:     raw&(1<<1) > 0,
		SoftwareResetImplemented:   raw&(1<<0) > 0,
	}
}
