package sff8079

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PassiveCableSpecifications maps a PassiveCableSpecification to a bool indicating its compliance (true) or not (false)
type PassiveCableSpecifications map[PassiveCableSpecification]bool
// PassiveCableSpecification a passive cable specification
type PassiveCableSpecification int

const (
    // PassiveCableSpecificationFCPI4AppendixH Compliant to FC-PI-4 Appendix H
	PassiveCableSpecificationFCPI4AppendixH PassiveCableSpecification = iota
    // PassiveCableSpecificationSFF8431AppendixE Compliant to SFF-8431 Appendix E
	PassiveCableSpecificationSFF8431AppendixE
)

func (p PassiveCableSpecification) String() string {
	return map[PassiveCableSpecification]string{
		PassiveCableSpecificationFCPI4AppendixH: "Compliant to FC-PI-4 Appendix H",
		PassiveCableSpecificationSFF8431AppendixE: "Compliant to SFF-8431 Appendix E",
	}[p]
}

func (p PassiveCableSpecifications) String() string {
	builder := &strings.Builder{}
	for passiveCableSpecification, status := range p {
		if status {
			fmt.Fprintf(builder, " * %s\n", passiveCableSpecification.String())
		}
	}
	return builder.String()
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (p PassiveCableSpecifications) MarshalJSON() ([]byte, error) {
	ret := []string{}
	for passiveCableSpecification, status := range p {
		if status {
			ret = append(ret, passiveCableSpecification.String())
		}
	}
	return json.Marshal(ret)
}

// NewPassiveCableSpecifications decode a [2]byte into a PassiveCableSpecifications
func NewPassiveCableSpecifications(raw [2]byte) PassiveCableSpecifications {
	p := PassiveCableSpecifications{}
	p[PassiveCableSpecificationFCPI4AppendixH] = raw[0]&(1<<1) > 0
	p[PassiveCableSpecificationSFF8431AppendixE] = raw[0]&(1<<0) > 0
	return p
}
