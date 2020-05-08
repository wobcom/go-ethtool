package sff8636

import (
	"encoding/json"
)

// StatusIndicators as defined in SFF-8636 rev 2.10a table 6-2
type StatusIndicators struct {
	RevisionCompliance RevisionCompliance
	StatusIndicator    *StatusIndicator
}

// RevisionCompliance compliance regarding the SFF-8636 standard, see SFF-8636 rev 2.10 table 6-3
type RevisionCompliance byte
// StatusIndicator as defined in SFF-8636 rev 2.10 table 6-2
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
    // RevisionUnspecified Revision not specified, might be anything < SFF-8636 Rev 2.5
	RevisionUnspecified              RevisionCompliance = 0x00
    // Revision4dot8orEarlier SFF-8436 Rev 4.8 or earlier
	Revision4dot8orEarlier      RevisionCompliance = 0x01
    // RevisionSFF8436andSFF8636 Includes functionality described in revision 4.8 or earlier of SFF-8436, except that this byte and Bytes 186-189 are as defined in this document
	RevisionSFF8436andSFF8636                    = 0x02
    // Revision1dot3orEarlier SFF-8636 Rev 1.3 or earlier
	Revision1dot3orEarlier      RevisionCompliance = 0x03
    // Revision1dot4 SFF-8636 Rev 1.4
	Revision1dot4              RevisionCompliance = 0x04
    // Revision1dot5 SFF-8636 Rev 1.5
	Revision1dot5              RevisionCompliance = 0x05
    // Revision2dot0 SFF-8636 Rev 2.0
	Revision2dot0              RevisionCompliance = 0x06
    // Revision2dot5and2dot6and2dot7 SFF-8636 Rev 2.5, 2.6 and 2.7
	Revision2dot5and2dot6and2dot7    RevisionCompliance = 0x07
    // Revision2dot8and2dot9and2dot10 SFF-8636 Rev 2.8, 2.9 and 2.10
	Revision2dot8and2dot9and2dot10   RevisionCompliance = 0x08
)

const (
	flatMemBitoffset       = 0x02
	intlBitOffset           = 0x01
	dataNotReadyBitoffset = 0x00
)

func (r RevisionCompliance) String() string {
	return map[RevisionCompliance]string{
		RevisionUnspecified:              "Revision not specified, might be anything < SFF-8636 Rev 2.5",
		Revision4dot8orEarlier:      "SFF-8436 Rev 4.8 or earlier",
		RevisionSFF8436andSFF8636: "Includes functionality described in revision 4.8 or earlier of SFF-8436, except that this byte and Bytes 186-189 are as defined in this document",
		Revision1dot3orEarlier:      "SFF-8636 Rev 1.3 or earlier",
		Revision1dot4:              "SFF-8636 Rev 1.4",
		Revision1dot5:              "SFF-8636 Rev 1.5",
		Revision2dot0:              "SFF-8636 Rev 2.0",
		Revision2dot5and2dot6and2dot7:    "SFF-8636 Rev 2.5, 2.6 and 2.7",
		Revision2dot8and2dot9and2dot10:   "SFF-8636 Rev 2.8, 2.9 and 2.10",
	}[r]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (s *StatusIndicators) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"revisionCompliance": s.RevisionCompliance,
		"statusIndicator":    s.StatusIndicator,
	})
}

// NewStatusIndiciator parses a byte into a new StatusIndicator instance
func NewStatusIndiciator(raw byte) *StatusIndicator {
	return &StatusIndicator{
		FlatMemory:   raw&(1<<flatMemBitoffset) > 0,
		IntL:         raw&(1<<intlBitOffset) > 0,
		DataNotReady: raw&(1<<dataNotReadyBitoffset) > 0,
	}
}

// NewStatusIndicators parses [2]byte into a new StatusIndicators instance
func NewStatusIndicators(raw [2]byte) *StatusIndicators {
	return &StatusIndicators{
		RevisionCompliance: RevisionCompliance(raw[0]),
		StatusIndicator:    NewStatusIndiciator(raw[1]),
	}
}
