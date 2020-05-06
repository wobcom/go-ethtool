package SFF8636

type ExtendedModuleCodeValues map[ExtendedModuleCode]bool

type ExtendedModuleCode int

const (
	ExtendedModuleCodeHDR ExtendedModuleCode = iota
	ExtendedModuleCodeEDR
	ExtendedModuleCodeFDR
	ExtendedModuleCodeQDR
	ExtendedModuleCodeDDR
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

func NewExtendedModuleCodeValues(raw byte) *ExtendedModuleCodeValues {
	e := ExtendedModuleCodeValues{}

	for bitIndex, extendedModuleCode := range extendedModuleCodeMemoryMap {
		e[extendedModuleCode] = raw&(1<<bitIndex) > 0
	}
	return &e
}
