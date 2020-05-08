package sff8636

// EnhancedOptions as of SFF-8636 rev 2.10a table 6-24
type EnhancedOptions struct {
	InitializationCompleteFlag bool
	RateSelectImplemented      bool
	TCReadinessImplemented     bool
	SoftwareResetImplemented   bool
}

// NewEnhancedOptions parses a byte into a new EnhancedOptions instance
func NewEnhancedOptions(raw byte) *EnhancedOptions {
	return &EnhancedOptions{
		InitializationCompleteFlag: raw&(1<<4) > 0,
		RateSelectImplemented:      raw&(1<<3) > 0,
		TCReadinessImplemented:     raw&(1<<1) > 0,
		SoftwareResetImplemented:   raw&(1<<0) > 0,
	}
}
