package SFF8079

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Compliance map[ComplianceFlag]bool
type ComplianceFlag int

const (
	/* 10G Ethernet ComplianceFlag Codes */
	ComplianceFlag10GBaseER ComplianceFlag = iota
	ComplianceFlag10GBaseLRM
	ComplianceFlag10GBaseLR
	ComplianceFlag10GBaseSR

	/* Infiband ComplianceFlag Codes */
	ComplianceFlag1XSX
	ComplianceFlag1XLX
	ComplianceFlag1XCopperActive
	ComplianceFlag1XCopperPassive

	/* ESCON ComplianceFlag Codes */
	ComplianceFlagEsconMMF1310Led
	ComplianceFlagEsconMMF1310Laser

	/* SONET ComplianceFlag Codes */
	ComplianceFlagOC192ShortReach
	ComplianceFlagSonetReachSpecifierBit1
	ComplianceFlagSonetReachSpecifierBit2
	ComplianceFlagOC48LongReach
	ComplianceFlagOC48IntermediateReach
	ComplianceFlagOC48ShortReach
	ComplianceFlagOC12SingleModeLongReach
	ComplianceFlagOC12SingleModeIntermediateReach
	ComplianceFlagOC12ShortReach
	ComplianceFlagOC3SingleModeLongReach
	ComplianceFlagOC3SingleModeIntermediateReach
	ComplianceFlagOC3SingleModeShortReach

	/* Ethernet ComplianceFlag Codes */
	ComplianceFlagBasePX
	ComplianceFlagBaseBX10
	ComplianceFlag100BaseFX
	ComplianceFlag100BaseLX
	ComplianceFlag1000BaseT
	ComplianceFlag1000BaseCX
	ComplianceFlag1000BaseLX
	ComplianceFlag1000BaseSX

	/* Fibre Channel Link Length */
	ComplianceFlagVeryLongDistance
	ComplianceFlagShortDistance
	ComplianceFlagIntermediateDistance
	ComplianceFlagLongDistance
	ComplianceFlagMediumDistance

	/* Fibre Channel Technology */
	ComplianceFlagShortwaveLaserSA
	ComplianceFlagLongwaveLaserLC
	ComplianceFlagElectricalInterEnclosureEL
	ComplianceFlagElectricalIntraEnclosureEL
	ComplianceFlagShortwaveLaserWithoutOFC
	ComplianceFlagShortwaveLaserWithOFC
	ComplianceFlagLongwaveLaserLL

	/* SFP+ Technology */
	ComplianceFlagActiveCable
	ComplianceFlagPassiveCable

	/* Fibre Channel Transmission Media */
	ComplianceFlagTwinAxialPair
	ComplianceFlagTwistedPair
	ComplianceFlagMiniatureCoax
	ComplianceFlagVideoCoax
	ComplianceFlagMultimodeM6
	ComplianceFlagMultimodeM5
	ComplianceFlagSingleMode

	/* FibreChannelSpeed */
	ComplianceFlag1200MBps
	ComplianceFlag800MBps
	ComplianceFlag1600MBps
	ComplianceFlag400MBps
	ComplianceFlag3200MBps
	ComplianceFlag200MBps
	ComplianceFlag100MBps
)

var complianceMemoryMap = map[uint]map[uint]ComplianceFlag{
	0x00: map[uint]ComplianceFlag{
		0x07: ComplianceFlag10GBaseER,
		0x06: ComplianceFlag10GBaseLRM,
		0x05: ComplianceFlag10GBaseLR,
		0x04: ComplianceFlag10GBaseSR,
		0x03: ComplianceFlag1XSX,
		0x02: ComplianceFlag1XLX,
		0x01: ComplianceFlag1XCopperActive,
		0x00: ComplianceFlag1XCopperPassive,
	},
	0x01: map[uint]ComplianceFlag{
		0x07: ComplianceFlagEsconMMF1310Led,
		0x06: ComplianceFlagEsconMMF1310Laser,
		0x05: ComplianceFlagOC192ShortReach,
		0x04: ComplianceFlagSonetReachSpecifierBit1,
		0x03: ComplianceFlagSonetReachSpecifierBit2,
		0x02: ComplianceFlagOC48LongReach,
		0x01: ComplianceFlagOC48IntermediateReach,
		0x00: ComplianceFlagOC48ShortReach,
	},
	0x02: map[uint]ComplianceFlag{
		0x06: ComplianceFlagOC12SingleModeLongReach,
		0x05: ComplianceFlagOC12SingleModeIntermediateReach,
		0x04: ComplianceFlagOC12ShortReach,
		0x02: ComplianceFlagOC3SingleModeLongReach,
		0x01: ComplianceFlagOC3SingleModeIntermediateReach,
		0x00: ComplianceFlagOC3SingleModeShortReach,
	},
	0x03: map[uint]ComplianceFlag{
		0x07: ComplianceFlagBasePX,
		0x06: ComplianceFlagBaseBX10,
		0x05: ComplianceFlag100BaseFX,
		0x04: ComplianceFlag100BaseLX,
		0x03: ComplianceFlag1000BaseT,
		0x02: ComplianceFlag1000BaseCX,
		0x01: ComplianceFlag1000BaseLX,
		0x00: ComplianceFlag1000BaseSX,
	},
	0x04: map[uint]ComplianceFlag{
		0x07: ComplianceFlagVeryLongDistance,
		0x06: ComplianceFlagShortDistance,
		0x05: ComplianceFlagIntermediateDistance,
		0x04: ComplianceFlagLongDistance,
		0x03: ComplianceFlagMediumDistance,
		0x02: ComplianceFlagShortwaveLaserSA,
		0x01: ComplianceFlagLongwaveLaserLC,
		0x00: ComplianceFlagElectricalInterEnclosureEL,
	},
	0x05: map[uint]ComplianceFlag{
		0x07: ComplianceFlagElectricalIntraEnclosureEL,
		0x06: ComplianceFlagShortwaveLaserWithoutOFC,
		0x05: ComplianceFlagShortwaveLaserWithOFC,
		0x04: ComplianceFlagLongwaveLaserLL,
		0x03: ComplianceFlagActiveCable,
		0x02: ComplianceFlagPassiveCable,
	},
	0x06: map[uint]ComplianceFlag{
		0x07: ComplianceFlagTwinAxialPair,
		0x06: ComplianceFlagTwistedPair,
		0x05: ComplianceFlagMiniatureCoax,
		0x04: ComplianceFlagVideoCoax,
		0x03: ComplianceFlagMultimodeM6,
		0x02: ComplianceFlagMultimodeM5,
		0x00: ComplianceFlagSingleMode,
	},
	0x07: map[uint]ComplianceFlag{
		0x07: ComplianceFlag1200MBps,
		0x06: ComplianceFlag800MBps,
		0x05: ComplianceFlag1600MBps,
		0x04: ComplianceFlag400MBps,
		0x03: ComplianceFlag3200MBps,
		0x02: ComplianceFlag200MBps,
		// 0x01 special use
		0x00: ComplianceFlag100MBps,
	},
}

func (c ComplianceFlag) String() string {
	return map[ComplianceFlag]string{
		ComplianceFlag10GBaseER:  "10G Base-ER",
		ComplianceFlag10GBaseLRM: "10G Base-LRM",
		ComplianceFlag10GBaseLR:  "10G BASE-LR",
		ComplianceFlag10GBaseSR:  "10G Base-SR",

		/* Infiband ComplianceFlag Codes */
		ComplianceFlag1XSX:            "1X SX",
		ComplianceFlag1XLX:            "1X LX",
		ComplianceFlag1XCopperActive:  "1X Copper Active",
		ComplianceFlag1XCopperPassive: "1X Copper Passive",

		/* ESCON ComplianceFlag Codes */
		ComplianceFlagEsconMMF1310Led:   "ESCON MMF, 1310nm LED",
		ComplianceFlagEsconMMF1310Laser: "ESCON MMF, 1310nm Laser",

		/* SONET ComplianceFlag Codes */
		ComplianceFlagOC192ShortReach:                 "OC-192, short reach",
		ComplianceFlagSonetReachSpecifierBit1:         "SONET reach specifier bit 1",
		ComplianceFlagSonetReachSpecifierBit2:         "SONET reach specifier bit 2",
		ComplianceFlagOC48LongReach:                   "OC-48, long reach",
		ComplianceFlagOC48IntermediateReach:           "OC-48, intermediate reach",
		ComplianceFlagOC48ShortReach:                  "OC-48, short reach",
		ComplianceFlagOC12SingleModeLongReach:         "OC-12, single mode, long reach",
		ComplianceFlagOC12SingleModeIntermediateReach: "OC-12, single mode, intermediate reach",
		ComplianceFlagOC12ShortReach:                  "OC-12, short reach",
		ComplianceFlagOC3SingleModeLongReach:          "OC-3, single mode, long reach",
		ComplianceFlagOC3SingleModeIntermediateReach:  "OC-3, single mode, intermediate reach",
		ComplianceFlagOC3SingleModeShortReach:         "OC-3, single mode, short reach",

		/* Ethernet ComplianceFlag Codes */
		ComplianceFlagBasePX:     "BASE-PX",
		ComplianceFlagBaseBX10:   "BASE-BX10",
		ComplianceFlag100BaseFX:  "100BASE-FX",
		ComplianceFlag100BaseLX:  "100BASE-LX/LX10",
		ComplianceFlag1000BaseT:  "1000BaseT",
		ComplianceFlag1000BaseCX: "1000BaseCX",
		ComplianceFlag1000BaseLX: "1000BaseLX",
		ComplianceFlag1000BaseSX: "1000BaseSX",

		/* Fibre Channel Link Length */
		ComplianceFlagVeryLongDistance:     "very long distance (V)",
		ComplianceFlagShortDistance:        "short distance (S)",
		ComplianceFlagIntermediateDistance: "intermediate distance (I)",
		ComplianceFlagLongDistance:         "long distance (L)",
		ComplianceFlagMediumDistance:       "medium distance (M)",

		/* Fibre Channel Technology */
		ComplianceFlagShortwaveLaserSA:           "Shortwave laser, linear Rx (SA)",
		ComplianceFlagLongwaveLaserLC:            "Longwave laser (LC)",
		ComplianceFlagElectricalInterEnclosureEL: "Electrical inter-enclosure (EL)",
		ComplianceFlagElectricalIntraEnclosureEL: "Electrical intra-enclosure (EL)",
		ComplianceFlagShortwaveLaserWithoutOFC:   "Shortwave laser w/o OFC (SN)",
		ComplianceFlagShortwaveLaserWithOFC:      "Shortwave laser with OFC (SL)",
		ComplianceFlagLongwaveLaserLL:            "Longwave laser (LL)",

		/* SFP+ Technology */
		ComplianceFlagActiveCable:  "Actice Cable",
		ComplianceFlagPassiveCable: "Passive Cable",

		/* Fibre Channel Transmission Media */
		ComplianceFlagTwinAxialPair: "Twin Axial Pair (TP)",
		ComplianceFlagTwistedPair:   "Twisted Pair (TP)",
		ComplianceFlagMiniatureCoax: "Miniature Coax (MI)",
		ComplianceFlagVideoCoax:     "Video Coax (TV)",
		ComplianceFlagMultimodeM6:   "Multimode, 62.5um (M6)",
		ComplianceFlagMultimodeM5:   "Multimode, 50um (M5, M5E)",
		ComplianceFlagSingleMode:    "Single Mode (SM)",

		/* FibreChannelSpeed */
		ComplianceFlag1200MBps: "12000 MBps",
		ComplianceFlag800MBps:  "800 MBps",
		ComplianceFlag1600MBps: "1600 MBps",
		ComplianceFlag400MBps:  "400 MBps",
		ComplianceFlag3200MBps: "3200 MBps",
		ComplianceFlag200MBps:  "200 MBps",
		ComplianceFlag100MBps:  "100 MBps",
	}[c]
}

func (c Compliance) String() string {
	builder := &strings.Builder{}
	for complianceFlag, status := range c {
		if status {
			fmt.Fprintf(builder, " * %s\n", complianceFlag.String())
		}
	}
	return builder.String()
}

func (c Compliance) MarshalJSON() ([]byte, error) {
	ret := []string{}
	for complianceFlag, status := range c {
		if status {
			ret = append(ret, complianceFlag.String())
		}
	}
	return json.Marshal(ret)
}

func NewCompliance(raw [8]byte) Compliance {
	c := Compliance{}
	for byteOffset, bitMap := range complianceMemoryMap {
		for bitOffset, Compliance := range bitMap {
			c[Compliance] = raw[byteOffset]&(1<<bitOffset) > 0
		}
	}
	return c
}

func (c Compliance) IsSFPCableImplementation() bool {
	return c[ComplianceFlagActiveCable] || c[ComplianceFlagPassiveCable]
}
