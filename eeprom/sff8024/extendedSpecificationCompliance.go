package SFF8024

type ExtendedSpecificationCompliance byte

const (
	ExtendedSpecificationCompliance100G_AOC_25GAUI                                    ExtendedSpecificationCompliance = 0x01
	ExtendedSpecificationCompliance100GBaseSR4_25GBaseSR                              ExtendedSpecificationCompliance = 0x02
	ExtendedSpecificationCompliance100GBaseLR4_25GBaseLR                              ExtendedSpecificationCompliance = 0x03
	ExtendedSpecificationCompliance100GBaseER4_25GBaseER                              ExtendedSpecificationCompliance = 0x04
	ExtendedSpecificationCompliance100GBaseSR10                                       ExtendedSpecificationCompliance = 0x05
	ExtendedSpecificationCompliance100G_CWDM4                                         ExtendedSpecificationCompliance = 0x06
	ExtendedSpecificationCompliance100G_PSM4                                          ExtendedSpecificationCompliance = 0x07
	ExtendedSpecificationCompliance100G_ACC_25GAUI                                    ExtendedSpecificationCompliance = 0x08
	ExtendedSpecificationComplianceObsolete                                           ExtendedSpecificationCompliance = 0x09
	ExtendedSpecificationCompliance100GBaseCR4_25GBaseCR_CA25GL_50GBaseCR2Clause91FEC ExtendedSpecificationCompliance = 0x0B
	ExtendedSpecificationCompliance25GBaseCR_CA25GS_50GBaseCR2Clause74FEC             ExtendedSpecificationCompliance = 0x0C
	ExtendedSpecificationCompliance25GBaseCR_CA25GN_50GBaseCR2NoFEC                   ExtendedSpecificationCompliance = 0x0D
	ExtendedSpecificationCompliance40GBaseER4                                         ExtendedSpecificationCompliance = 0x10
	ExtendedSpecificationCompliance4x10GBaseSR                                        ExtendedSpecificationCompliance = 0x11
	ExtendedSpecificationCompliance40G_PSM4                                           ExtendedSpecificationCompliance = 0x12
	ExtendedSpecificationComplianceG959_1_Profile1_1_2D1                              ExtendedSpecificationCompliance = 0x13
	ExtendedSpecificationComplianceG959_1_Profile1S1_2D2                              ExtendedSpecificationCompliance = 0x14
	ExtendedSpecificationComplianceG959_1_Profile1L1_2D2                              ExtendedSpecificationCompliance = 0x15
	ExtendedSpecificationCompliance10GBaseTSFI                                        ExtendedSpecificationCompliance = 0x16
	ExtendedSpecificationCompliance100G_CLR4                                          ExtendedSpecificationCompliance = 0x17
	ExtendedSpecificationCompliance100G_AOC_25GAUI_C2M_AOC                            ExtendedSpecificationCompliance = 0x18
	ExtendedSpecificationCompliance100G_AOC_25GAUI_C2M_ACC                            ExtendedSpecificationCompliance = 0x19
	ExtendedSpecificationCompliance100GE_DWDM2                                        ExtendedSpecificationCompliance = 0x1A
	ExtendedSpecificationCompliance100G_WDM4                                          ExtendedSpecificationCompliance = 0x1B
	ExtendedSpecificationCompliance10GBaseT                                           ExtendedSpecificationCompliance = 0x1C
	ExtendedSpecificationCompliance5GBaseT                                            ExtendedSpecificationCompliance = 0x1D
	ExtendedSpecificationCompliance2_5GBaseT                                          ExtendedSpecificationCompliance = 0x1E
	ExtendedSpecificationCompliance40G_SWDM4                                          ExtendedSpecificationCompliance = 0x1F
	ExtendedSpecificationCompliance100G_SWDM4                                         ExtendedSpecificationCompliance = 0x20
	ExtendedSpecificationCompliance100G_PAM4_BiDi                                     ExtendedSpecificationCompliance = 0x21
	ExtendedSpecificationCompliance4WDM10_MSA                                         ExtendedSpecificationCompliance = 0x22

	ExtendedSpecificationCompliance4WDM20                            ExtendedSpecificationCompliance = 0x23
	ExtendedSpecificationCompliance4WDM40                            ExtendedSpecificationCompliance = 0x24
	ExtendedSpecificationCompliance100GBaseDR                        ExtendedSpecificationCompliance = 0x25
	ExtendedSpecificationCompliance100GFR_100GBaseFR1                ExtendedSpecificationCompliance = 0x26
	ExtendedSpecificationCompliance100G_LR_100GBaseLR1               ExtendedSpecificationCompliance = 0x27
	ExtendedSpecificationComplianceACC50GAUI_100GAUI2_200GAUI4_BER6  ExtendedSpecificationCompliance = 0x30
	ExtendedSpecificationComplianceAOC50GAUI_100GAUI2_200GAUI4_BER6  ExtendedSpecificationCompliance = 0x31
	ExtendedSpecificationComplianceACC50GAUI_100GAUI2_200GAUI4_BER4  ExtendedSpecificationCompliance = 0x32
	ExtendedSpecificationComplianceAOC50GAUI_100GAUI2_200GAUI4_BER4  ExtendedSpecificationCompliance = 0x33
	ExtendedSpecificationCompliance50GBaseCR_100GBaseCR2_200GBaseCR4 ExtendedSpecificationCompliance = 0x40
	ExtendedSpecificationCompliance50GBaseSR_100GBaseSR2_200GBaseSR2 ExtendedSpecificationCompliance = 0x41
	ExtendedSpecificationCompliance50GBaseFR_200GBaseDR4             ExtendedSpecificationCompliance = 0x42
	ExtendedSpecificationCompliance200GBaseFR4                       ExtendedSpecificationCompliance = 0x43
	ExtendedSpecificationCompliance200G_PSM4                         ExtendedSpecificationCompliance = 0x44
	ExtendedSpecificationCompliance50GBaseLR                         ExtendedSpecificationCompliance = 0x45
	ExtendedSpecificationCompliance200GBaseLR4                       ExtendedSpecificationCompliance = 0x46
	ExtendedSpecificationCompliance64GFC_EA                          ExtendedSpecificationCompliance = 0x50
	ExtendedSpecificationCompliance64GFC_SW                          ExtendedSpecificationCompliance = 0x51
	ExtendedSpecificationCompliance64GFC_LW                          ExtendedSpecificationCompliance = 0x52
	ExtendedSpecificationCompliance128GFC_EA                         ExtendedSpecificationCompliance = 0x53
	ExtendedSpecificationCompliance128GFC_SW                         ExtendedSpecificationCompliance = 0x54
	ExtendedSpecificationCompliance128GFC_LW                         ExtendedSpecificationCompliance = 0x55
)
