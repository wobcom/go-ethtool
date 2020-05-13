package sff8636

// Control functions
type Control struct {
	ChannelControls      [4]ChannelControl
	PowerClass8Enable    bool
	PowerClass5To7Enable bool
	LowPowerMode         bool
	PowerOverride        bool
	LPModeTxDIS          bool
	IntlLOSL             bool
}

// ChannelControl channel specific controls
type ChannelControl struct {
	TxDisable    bool
	RxRateSelect RateSelect
	TxRateSelect RateSelect
	TxCDR        bool
	RxCDR        bool
}

// RateSelect used for software rate select
type RateSelect struct {
	MSB bool
	LSB bool
}

var controlMemoryMap = map[uint]map[uint]func(*Control, bool){
	0x00: map[uint]func(*Control, bool){
		// 0x07-0x04 reserved
		0x03: func(c *Control, v bool) { c.ChannelControls[3].TxDisable = v },
		0x02: func(c *Control, v bool) { c.ChannelControls[2].TxDisable = v },
		0x01: func(c *Control, v bool) { c.ChannelControls[1].TxDisable = v },
		0x00: func(c *Control, v bool) { c.ChannelControls[0].TxDisable = v },
	},
	0x01: map[uint]func(*Control, bool){
		0x07: func(c *Control, v bool) { c.ChannelControls[3].RxRateSelect.MSB = v },
		0x06: func(c *Control, v bool) { c.ChannelControls[3].RxRateSelect.LSB = v },
		0x05: func(c *Control, v bool) { c.ChannelControls[2].RxRateSelect.MSB = v },
		0x04: func(c *Control, v bool) { c.ChannelControls[2].RxRateSelect.LSB = v },
		0x03: func(c *Control, v bool) { c.ChannelControls[1].RxRateSelect.MSB = v },
		0x02: func(c *Control, v bool) { c.ChannelControls[1].RxRateSelect.LSB = v },
		0x01: func(c *Control, v bool) { c.ChannelControls[0].RxRateSelect.MSB = v },
		0x00: func(c *Control, v bool) { c.ChannelControls[0].RxRateSelect.LSB = v },
	},
	0x02: map[uint]func(*Control, bool){
		0x07: func(c *Control, v bool) { c.ChannelControls[3].TxRateSelect.MSB = v },
		0x06: func(c *Control, v bool) { c.ChannelControls[3].TxRateSelect.LSB = v },
		0x05: func(c *Control, v bool) { c.ChannelControls[2].TxRateSelect.MSB = v },
		0x04: func(c *Control, v bool) { c.ChannelControls[2].TxRateSelect.LSB = v },
		0x03: func(c *Control, v bool) { c.ChannelControls[1].TxRateSelect.MSB = v },
		0x02: func(c *Control, v bool) { c.ChannelControls[1].TxRateSelect.LSB = v },
		0x01: func(c *Control, v bool) { c.ChannelControls[0].TxRateSelect.MSB = v },
		0x00: func(c *Control, v bool) { c.ChannelControls[0].TxRateSelect.LSB = v },
	},
	// 0x03-0x06 reserved (prior to rev 2.10 used for SFF-8079 - now deprecated)
	0x07: map[uint]func(*Control, bool){
		// 0x07 SW reset - read only
		// 0x06-0x04 reserved
		0x03: func(c *Control, v bool) { c.PowerClass8Enable = v },
		0x02: func(c *Control, v bool) { c.PowerClass5To7Enable = v },
		0x01: func(c *Control, v bool) { c.LowPowerMode = v },
		0x00: func(c *Control, v bool) { c.PowerOverride = v },
	},
	// 0x08-0xB reserved
	0x0C: map[uint]func(*Control, bool){
		0x07: func(c *Control, v bool) { c.ChannelControls[3].TxCDR = v },
		0x06: func(c *Control, v bool) { c.ChannelControls[2].TxCDR = v },
		0x05: func(c *Control, v bool) { c.ChannelControls[1].TxCDR = v },
		0x04: func(c *Control, v bool) { c.ChannelControls[0].TxCDR = v },
		0x03: func(c *Control, v bool) { c.ChannelControls[3].RxCDR = v },
		0x02: func(c *Control, v bool) { c.ChannelControls[2].RxCDR = v },
		0x01: func(c *Control, v bool) { c.ChannelControls[1].RxCDR = v },
		0x00: func(c *Control, v bool) { c.ChannelControls[0].RxCDR = v },
	},
	0x0D: map[uint]func(*Control, bool){
		// 0x07-0x02 reserved
		0x01: func(c *Control, v bool) { c.LPModeTxDIS = v },
		0x00: func(c *Control, v bool) { c.IntlLOSL = v },
	},
}

// NewControl parses [14]byte into a new Control instance
func NewControl(raw [14]byte) *Control {
	c := &Control{}

	for byteOffset, bitMap := range controlMemoryMap {
		for bitOffset, callback := range bitMap {
			value := raw[byteOffset]&(1<<bitOffset) > 0
			callback(c, value)
		}
	}
	return c
}
