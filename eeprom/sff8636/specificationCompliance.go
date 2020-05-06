package SFF8636

import (
	"encoding/json"
)

type SpecificationCompliance map[Specification]bool

type Specification int

const (
	/* 10G / 40G / 100G Ethernet Compliance Codes */
	SpecExteded Specification = iota
	Spec10GBaseLRM
	Spec10GBaseLR
	Spec10GBaseSR
	Spec40GBaseCR4
	Spec40GBaseSR4
	Spec40GBaseLR4
	Spec40GXLPPI

	/* SONET Compliance Codes */
	SpecOC48LongReach
	SpecOC48IntermediateReach
	SpecOC48ShortReach

	/* SAS/SATA Compliance Codes */
	SpecSAS24G
	SpecSAS12G
	SpecSAS6G
	SpecSAS3G

	/* Gigabit Ethernet Compliance Codes */
	Spec1000BaseT
	Spec1000BaseCX
	Spec1000BaseLX
	Spec1000BaseSX

	/* Fibre Channel Link Length */
	SpecVeryLongDistance
	SpecShortDistance
	SpecIntermediateDistance
	SpecLongDistance
	SpecMediumDistance

	/* Fibre Channel Transmitter Technology */
	SpecLongwaveLaserLC
	SpecElectricalInterEnclosure
	SpecElectricalIntraEnclosure
	SpecShortwaveLaserWithoutOFC
	SpecShortwaveLaserWithOFC
	SpecLongwaveLaserLL

	/* Fibre Channel Transmission Media */
	SpecTwinAxialPair
	SpecShieldedTwistedPair
	SpecMiniatureCoax
	SpecVideoCoax
	SpecMultimodeM6
	SpecMultimodeM5
	SpecMultimodeOM3
	SpecSingleMode

	/* Fibre Channel Speed */
	Spec1200MBpsPerChannel
	Spec800MBps
	Spec1600MBpsPerChannel
	Spec400MBps
	Spec3200MBpsPerChannel
	Spec200MBps
	Spec100MBps
)

func (s Specification) String() string {
	return map[Specification]string{
		Spec10GBaseLRM: "10GBASE-LRM",
		Spec10GBaseLR:  "10GBASE-LR",
		Spec10GBaseSR:  "10GBASE-SR",
		Spec40GBaseCR4: "40GBASE-CR4",
		Spec40GBaseSR4: "40GBASE-SR4",
		Spec40GBaseLR4: "40GBASE-LR4",
		Spec40GXLPPI:   "40G Active Cable (XLPPI)",

		SpecOC48LongReach:         "OC 48, long reach",
		SpecOC48IntermediateReach: "OC 48, intermediate reach",
		SpecOC48ShortReach:        "OC 48, short reach",

		SpecSAS24G: "SAS 24.0 Gbps",
		SpecSAS12G: "SAS 12.0 Gbps",
		SpecSAS6G:  "SAS 6.0 Gbps",
		SpecSAS3G:  "SAS 3.0 Gbps",

		Spec1000BaseT:  "1000BASE-T",
		Spec1000BaseCX: "1000BASE-CX",
		Spec1000BaseLX: "1000Base-LX",
		Spec1000BaseSX: "1000Base-SX",

		SpecVeryLongDistance:     "Very long distance (v)",
		SpecShortDistance:        "Short distance (S)",
		SpecIntermediateDistance: "Intermediate distance (I)",
		SpecLongDistance:         "Long distance (L)",
		SpecMediumDistance:       "Medium (M)",

		SpecLongwaveLaserLC:          "Longwave laser (LC)",
		SpecElectricalInterEnclosure: "Electrical inter-enclosure (EL)",
		SpecElectricalIntraEnclosure: "Electrical intra-enclosure",
		SpecShortwaveLaserWithoutOFC: "Shortwave laser w/o OFC (SN)",
		SpecShortwaveLaserWithOFC:    "Shortwave laser w OFC (SL)",
		SpecLongwaveLaserLL:          "Longwave laser (LL)",

		SpecTwinAxialPair:       "Twin Axial Pair (TP)",
		SpecShieldedTwistedPair: "Shielded Twisted Pair (TP)",
		SpecMiniatureCoax:       "Miniature Coax (MI)",
		SpecVideoCoax:           "Video Coax (TV)",
		SpecMultimodeM6:         "Multi-mode 62.5 um (M6)",
		SpecMultimodeM5:         "Multi-mode 50 um (M5)",
		SpecMultimodeOM3:        "Multi-mode 50 um (OM3)",
		SpecSingleMode:          "Single Mode (SM)",

		Spec1200MBpsPerChannel: "1200 MBps (per channel)",
		Spec800MBps:            "800 MBps",
		Spec1600MBpsPerChannel: "1600 MBps (per channel)",
		Spec400MBps:            "400 MBps",
		Spec3200MBpsPerChannel: "3200 MBps (per channel)",
		Spec200MBps:            "200 MBps",
		Spec100MBps:            "100 MBps",
	}[s]
}

func (s SpecificationCompliance) MarshalJSON() ([]byte, error) {
	ret := []string{}
	for specification, status := range s {
		if status {
			ret = append(ret, specification.String())
		}
	}
	return json.Marshal(ret)
}

var specificationComplianceMemoryMap = map[uint]map[uint]Specification{
	0x00: map[uint]Specification{
		0x06: Spec10GBaseLRM,
		0x05: Spec10GBaseLR,
		0x04: Spec10GBaseSR,
		0x03: Spec40GBaseCR4,
		0x02: Spec40GBaseSR4,
		0x01: Spec40GBaseLR4,
		0x00: Spec40GXLPPI,
	},
	0x01: map[uint]Specification{
		0x02: SpecOC48LongReach,
		0x01: SpecOC48IntermediateReach,
		0x00: SpecOC48ShortReach,
	},
	0x02: map[uint]Specification{
		0x07: SpecSAS24G,
		0x06: SpecSAS12G,
		0x05: SpecSAS6G,
		0x04: SpecSAS3G,
	},
	0x03: map[uint]Specification{
		0x03: Spec1000BaseT,
		0x02: Spec1000BaseCX,
		0x01: Spec1000BaseLX,
		0x00: Spec1000BaseSX,
	},
	0x04: map[uint]Specification{
		0x07: SpecVeryLongDistance,
		0x06: SpecShortDistance,
		0x05: SpecIntermediateDistance,
		0x04: SpecLongDistance,
		0x03: SpecMediumDistance,
		0x01: SpecLongwaveLaserLL,
		0x00: SpecElectricalInterEnclosure,
	},
	0x05: map[uint]Specification{
		0x07: SpecElectricalIntraEnclosure,
		0x06: SpecShortwaveLaserWithoutOFC,
		0x05: SpecShortwaveLaserWithOFC,
		0x04: SpecLongwaveLaserLC,
	},
	0x06: map[uint]Specification{
		0x07: SpecTwinAxialPair,
		0x06: SpecShieldedTwistedPair,
		0x05: SpecMiniatureCoax,
		0x04: SpecVideoCoax,
		0x03: SpecMultimodeM6,
		0x02: SpecMultimodeM5,
		0x01: SpecMultimodeOM3,
		0x00: SpecSingleMode,
	},
	0x07: map[uint]Specification{
		0x07: Spec1200MBpsPerChannel,
		0x06: Spec800MBps,
		0x05: Spec1600MBpsPerChannel,
		0x04: Spec400MBps,
		0x03: Spec3200MBpsPerChannel,
		0x02: Spec200MBps,
		0x01: Spec100MBps,
	},
}

func (s SpecificationCompliance) IsNonOpticalImplementation() bool {
	return s[Spec40GBaseCR4]
}

func NewSpecificationCompliance(raw [8]byte) SpecificationCompliance {
	s := SpecificationCompliance{}

	for byteOffset, bitMap := range specificationComplianceMemoryMap {
		for bitOffset, specifcation := range bitMap {
			s[specifcation] = raw[byteOffset]&(1<<bitOffset) > 0
		}
	}
	return s
}
