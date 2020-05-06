package SFF8636

import (
	"encoding/json"
	"fmt"
)

type StatusIndicators struct {
	RevisionCompliance RevisionCompliance
	StatusIndicator    *StatusIndicator
}

type RevisionCompliance byte
type StatusIndicator struct {
	/* Upper memory flat or paged.
	   Bit 2 = 1b: Flat memory (lower and upper pages 00h only),
	   Bit 2 = 0b: Paging (at least upper page 03h implemented),
	    See Page 00h, Byte 195 for additional advertising. */
	FlatMemory bool `json:"flatMemory"`
	/* Digital state of the IntL Interrupt output pin. 1 =
	   IntL not asserted, 0 = IntL asserted. Default = 1. */
	IntL bool `json:"intL"`
	/* Indicates free-side does not yet have valid monitor
	   data. The bit remains high until valid data can be
	   read at which time the bit goes low. */
	DataNotReady bool `json:"dataNotReady"`
}

const (
	REVISION_UNSPECIFIED              RevisionCompliance = 0x00
	REVISION_SFF8436_4_8_EARLIER      RevisionCompliance = 0x01
	REVISION_SFF8436_AND_SOME_SFF8636                    = 0x02
	REVISION_SFF8636_1_3_EARLIER      RevisionCompliance = 0x03
	REVISION_SFF8636_1_4              RevisionCompliance = 0x04
	REVISION_SFF8636_1_5              RevisionCompliance = 0x05
	REVISION_SFF8636_2_0              RevisionCompliance = 0x06
	REVISION_SFF8636_2_5__2_6__2_7    RevisionCompliance = 0x07
	REVISION_SFF8636_2_8__2_9__2_10   RevisionCompliance = 0x08
)

const (
	STATUS_INDICATOR2_FLAT_MEM_BIT       = 0x02
	STATUS_INDICATOR2_INTL_BIT           = 0x01
	STATUS_INDICATOR2_DATA_NOT_READY_BIT = 0x00
)

func (r RevisionCompliance) String() string {
	return map[RevisionCompliance]string{
		REVISION_UNSPECIFIED:              "Revision not specified, might be anything < SFF-8636 Rev 2.5",
		REVISION_SFF8436_4_8_EARLIER:      "SFF-8436 Rev 4.8 or earlier",
		REVISION_SFF8436_AND_SOME_SFF8636: "Includes functionality described in revision 4.8 or earlier of SFF-8436, except that this byte and Bytes 186-189 are as defined in this document",
		REVISION_SFF8636_1_3_EARLIER:      "SFF-8636 Rev 1.3 or earlier",
		REVISION_SFF8636_1_4:              "SFF-8636 Rev 1.4",
		REVISION_SFF8636_1_5:              "SFF-8636 Rev 1.5",
		REVISION_SFF8636_2_0:              "SFF-8636 Rev 2.0",
		REVISION_SFF8636_2_5__2_6__2_7:    "SFF-8636 Rev 2.5, 2.6 and 2.7",
		REVISION_SFF8636_2_8__2_9__2_10:   "SFF-8636 Rev 2.8, 2.9 and 2.10",
	}[r]
}

func (r RevisionCompliance) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"hex":   fmt.Sprintf("%#02X", byte(r)),
		"ascii": r.String(),
	}
}

func (s *StatusIndicators) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"revisionCompliance": s.RevisionCompliance.MarshalJson(),
		"statusIndicator":    s.StatusIndicator,
	})
}

func NewStatusIndiciator(raw byte) *StatusIndicator {
	return &StatusIndicator{
		FlatMemory:   raw&(1<<STATUS_INDICATOR2_FLAT_MEM_BIT) > 0,
		IntL:         raw&(1<<STATUS_INDICATOR2_INTL_BIT) > 0,
		DataNotReady: raw&(1<<STATUS_INDICATOR2_DATA_NOT_READY_BIT) > 0,
	}
}

func NewStatusIndicators(raw [2]byte) *StatusIndicators {
	return &StatusIndicators{
		RevisionCompliance: RevisionCompliance(raw[0]),
		StatusIndicator:    NewStatusIndiciator(raw[1]),
	}
}
