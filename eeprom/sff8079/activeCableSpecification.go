package SFF8079

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ActiveCableSpecifications map[ActiveCableSpecification]bool
type ActiveCableSpecification int

const (
	ActiveCableSpecificationFC_PI_4_Limiting ActiveCableSpecification = iota
	ActiveCableSpecificationSFF8431_Limiting
	ActiveCableSpecificationFC_PI_4_AppendixH
	ActiveCableSpecificationSFF8431_AppendixE
)

func (a ActiveCableSpecification) String() string {
	return map[ActiveCableSpecification]string{
		ActiveCableSpecificationFC_PI_4_Limiting:  "Compliant to FC-PI-4 Limiting",
		ActiveCableSpecificationSFF8431_Limiting:  "Compliant to SFF-8431 Limiting",
		ActiveCableSpecificationFC_PI_4_AppendixH: "Compliant to FC-PI-4 Appendix H",
		ActiveCableSpecificationSFF8431_AppendixE: "Compliant to SFF-8431 Appendix E",
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

func (a ActiveCableSpecifications) MarshalJSON() ([]byte, error) {
	ret := []string{}
	for activeCableSpecification, status := range a {
		if status {
			ret = append(ret, activeCableSpecification.String())
		}
	}
	return json.Marshal(ret)
}

func NewActiveCableSpecifications(raw [2]byte) ActiveCableSpecifications {
	a := ActiveCableSpecifications{}
	a[ActiveCableSpecificationFC_PI_4_Limiting] = raw[0]&(1<<3) > 0
	a[ActiveCableSpecificationSFF8431_Limiting] = raw[0]&(1<<2) > 0
	a[ActiveCableSpecificationFC_PI_4_AppendixH] = raw[0]&(1<<1) > 0
	a[ActiveCableSpecificationSFF8431_AppendixE] = raw[0]&(1<<0) > 0
	return a
}
