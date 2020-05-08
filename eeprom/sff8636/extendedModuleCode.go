package sff8636

// ExtendedModuleCodeValues maps a ExtendedModuleCode to a bool to indicate support of a electronic or optical interface for InfiniBand
type ExtendedModuleCodeValues map[ExtendedModuleCode]bool
// ExtendedModuleCode electronic or optical interface for InfiniBand
type ExtendedModuleCode int

const (
    // ExtendedModuleCodeHDR HDR
	ExtendedModuleCodeHDR ExtendedModuleCode = iota
    // ExtendedModuleCodeEDR EDR
	ExtendedModuleCodeEDR
    // ExtendedModuleCodeFDR FDR
	ExtendedModuleCodeFDR
    // ExtendedModuleCodeQDR QDR
	ExtendedModuleCodeQDR
    // ExtendedModuleCodeDDR DDR
	ExtendedModuleCodeDDR
    // ExtendedModuleCodeSDR SDR
	ExtendedModuleCodeSDR
)

var extendedModuleCodeMemoryMap = map[uint]ExtendedModuleCode{
	0x05: ExtendedModuleCodeHDR,
	0x04: ExtendedModuleCodeEDR,
	0x03: ExtendedModuleCodeFDR,
	0x02: ExtendedModuleCodeQDR,
	0x01: ExtendedModuleCodeDDR,
	0x00: ExtendedModuleCodeSDR,
}

// NewExtendedModuleCodeValues parses a byte into a new ExtendedModuleCodeValues instance
func NewExtendedModuleCodeValues(raw byte) *ExtendedModuleCodeValues {
	e := ExtendedModuleCodeValues{}

	for bitIndex, extendedModuleCode := range extendedModuleCodeMemoryMap {
		e[extendedModuleCode] = raw&(1<<bitIndex) > 0
	}
	return &e
}
