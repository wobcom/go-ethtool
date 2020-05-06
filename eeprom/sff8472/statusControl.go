package SFF8472

type StatusControl struct {
	TxDisableState         bool
	SoftTxDisableSelect    bool
	InputPinRS1State       bool
	InputPinRS0State       bool
	FullbandwidthOperation bool
	TxFaultState           bool
	RxLosState             bool
	DataReadyBarState      bool
}

func NewStatusControl(raw byte) *StatusControl {
	return &StatusControl{
		TxDisableState:         raw&(1<<7) > 0,
		SoftTxDisableSelect:    raw&(1<<6) > 0,
		InputPinRS1State:       raw&(1<<5) > 0,
		InputPinRS0State:       raw&(1<<4) > 0,
		FullbandwidthOperation: raw&(1<<3) > 0,
		TxFaultState:           raw&(1<<2) > 0,
		RxLosState:             raw&(1<<1) > 0,
		DataReadyBarState:      raw&(1<<0) > 0,
	}
}
