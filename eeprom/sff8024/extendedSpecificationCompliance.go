package sff8024

import (
    "encoding/json"
    "fmt"
)

// ExtendedSpecificationCompliance as defined in SFF-8-24
type ExtendedSpecificationCompliance byte

const (
    // ExtendedSpecificationCompliance100GAOC25GAUI 100G AOC (Active Optical Cable) or 25GAUI C2M AOC. Providing a worst BER of 5 × 10-5
    ExtendedSpecificationCompliance100GAOC25GAUI                                    ExtendedSpecificationCompliance = 0x01
    // ExtendedSpecificationCompliance100GBaseSR425GBaseSR 100GBASE-SR4 or 25GBASE-SR
    ExtendedSpecificationCompliance100GBaseSR425GBaseSR                              ExtendedSpecificationCompliance = 0x02
    // ExtendedSpecificationCompliance100GBaseLR425GBaseLR 100GBASE-LR4 or 25GBASE-LR
    ExtendedSpecificationCompliance100GBaseLR425GBaseLR                              ExtendedSpecificationCompliance = 0x03
    // ExtendedSpecificationCompliance100GBaseER425GBaseER 100GBASE-LR4 or 25GBASE-ER
    ExtendedSpecificationCompliance100GBaseER425GBaseER                              ExtendedSpecificationCompliance = 0x04
    // ExtendedSpecificationCompliance100GBaseSR10 100GBASE-SR10
    ExtendedSpecificationCompliance100GBaseSR10                                       ExtendedSpecificationCompliance = 0x05
    // ExtendedSpecificationCompliance100GCWDM4 100G CWDM4
    ExtendedSpecificationCompliance100GCWDM4                                         ExtendedSpecificationCompliance = 0x06
    // ExtendedSpecificationCompliance100GPSM4 100G PSM4 Parallel SMF
    ExtendedSpecificationCompliance100GPSM4                                          ExtendedSpecificationCompliance = 0x07
    // ExtendedSpecificationCompliance100GACC25GAUI 100G ACC (Active Copper Cable) or 25GAUI C2M ACC. Providing a worst BER of 5 × 10 -5
    ExtendedSpecificationCompliance100GACC25GAUI                                    ExtendedSpecificationCompliance = 0x08
    // ExtendedSpecificationComplianceObsolete Obsolete (assigned before 100G CWDM4 MSA required FEC)
    ExtendedSpecificationComplianceObsolete                                           ExtendedSpecificationCompliance = 0x09
    // ExtendedSpecificationCompliance100GBaseCR425GBaseCRCA25GL50GBaseCR2Clause91FEC 100GBASE-CR4, 25GBASE-CR CA-25G-L or 50GBASE-CR2 with RS (Clause91) FEC
    ExtendedSpecificationCompliance100GBaseCR425GBaseCRCA25GL50GBaseCR2Clause91FEC ExtendedSpecificationCompliance = 0x0B
    // ExtendedSpecificationCompliance25GBaseCRCA25GS50GBaseCR2Clause74FEC 25GBASE-CR CA-25G-S or 50GBASE-CR2 with BASE-R (Clause 74 Fire code) FEC
    ExtendedSpecificationCompliance25GBaseCRCA25GS50GBaseCR2Clause74FEC             ExtendedSpecificationCompliance = 0x0C
    // ExtendedSpecificationCompliance25GBaseCRCA25GN50GBaseCR2NoFEC 25GBASE-CR CA-25G-N or 50GBASE-CR2 with no FEC
    ExtendedSpecificationCompliance25GBaseCRCA25GN50GBaseCR2NoFEC                   ExtendedSpecificationCompliance = 0x0D
    // ExtendedSpecificationCompliance40GBaseER4 40GBASE-ER4
    ExtendedSpecificationCompliance40GBaseER4                                         ExtendedSpecificationCompliance = 0x10
    // ExtendedSpecificationCompliance4x10GBaseSR 4 x 10GBASE-SR
    ExtendedSpecificationCompliance4x10GBaseSR                                        ExtendedSpecificationCompliance = 0x11
    // ExtendedSpecificationCompliance40GPSM4 40G PSM4 Parallel SMF
    ExtendedSpecificationCompliance40GPSM4                                           ExtendedSpecificationCompliance = 0x12
    // ExtendedSpecificationComplianceG9591Profile112D1 G959.1 profile P1I1-2D1 (10709 MBd, 2km, 1310 nm SM)
    ExtendedSpecificationComplianceG9591Profile112D1                              ExtendedSpecificationCompliance = 0x13
    // ExtendedSpecificationComplianceG9591Profile1S12D2 G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550 nm SM)
    ExtendedSpecificationComplianceG9591Profile1S12D2                              ExtendedSpecificationCompliance = 0x14
    // ExtendedSpecificationComplianceG9591Profile1L12D2 G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550 nm SM)
    ExtendedSpecificationComplianceG9591Profile1L12D2                              ExtendedSpecificationCompliance = 0x15
    // ExtendedSpecificationCompliance10GBaseTSFI 10GBASE-T with SFI electrical interface
    ExtendedSpecificationCompliance10GBaseTSFI                                        ExtendedSpecificationCompliance = 0x16
    // ExtendedSpecificationCompliance100GCLR4 100G CLR4
    ExtendedSpecificationCompliance100GCLR4                                          ExtendedSpecificationCompliance = 0x17
    // ExtendedSpecificationCompliance100GAOC25GAUIC2MAOC 100G AOC or 25GAUI C2M AOC. Providing a worst BER of 10 -12 or below
    ExtendedSpecificationCompliance100GAOC25GAUIC2MAOC                            ExtendedSpecificationCompliance = 0x18
    // ExtendedSpecificationCompliance100GAOC25GAUIC2MACC 100G ACC or 25GAUI C2M ACC. Providing a worst BER of 10 -12 or below
    ExtendedSpecificationCompliance100GAOC25GAUIC2MACC                            ExtendedSpecificationCompliance = 0x19
    // ExtendedSpecificationCompliance100GEDWDM2 100GE-DWDM2 (DWDM transceiver using 2 wavelengths on a 1550 nm DWDM grid with a reach up to 80 km)
    ExtendedSpecificationCompliance100GEDWDM2                                        ExtendedSpecificationCompliance = 0x1A
    // ExtendedSpecificationCompliance100GWDM4 100G 1550nm WDM (4 wavelengths)
    ExtendedSpecificationCompliance100GWDM4                                          ExtendedSpecificationCompliance = 0x1B
    // ExtendedSpecificationCompliance10GBaseT 10GBASE-T Short Reach (30 meters)
    ExtendedSpecificationCompliance10GBaseT                                           ExtendedSpecificationCompliance = 0x1C
    // ExtendedSpecificationCompliance5GBaseT 5GBASE-T
    ExtendedSpecificationCompliance5GBaseT                                            ExtendedSpecificationCompliance = 0x1D
    // ExtendedSpecificationCompliance25GBaseT 2.5GBASE-T
    ExtendedSpecificationCompliance25GBaseT                                          ExtendedSpecificationCompliance = 0x1E
    // ExtendedSpecificationCompliance40GSWDM4 40G SWDM4
    ExtendedSpecificationCompliance40GSWDM4                                          ExtendedSpecificationCompliance = 0x1F
    // ExtendedSpecificationCompliance100GSWDM4 100G SWDM4
    ExtendedSpecificationCompliance100GSWDM4                                         ExtendedSpecificationCompliance = 0x20
    // ExtendedSpecificationCompliance100GPAM4BiDi 100G PAM4 BiDi
    ExtendedSpecificationCompliance100GPAM4BiDi                                     ExtendedSpecificationCompliance = 0x21
    // ExtendedSpecificationCompliance4WDM10MSA 4WDM-10 MSA (10km version of 100G CWDM4 with same RS(528,514) FEC in host system)
    ExtendedSpecificationCompliance4WDM10MSA                                         ExtendedSpecificationCompliance = 0x22
    // ExtendedSpecificationCompliance4WDM20 4WDM-20 MSA (20km version of 100GBASE-LR4 with RS(528,514) FEC in host system)
    ExtendedSpecificationCompliance4WDM20                            ExtendedSpecificationCompliance = 0x23
    // ExtendedSpecificationCompliance4WDM40 4WDM-40 MSA (40km reach with APD receiver and RS(528,514) FEC in host system)
    ExtendedSpecificationCompliance4WDM40                            ExtendedSpecificationCompliance = 0x24
    // ExtendedSpecificationCompliance100GBaseDR 100GBASE-DR (clause 140), CAUI-4 (no FEC)
    ExtendedSpecificationCompliance100GBaseDR                        ExtendedSpecificationCompliance = 0x25
    // ExtendedSpecificationCompliance100GFR100GBaseFR1 100G-FR or 100GBASE-FR1 (clause 140), CAUI-4 (no FEC)
    ExtendedSpecificationCompliance100GFR100GBaseFR1                ExtendedSpecificationCompliance = 0x26
    // ExtendedSpecificationCompliance100GLR100GBaseLR1 100G-LR or 100GBASE-LR1 (clause 140), CAUI-4 (no FEC)
    ExtendedSpecificationCompliance100GLR100GBaseLR1               ExtendedSpecificationCompliance = 0x27
    // ExtendedSpecificationComplianceACC50GAUI100GAUI2200GAUI4BER6 Active Copper Cable with 50GAUI, 100GAUI-2 or 200GAUI-4 C2M. Providing a worst BER of 10-6 or below
    ExtendedSpecificationComplianceACC50GAUI100GAUI2200GAUI4BER6  ExtendedSpecificationCompliance = 0x30
    // ExtendedSpecificationComplianceAOC50GAUI100GAUI2200GAUI4BER6 4WDM-40 MSA (40km reach with APD receiver and RS(528,514) FEC in host system)
    ExtendedSpecificationComplianceAOC50GAUI100GAUI2200GAUI4BER6  ExtendedSpecificationCompliance = 0x31
    // ExtendedSpecificationComplianceACC50GAUI100GAUI2200GAUI4BER4 Active Copper Cable with 50GAUI, 100GAUI-2 or 200GAUI-4 C2M. Providing a worst BER of 2.6x10-4 for ACC, 10-5 for AUI, or below
    ExtendedSpecificationComplianceACC50GAUI100GAUI2200GAUI4BER4  ExtendedSpecificationCompliance = 0x32
    // ExtendedSpecificationComplianceAOC50GAUI100GAUI2200GAUI4BER4 Active Optical Cable with 50GAUI, 100GAUI-2 or 200GAUI-4 C2M. Providing a worst BER of 2.6x10-4 for AOC, 10-5 for AUI, or below
    ExtendedSpecificationComplianceAOC50GAUI100GAUI2200GAUI4BER4  ExtendedSpecificationCompliance = 0x33
    // ExtendedSpecificationCompliance50GBaseCR100GBaseCR2200GBaseCR4 50GBASE-CR, 100GBASE-CR2, or 200GBASE-CR4
    ExtendedSpecificationCompliance50GBaseCR100GBaseCR2200GBaseCR4 ExtendedSpecificationCompliance = 0x40
    // ExtendedSpecificationCompliance50GBaseSR100GBaseSR2200GBaseSR2 50GBASE-SR, 100GBASE-SR2, or 200GBASE-SR4
    ExtendedSpecificationCompliance50GBaseSR100GBaseSR2200GBaseSR2 ExtendedSpecificationCompliance = 0x41
    // ExtendedSpecificationCompliance50GBaseFR200GBaseDR4 50GBASE-FR or 200GBASE-DR4
    ExtendedSpecificationCompliance50GBaseFR200GBaseDR4             ExtendedSpecificationCompliance = 0x42
    // ExtendedSpecificationCompliance200GBaseFR4 200GBASE-FR4
    ExtendedSpecificationCompliance200GBaseFR4                       ExtendedSpecificationCompliance = 0x43
    // ExtendedSpecificationCompliance200GPSM4 200G 1550 nm PSM4
    ExtendedSpecificationCompliance200GPSM4                         ExtendedSpecificationCompliance = 0x44
    // ExtendedSpecificationCompliance50GBaseLR 50GBASE-LR
    ExtendedSpecificationCompliance50GBaseLR                         ExtendedSpecificationCompliance = 0x45
    // ExtendedSpecificationCompliance200GBaseLR4 200GBASE-LR4
    ExtendedSpecificationCompliance200GBaseLR4                       ExtendedSpecificationCompliance = 0x46
    // ExtendedSpecificationCompliance64GFCEA 64GFC EA
    ExtendedSpecificationCompliance64GFCEA                          ExtendedSpecificationCompliance = 0x50
    // ExtendedSpecificationCompliance64GFCSW 64GFC SW
    ExtendedSpecificationCompliance64GFCSW                          ExtendedSpecificationCompliance = 0x51
    // ExtendedSpecificationCompliance64GFCLW 64GFC LW
    ExtendedSpecificationCompliance64GFCLW                          ExtendedSpecificationCompliance = 0x52
    // ExtendedSpecificationCompliance128GFCEA 128GFC EA
    ExtendedSpecificationCompliance128GFCEA                         ExtendedSpecificationCompliance = 0x53
    // ExtendedSpecificationCompliance128GFCSW 128GFC SW
    ExtendedSpecificationCompliance128GFCSW                         ExtendedSpecificationCompliance = 0x54
    // ExtendedSpecificationCompliance128GFCLW 128GFC LW
    ExtendedSpecificationCompliance128GFCLW                         ExtendedSpecificationCompliance = 0x55
)

func (e ExtendedSpecificationCompliance) String() string {
    return map[ExtendedSpecificationCompliance]string{
        ExtendedSpecificationCompliance100GAOC25GAUI: "100G AOC (Active Optical Cable) or 25GAUI C2M AOC. Providing a worst BER of 5 × 10-5",
        ExtendedSpecificationCompliance100GBaseSR425GBaseSR: "100GBASE-SR4 or 25GBASE-SR",
        ExtendedSpecificationCompliance100GBaseLR425GBaseLR: "100GBASE-LR4 or 25GBASE-LR",
        ExtendedSpecificationCompliance100GBaseER425GBaseER: "100GBASE-LR4 or 25GBASE-ER",
        ExtendedSpecificationCompliance100GBaseSR10: "100GBASE-SR10",
        ExtendedSpecificationCompliance100GCWDM4: "100G CWDM4",
        ExtendedSpecificationCompliance100GPSM4: "100G PSM4 Parallel SMF",
        ExtendedSpecificationCompliance100GACC25GAUI: "100G ACC (Active Copper Cable) or 25GAUI C2M ACC. Providing a worst BER of 5 × 10 -5",
        ExtendedSpecificationComplianceObsolete: "Obsolete (assigned before 100G CWDM4 MSA required FEC)",
        ExtendedSpecificationCompliance100GBaseCR425GBaseCRCA25GL50GBaseCR2Clause91FEC: "100GBASE-CR4, 25GBASE-CR CA-25G-L or 50GBASE-CR2 with RS (Clause91) FEC",
        ExtendedSpecificationCompliance25GBaseCRCA25GS50GBaseCR2Clause74FEC: "25GBASE-CR CA-25G-S or 50GBASE-CR2 with BASE-R (Clause 74 Fire code) FEC",
        ExtendedSpecificationCompliance25GBaseCRCA25GN50GBaseCR2NoFEC: "25GBASE-CR CA-25G-N or 50GBASE-CR2 with no FEC",
        ExtendedSpecificationCompliance40GBaseER4: "40GBASE-ER4",
        ExtendedSpecificationCompliance4x10GBaseSR: "4 x 10GBASE-SR",
        ExtendedSpecificationCompliance40GPSM4: "40G PSM4 Parallel SMF",
        ExtendedSpecificationComplianceG9591Profile112D1: "G959.1 profile P1I1-2D1 (10709 MBd, 2km, 1310 nm SM)",
        ExtendedSpecificationComplianceG9591Profile1S12D2: "G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550 nm SM)",
        ExtendedSpecificationComplianceG9591Profile1L12D2: "G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550 nm SM)",
        ExtendedSpecificationCompliance10GBaseTSFI: "10GBASE-T with SFI electrical interface",
        ExtendedSpecificationCompliance100GCLR4: "100G CLR4",
        ExtendedSpecificationCompliance100GAOC25GAUIC2MAOC: "100G AOC or 25GAUI C2M AOC. Providing a worst BER of 10 -12 or below",
        ExtendedSpecificationCompliance100GAOC25GAUIC2MACC: "100G ACC or 25GAUI C2M ACC. Providing a worst BER of 10 -12 or below",
        ExtendedSpecificationCompliance100GEDWDM2: "100GE-DWDM2 (DWDM transceiver using 2 wavelengths on a 1550 nm DWDM grid with a reach up to 80 km)",
        ExtendedSpecificationCompliance100GWDM4: "100G 1550nm WDM (4 wavelengths)",
        ExtendedSpecificationCompliance10GBaseT: "10GBASE-T Short Reach (30 meters)",
        ExtendedSpecificationCompliance5GBaseT: "5GBASE-T",
        ExtendedSpecificationCompliance25GBaseT: "2.5GBASE-T",
        ExtendedSpecificationCompliance40GSWDM4: "100G SWDM4",
        ExtendedSpecificationCompliance100GPAM4BiDi: "100G PAM4 BiDi",
        ExtendedSpecificationCompliance4WDM10MSA: "4WDM-10 MSA (10km version of 100G CWDM4 with same RS(528,514) FEC in host system)",
        ExtendedSpecificationCompliance4WDM20: "4WDM-20 MSA (20km version of 100GBASE-LR4 with RS(528,514) FEC in host system)",
        ExtendedSpecificationCompliance4WDM40: "4WDM-40 MSA (40km reach with APD receiver and RS(528,514) FEC in host system)",
        ExtendedSpecificationCompliance100GBaseDR: "100GBASE-DR (clause 140), CAUI-4 (no FEC)",
        ExtendedSpecificationCompliance100GFR100GBaseFR1: "100G-FR or 100GBASE-FR1 (clause 140), CAUI-4 (no FEC)",
        ExtendedSpecificationCompliance100GLR100GBaseLR1: "100G-LR or 100GBASE-LR1 (clause 140), CAUI-4 (no FEC)",
        ExtendedSpecificationComplianceACC50GAUI100GAUI2200GAUI4BER6: "Active Copper Cable with 50GAUI, 100GAUI-2 or 200GAUI-4 C2M. Providing a worst BER of 10-6 or below",
        ExtendedSpecificationComplianceAOC50GAUI100GAUI2200GAUI4BER6: "4WDM-40 MSA (40km reach with APD receiver and RS(528,514) FEC in host system)",
        ExtendedSpecificationComplianceACC50GAUI100GAUI2200GAUI4BER4: "Active Copper Cable with 50GAUI, 100GAUI-2 or 200GAUI-4 C2M. Providing a worst BER of 2.6x10-4 for ACC, 10-5 for AUI, or below",
        ExtendedSpecificationComplianceAOC50GAUI100GAUI2200GAUI4BER4: "Active Optical Cable with 50GAUI, 100GAUI-2 or 200GAUI-4 C2M. Providing a worst BER of 2.6x10-4 for AOC, 10-5 for AUI, or below",
        ExtendedSpecificationCompliance50GBaseCR100GBaseCR2200GBaseCR4: "50GBASE-CR, 100GBASE-CR2, or 200GBASE-CR4",
        ExtendedSpecificationCompliance50GBaseSR100GBaseSR2200GBaseSR2: "50GBASE-SR, 100GBASE-SR2, or 200GBASE-SR4",
        ExtendedSpecificationCompliance50GBaseFR200GBaseDR4: "50GBASE-FR or 200GBASE-DR4",
        ExtendedSpecificationCompliance200GBaseFR4: "200GBASE-FR4",
        ExtendedSpecificationCompliance200GPSM4: "200G 1550 nm PSM4",
        ExtendedSpecificationCompliance50GBaseLR: "200GBASE-LR4",
        ExtendedSpecificationCompliance64GFCEA: "64GFC EA",
        ExtendedSpecificationCompliance64GFCSW: "64GFC SW",
        ExtendedSpecificationCompliance64GFCLW: "64GFC LW",
        ExtendedSpecificationCompliance128GFCEA: "128GFC EA",
        ExtendedSpecificationCompliance128GFCSW: "128GFC SW",
        ExtendedSpecificationCompliance128GFCLW: "128GFC LW",
    }[e]
}

// MarshalJSON implements the encoding/json/Marshaler interface's MarshalJSON function
func (e ExtendedSpecificationCompliance) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]string {
        "ascii": e.String(),
        "hex":   fmt.Sprintf("%#02x", byte(e)),
    })  
}
