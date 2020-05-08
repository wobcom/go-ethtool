package sff8472

import "encoding/json"

// ExtendedStatusControl Extended module control and status bytes as of SFF-847 Rev 12.3 Table 10-1
type ExtendedStatusControl struct {
	SoftRS1Select            bool
	PowerLevelOperationState PowerLevelOperationState
	Gfc64ModeTxConfigured    bool
	Gfc64ModeRxConfigured    bool
	Gfc64Mode                bool
	TxCdrUnlocked            bool
	RxCdrUnlocked            bool
}

// PowerLevelOperationState power level the transceiver is operating at
type PowerLevelOperationState bool

const (
    // PowerLevelOperationStatePowerLevel1 Power Level 1 operation (1.0 Watt max)
	PowerLevelOperationStatePowerLevel1   PowerLevelOperationState = false
    // PowerLevelOperationStatePowerLevel2_3 Power Level 2 or 3 operation (1.5 or 2.0 Watt max)
	PowerLevelOperationStatePowerLevel2_3 PowerLevelOperationState = true
)

func (p PowerLevelOperationState) String() string {
	if p == PowerLevelOperationStatePowerLevel1 {
		return "Power Level 1 operation (1.0 Watt max)"
	}
	return "Power Level 2 or 3 operation (1.5 or 2.0 Watt max)"
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (p PowerLevelOperationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// NewExtendedStatusControl parses [2]byte into a new ExtendedStatusControl instance
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
