package sff8079

import (
	"encoding/json"
	"fmt"
)

// RateIdentifier as of SFF-8079
type RateIdentifier byte

const (
	// RateIdentifierUnspecified unspecified
	RateIdentifierUnspecified RateIdentifier = 0x00
	// RateIdentifier421G SFF-8079 (4/2/1G Rate_Select & AS0/AS1)
	RateIdentifier421G RateIdentifier = 0x01
	// RateIdentifier842RxRateSelectOnly SFF-8431 (8/4/2G Rx Rate_Select only)
	RateIdentifier842RxRateSelectOnly RateIdentifier = 0x02
	// RateIdentifier842TxRateSelectOnly SFF-8431 (8/4/2G Tx Rate_Select only)
	RateIdentifier842TxRateSelectOnly = 0x04
	// RateIdentifier842IndependentRxTxRateSelect SFF-8431 (8/4/2G Independent Rx & Tx Rate_select)
	RateIdentifier842IndependentRxTxRateSelect RateIdentifier = 0x06
	// RateIdentifier1684RxRateSelectOnly FC-PI-5 (16/8/4G Rx Rate_select only) High=16G only, Low=8G/4G
	RateIdentifier1684RxRateSelectOnly RateIdentifier = 0x08
	// RateIdentifier1684IndependendRxTxRateSelect FC-PI-5 (16/8/4G Independent Rx, Tx Rate_select) High=16G only, Low=8G/4G
	RateIdentifier1684IndependendRxTxRateSelect RateIdentifier = 0x0A
	// RateIdentifier32168IndependendRxTxRateSelect FC-PI-6 (32/16/8G Independent Rx, Tx Rate_Select) High=32G only, Low = 16G/8G
	RateIdentifier32168IndependendRxTxRateSelect RateIdentifier = 0x0C
	// RateIdentifier108GRxTx 10/8G Rx and Tx Rate_Select controlling the operation or locking modes of the internal signal conditioner, retimer or CDR, according to the logic table defined in SFF-8472 rev 12.3 Table 10-2, High Bit Rate (10G) =9.95-11.3 Gb/s; Low Bit Rate (8G) = 8.5 Gb/s. In this mode, the default value of bit 110.3 (Soft Rate Select RS(0), Table 9-11) and of bit 118.3 (Soft Rate Select RS(1), Table 10-1) is 1.
	RateIdentifier108GRxTx RateIdentifier = 0x0E
	// RateIdentifier643216IndependendRxTxRateSelect FC-PI-7 (64/32/16G Independent Rx, Tx Rate Select) High = 32GFC and 64GFC. Low = 16GFC.
	RateIdentifier643216IndependendRxTxRateSelect RateIdentifier = 0x10
)

func (r RateIdentifier) String() string {
	return map[RateIdentifier]string{
		RateIdentifierUnspecified:                     "unspecified",
		RateIdentifier421G:                            "SFF-8079 (4/2/1G Rate_Select & AS0/AS1)",
		RateIdentifier842RxRateSelectOnly:             "SFF-8431 (8/4/2G Rx Rate_Select only)",
		RateIdentifier842TxRateSelectOnly:             "SFF-8431 (8/4/2G Tx Rate_Select only)",
		RateIdentifier842IndependentRxTxRateSelect:    "SFF-8431 (8/4/2G Independent Rx & Tx Rate_select)",
		RateIdentifier1684RxRateSelectOnly:            "FC-PI-5 (16/8/4G Rx Rate_select only) High=16G only, Low=8G/4G",
		RateIdentifier1684IndependendRxTxRateSelect:   "FC-PI-5 (16/8/4G Independent Rx, Tx Rate_select) High=16G only, Low=8G/4G",
		RateIdentifier32168IndependendRxTxRateSelect:  "FC-PI-6 (32/16/8G Independent Rx, Tx Rate_Select) High=32G only, Low = 16G/8G",
		RateIdentifier108GRxTx:                        " 10/8G Rx and Tx Rate_Select controlling the operation or locking modes of the internal signal conditioner, retimer or CDR, according to the logic table defined in SFF-8472 rev 12.3 Table 10-2, High Bit Rate (10G) =9.95-11.3 Gb/s; Low Bit Rate (8G) = 8.5 Gb/s. In this mode, the default value of bit 110.3 (Soft Rate Select RS(0), Table 9-11) and of bit 118.3 (Soft Rate Select RS(1), Table 10-1) is 1.",
		RateIdentifier643216IndependendRxTxRateSelect: "FC-PI-7 (64/32/16G Independent Rx, Tx Rate Select) High = 32GFC and 64GFC. Low = 16GFC.",
	}[r]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (r RateIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"ascii": r.String(),
		"hex":   fmt.Sprintf("%#02X", byte(r)),
	})
}
