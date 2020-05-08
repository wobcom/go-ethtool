package sff8636

// ChannelMonitors Real-time channel monitoring for each transmit and receive channel includes optical input power and Tx bias current.
type ChannelMonitors [4]ChannelMonitor

// ChannelMonitor Real-time channel monitoring data
type ChannelMonitor struct {
	RxPower Power
	TxPower Power
	Bias    float64
}

var channelMonitorsMemoryMap = map[uint]func(*ChannelMonitors, byte, byte){
	0x00: func(c *ChannelMonitors, msb byte, lsb byte) { c[0].RxPower = parsePower(msb, lsb) },
	0x02: func(c *ChannelMonitors, msb byte, lsb byte) { c[1].RxPower = parsePower(msb, lsb) },
	0x04: func(c *ChannelMonitors, msb byte, lsb byte) { c[2].RxPower = parsePower(msb, lsb) },
	0x06: func(c *ChannelMonitors, msb byte, lsb byte) { c[3].RxPower = parsePower(msb, lsb) },
	0x08: func(c *ChannelMonitors, msb byte, lsb byte) { c[0].Bias = parseCurrent(msb, lsb) },
	0x0A: func(c *ChannelMonitors, msb byte, lsb byte) { c[1].Bias = parseCurrent(msb, lsb) },
	0x0C: func(c *ChannelMonitors, msb byte, lsb byte) { c[2].Bias = parseCurrent(msb, lsb) },
	0x0E: func(c *ChannelMonitors, msb byte, lsb byte) { c[3].Bias = parseCurrent(msb, lsb) },
	0x10: func(c *ChannelMonitors, msb byte, lsb byte) { c[0].TxPower = parsePower(msb, lsb) },
	0x12: func(c *ChannelMonitors, msb byte, lsb byte) { c[1].TxPower = parsePower(msb, lsb) },
	0x14: func(c *ChannelMonitors, msb byte, lsb byte) { c[2].TxPower = parsePower(msb, lsb) },
	0x16: func(c *ChannelMonitors, msb byte, lsb byte) { c[3].TxPower = parsePower(msb, lsb) },
	// 0x18-0x1F reserved channel monitor set 4
	// 0x20-0x27 reserved channel monitor set 5
	// 0x28-0x2F vendor specific
}

// NewChannelMonitors parses [48]byte into a new ChannelMonitors instance
func NewChannelMonitors(raw [48]byte) *ChannelMonitors {
	c := &ChannelMonitors{}

	for byteIndex, callback := range channelMonitorsMemoryMap {
		callback(c, raw[byteIndex], raw[byteIndex+1])
	}
	return c
}
