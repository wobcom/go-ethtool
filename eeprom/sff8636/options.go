package SFF8636

type Options struct {
	LPModeTxDisConfigurable                bool
	IntLRxLOSLConfigurable                 bool
	TxInputAdaptiveEqualizersFreezeCapable bool
	TxInputEqualizersAutoAdapativeCapable  bool
	TxInputEqualizersFixedProgrammable     bool
	RxOutputEmphasisFixedProgrammable      bool
	RxOutputAmplitudeFixedProgrammable     bool

	TxCDROnOffControlImplemented bool
	RxCDROnOffControlImplemented bool
	TxCDRLOLImplemented          bool
	RxCDRLOLImplemented          bool
	RxSquelchDisableImplemented  bool
	RxOutputDisableImplemented   bool
	TxSquelchDisableImplemented  bool
	TxSquelchImplemented         bool

	MemoryPage02hProvided          bool
	MemoryPage01hProvided          bool
	RateSelectMultiRateImplemented bool
	TxDisableImplemented           bool
	TxFaultImplemented             bool
	TxSquelchImplmenetedReducePave bool
	TxLOSImplemented               bool
	MemoryPages20h21hProvided      bool
}

var optionsMemoryMap = map[uint]map[uint]func(*Options, bool){
	0x00: map[uint]func(*Options, bool){
		0x06: func(o *Options, b bool) { o.LPModeTxDisConfigurable = b },
		0x05: func(o *Options, b bool) { o.IntLRxLOSLConfigurable = b },
		0x04: func(o *Options, b bool) { o.TxInputAdaptiveEqualizersFreezeCapable = b },
		0x03: func(o *Options, b bool) { o.TxInputEqualizersAutoAdapativeCapable = b },
		0x02: func(o *Options, b bool) { o.TxInputEqualizersFixedProgrammable = b },
		0x01: func(o *Options, b bool) { o.RxOutputEmphasisFixedProgrammable = b },
		0x00: func(o *Options, b bool) { o.RxOutputAmplitudeFixedProgrammable = b },
	},
	0x01: map[uint]func(*Options, bool){
		0x07: func(o *Options, b bool) { o.TxCDROnOffControlImplemented = b },
		0x06: func(o *Options, b bool) { o.RxCDROnOffControlImplemented = b },
		0x05: func(o *Options, b bool) { o.TxCDRLOLImplemented = b },
		0x04: func(o *Options, b bool) { o.RxCDRLOLImplemented = b },
		0x03: func(o *Options, b bool) { o.RxSquelchDisableImplemented = b },
		0x02: func(o *Options, b bool) { o.RxOutputDisableImplemented = b },
		0x01: func(o *Options, b bool) { o.TxSquelchDisableImplemented = b },
		0x00: func(o *Options, b bool) { o.TxSquelchImplemented = b },
	},
	0x02: map[uint]func(*Options, bool){
		0x07: func(o *Options, b bool) { o.MemoryPage02hProvided = b },
		0x06: func(o *Options, b bool) { o.MemoryPage01hProvided = b },
		0x05: func(o *Options, b bool) { o.RateSelectMultiRateImplemented = b },
		0x04: func(o *Options, b bool) { o.TxDisableImplemented = b },
		0x03: func(o *Options, b bool) { o.TxFaultImplemented = b },
		0x02: func(o *Options, b bool) { o.TxSquelchImplmenetedReducePave = b },
		0x01: func(o *Options, b bool) { o.TxLOSImplemented = b },
		0x00: func(o *Options, b bool) { o.MemoryPages20h21hProvided = b },
	},
}

func NewOptions(raw [3]byte) *Options {
	o := &Options{}
	for byteOffset, bitMap := range optionsMemoryMap {
		for bitOffset, callback := range bitMap {
			value := raw[byteOffset]&(1<<bitOffset) > 0
			callback(o, value)
		}
	}
	return o
}
