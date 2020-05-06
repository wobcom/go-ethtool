package SFF8472

import "encoding/json"

type ExtendedStatusControl struct {
	SoftRS1Select            bool
	PowerLevelOperationState PowerLevelOperationState
	Gfc64ModeTxConfigured    bool
	Gfc64ModeRxConfigured    bool
	Gfc64Mode                bool
	TxCdrUnlocked            bool
	RxCdrUnlocked            bool
}

type PowerLevelOperationState bool

const (
	PowerLevelOperationStatePowerLevel1   PowerLevelOperationState = false
	PowerLevelOperationStatePowerLevel2_3 PowerLevelOperationState = true
)

func (p PowerLevelOperationState) String() string {
	if p == PowerLevelOperationStatePowerLevel1 {
		return "Power Level 1 operation (1.0 Watt max)"
	}
	return "Power Level 2 or 3 operation (1.5 or 2.0 Watt max)"
}

func (p PowerLevelOperationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func NewExtendedStatusControl(raw [2]byte) *ExtendedStatusControl {
	return &ExtendedStatusControl{
		SoftRS1Select:            raw[0]&(1<<3) > 0,
		PowerLevelOperationState: PowerLevelOperationState(raw[0]&(1<<1) > 0),
		Gfc64ModeTxConfigured:    raw[1]&(1<<4) > 0,
		Gfc64ModeRxConfigured:    raw[1]&(1<<3) > 0,
		Gfc64Mode:                raw[1]&(1<<2) > 0,
		TxCdrUnlocked:            raw[1]&(1<<1) > 0,
		RxCdrUnlocked:            raw[1]&(1<<0) > 0,
	}
}
