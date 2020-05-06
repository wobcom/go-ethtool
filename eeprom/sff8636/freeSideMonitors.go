package SFF8636

type FreeSideMonitors struct {
	Temperature   float64 `json:"temperature"`
	SupplyVoltage float64 `json:"supplyVoltage"`
}

var freeSideMonitoringValuesMemoryMap = map[uint]func(*FreeSideMonitors, byte, byte){
	0x00: func(f *FreeSideMonitors, msb byte, lsb byte) {
		f.Temperature = parseTemperature(msb, lsb)
	},
	// 0x02-0x03 reserved
	0x04: func(f *FreeSideMonitors, msb byte, lsb byte) {
		f.SupplyVoltage = parseVoltage(msb, lsb)
	},
	// 0x06-0x07 reserved
	// 0x08-0x0B vendor specific
}

func NewFreeSideMonitors(raw [12]byte) *FreeSideMonitors {
	f := &FreeSideMonitors{}

	for byteIndex, callback := range freeSideMonitoringValuesMemoryMap {
		callback(f, raw[byteIndex], raw[byteIndex+1])
	}
	return f
}
