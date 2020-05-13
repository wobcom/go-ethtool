package sff8079

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ActiveCableSpecifications maps an ActiveCableSpecification to a bool indicating its compliance (true) or not (false)
type ActiveCableSpecifications map[ActiveCableSpecification]bool

// ActiveCableSpecification an active cable specification
type ActiveCableSpecification int

const (
	// ActiveCableSpecificationFCPI4Limiting Compliant to FC-PI-4 Limiting
	ActiveCableSpecificationFCPI4Limiting ActiveCableSpecification = iota
	// ActiveCableSpecificationSFF8431Limiting Compliant to SFF-8431 Limiting
	ActiveCableSpecificationSFF8431Limiting
	// ActiveCableSpecificationFCPI4AppendixH Compliant to FC-PI-4 Appendix H
	ActiveCableSpecificationFCPI4AppendixH
	// ActiveCableSpecificationSFF8431AppendixE Compliant to SFF-8431 Appendix E
	ActiveCableSpecificationSFF8431AppendixE
)

func (a ActiveCableSpecification) String() string {
	return map[ActiveCableSpecification]string{
		ActiveCableSpecificationFCPI4Limiting:    "Compliant to FC-PI-4 Limiting",
		ActiveCableSpecificationSFF8431Limiting:  "Compliant to SFF-8431 Limiting",
		ActiveCableSpecificationFCPI4AppendixH:   "Compliant to FC-PI-4 Appendix H",
		ActiveCableSpecificationSFF8431AppendixE: "Compliant to SFF-8431 Appendix E",
	}[a]
}

func (a ActiveCableSpecifications) String() string {
	builder := &strings.Builder{}
	for activeCableSpecification, status := range a {
		if status {
			fmt.Fprintf(builder, " * %s\n", activeCableSpecification.String())
		}
	}
	return builder.String()
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (a ActiveCableSpecifications) MarshalJSON() ([]byte, error) {
	ret := []string{}
	for activeCableSpecification, status := range a {
		if status {
			ret = append(ret, activeCableSpecification.String())
		}
	}
	return json.Marshal(ret)
}

// NewActiveCableSpecifications decodes a [2]byte into a ActiveCableSpecifications
func NewActiveCableSpecifications(raw [2]byte) ActiveCableSpecifications {
	a := ActiveCableSpecifications{}
	a[ActiveCableSpecificationFCPI4Limiting] = raw[0]&(1<<3) > 0
	a[ActiveCableSpecificationSFF8431Limiting] = raw[0]&(1<<2) > 0
	a[ActiveCableSpecificationFCPI4AppendixH] = raw[0]&(1<<1) > 0
	a[ActiveCableSpecificationSFF8431AppendixE] = raw[0]&(1<<0) > 0
	return a
}
