package sff8079

import "github.com/wobcom/go-ethtool/eeprom"

// Options Transceiver options
type Options struct {
	PowerLevel3Requirement          bool
	PagingImplemented               bool
	RetimeOrCDRPresent              bool
	CooledImplementation            bool
	PowerLevel2Requirement          bool
	LinearReceiverOutputImplemented bool

	ReceiverDecisionThresholdImplemented bool
	TunableTransmitterTechnology         bool
	RateSelectImplemented                bool
	TxDisableImplemented                 bool
	TxFaultImplemented                   bool
	RxLOSInvImplemented                  bool
	RxLOSImplemented                     bool
}

var optionsMemoryMap = map[uint]map[uint]func(*Options, bool){
	0x00: map[uint]func(*Options, bool){
		// 0x07 unallocated
		// 0x06 unallocated
		0x05: func(o *Options, b bool) { o.PowerLevel3Requirement = b },
		0x04: func(o *Options, b bool) { o.PagingImplemented = b },
		0x03: func(o *Options, b bool) { o.RetimeOrCDRPresent = b },
		0x02: func(o *Options, b bool) { o.CooledImplementation = b },
		0x01: func(o *Options, b bool) { o.PowerLevel2Requirement = b },
		0x00: func(o *Options, b bool) { o.LinearReceiverOutputImplemented = b },
	},
	0x01: map[uint]func(*Options, bool){
		0x07: func(o *Options, b bool) { o.ReceiverDecisionThresholdImplemented = b },
		0x06: func(o *Options, b bool) { o.TunableTransmitterTechnology = b },
		0x05: func(o *Options, b bool) { o.RateSelectImplemented = b },
		0x04: func(o *Options, b bool) { o.TxDisableImplemented = b },
		0x03: func(o *Options, b bool) { o.TxFaultImplemented = b },
		0x02: func(o *Options, b bool) { o.RxLOSInvImplemented = b },
		0x01: func(o *Options, b bool) { o.RxLOSImplemented = b },
		// 0x00 unallocated
	},
}

// GetPowerClass returns the PowerClass from the given Options
func (o *Options) GetPowerClass() eeprom.PowerClass {
	if o.PowerLevel3Requirement {
		return eeprom.PowerClass3
	}
	if o.PowerLevel2Requirement {
		return eeprom.PowerClass2
	}
	return eeprom.PowerClass1
}

// NewOptions decodes a [2]byte into a new Options instance
func NewOptions(raw [2]byte) *Options {
	o := &Options{}
	for byteOffset, bitMap := range optionsMemoryMap {
		for bitOffset, callback := range bitMap {
			callback(o, raw[byteOffset]&(1<<bitOffset) > 0)
		}
	}
	return o
}
