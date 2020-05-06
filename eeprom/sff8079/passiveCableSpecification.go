package SFF8079

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PassiveCableSpecifications map[PassiveCableSpecification]bool
type PassiveCableSpecification int

const (
	PassiveCableSpecificationFC_PI_4_AppendixH PassiveCableSpecification = iota
	PassiveCableSpecificationSFF8431_AppendixE
)

func (p PassiveCableSpecification) String() string {
	return map[PassiveCableSpecification]string{
		PassiveCableSpecificationFC_PI_4_AppendixH: "Compliant to FC-PI-4 Appendix H",
		PassiveCableSpecificationSFF8431_AppendixE: "Compliant to SFF-8431 Appendix E",
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

func (p PassiveCableSpecifications) MarshalJSON() ([]byte, error) {
	ret := []string{}
	for passiveCableSpecification, status := range p {
		if status {
			ret = append(ret, passiveCableSpecification.String())
		}
	}
	return json.Marshal(ret)
}

func NewPassiveCableSpecifications(raw [2]byte) PassiveCableSpecifications {
	p := PassiveCableSpecifications{}
	p[PassiveCableSpecificationFC_PI_4_AppendixH] = raw[0]&(1<<1) > 0
	p[PassiveCableSpecificationSFF8431_AppendixE] = raw[0]&(1<<0) > 0
	return p
}
