package SFF8472

type EnhancedOptions struct {
	AlarmWarningFlagsImplemented                    bool
	SoftTxDisableControlAndMonitoringImplemented    bool
	SoftTxFaultImplemented                          bool
	SoftRxLosImplemented                            bool
	SoftRateSelectControlAndMonitoringImplemented   bool
	ApplicationSelectControlImplementedAsPerSFF8079 bool
	SoftRateSelectImplementedAsPerSFF8431           bool
}

func NewEnhancedOptions(raw byte) *EnhancedOptions {
	return &EnhancedOptions{
		AlarmWarningFlagsImplemented:                    raw&(1<<7) > 0,
		SoftTxDisableControlAndMonitoringImplemented:    raw&(1<<6) > 0,
		SoftTxFaultImplemented:                          raw&(1<<5) > 0,
		SoftRxLosImplemented:                            raw&(1<<4) > 0,
		SoftRateSelectControlAndMonitoringImplemented:   raw&(1<<3) > 0,
		ApplicationSelectControlImplementedAsPerSFF8079: raw&(1<<2) > 0,
		SoftRateSelectImplementedAsPerSFF8431:           raw&(1<<1) > 0,
	}
}
