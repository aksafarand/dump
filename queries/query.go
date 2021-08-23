package queries

type Queries struct {
	query3g    string
	query4g    string
	query4gulo string
}

func GetQueries(tech string) (qry1 string, qry2 string) {

	qry := &Queries{}
	qry.query3g = `SELECT
	[ADD UCELLSETUP].[CELLNAME] AS cell_name,
	[ADD UCELLALGOSWITCH].[NE Name] AS ne_name,
	[ADD UCELLALGOSWITCH].[CELLID] AS cell_id,
	iif(Instr([ADD UCELLSETUP].[CELLNAME],'_') < 10 or Instr([ADD UCELLSETUP].[CELLNAME],'-') < 10,iif(len(left([ADD UCELLSETUP].[CELLNAME],Instr([ADD UCELLSETUP].[CELLNAME],'_')-1)) > 10, left([ADD UCELLSETUP].[CELLNAME],Instr([ADD UCELLSETUP].[CELLNAME],'-')-1),left([ADD UCELLSETUP].[CELLNAME],Instr([ADD UCELLSETUP].[CELLNAME],'_')-1)),[ADD UCELLSETUP].[CELLNAME]) as siteid,
	[ADD UCELLALGOSWITCH].HSPAPLUSSWITCH_FDPCH_CAPABILITY_INVALID AS hspaplusswitch_fdpch_capability_invalid,
	[ADD UCELLCAC].LOADBALANCERATIO AS loadbalanceratio,
	[ADD UCELLCAC].MAXHSUPA2MSTTIUSERNUM AS maxhsupa2msttiusernum,
	[ADD UCELLCOALGOENHPARA].CELLCOALGOENHSWITCH_SPEC_USER_HSDPA_DISABLED_SWITCH AS cellcoalgoenhswitch_spec_user_hsdpa_disabled_switch,
	[ADD UCELLNFREQPRIOINFO].EARFCN AS earfcn,
	[ADD UCELLNFREQPRIOINFO].NPRIORITY AS npriority,
	[ADD UCELLNFREQPRIOINFO].THDTOHIGH AS thdtohigh,
	[ADD UCELLNFREQPRIOINFO].THDTOLOW AS thdtolow,
	[ADD UCELLNFREQPRIOINFO].EMEASBW AS emeasbw,
	[ADD UCELLNFREQPRIOINFO].EQRXLEVMIN AS eqrxlevmin,
	[ADD UCELLNFREQPRIOINFO].SUPCNOPGRPINDEX AS supcnopgrpindex,
	[ADD UCELLNFREQPRIOINFO].BLACKLSTCELLNUMBER AS blacklstcellnumber,
	[ADD UCELLNFREQPRIOINFO].RSRQSWITCH AS rsrqswitch,
	[SET UNBMPARA].PERFENHANCESWITCH3_PERFENH_CQI_0_H2D_SWITCH AS perfenhanceswitch3_perfenh_cqi_0_h2d_switch
   FROM 
  ((((
  [ADD UCELLALGOSWITCH] 
  LEFT JOIN [ADD UCELLCAC] ON ([ADD UCELLALGOSWITCH].CELLID = [ADD UCELLCAC].CELLID) AND ([ADD UCELLALGOSWITCH].[NE Name] = [ADD UCELLCAC].[NE Name])) 
  LEFT JOIN [ADD UCELLCOALGOENHPARA] ON ([ADD UCELLALGOSWITCH].CELLID = [ADD UCELLCOALGOENHPARA].CELLID) AND ([ADD UCELLALGOSWITCH].[NE Name] = [ADD UCELLCOALGOENHPARA].[NE Name])) 
  LEFT JOIN [ADD UCELLNFREQPRIOINFO] ON ([ADD UCELLALGOSWITCH].CELLID = [ADD UCELLNFREQPRIOINFO].CELLID) AND ([ADD UCELLALGOSWITCH].[NE Name] = [ADD UCELLNFREQPRIOINFO].[NE Name])) 
  LEFT JOIN [ADD UCELLSETUP] ON ([ADD UCELLALGOSWITCH].CELLID = [ADD UCELLSETUP].CELLID) AND ([ADD UCELLALGOSWITCH].[NE Name] = [ADD UCELLSETUP].[NE Name])) 
  LEFT JOIN [SET UNBMPARA] ON [ADD UCELLALGOSWITCH].[NE Name] = [SET UNBMPARA].[NE Name]`

	qry.query4g = `SELECT DISTINCT
	  CELL.[NE NAME] as ne_name,
	  CELL.DlBandWidth as dlbandwidth,
	  CELL.UlBandWidth as ulbandwidth,
	  CELLALGOSWITCH.PucchAlgoSwitch as pucchalgoswitch,
	  CELLALGOSWITCH.DlPcAlgoSwitch as dlpcalgoswitch,
	  CELLALGOSWITCH.NonStandardBwAlgoSw as nonstandardbwalgosw,
	  CELLALGOSWITCH.CqiAdjAlgoSwitch as cqiadjalgoswitch,
	  CELLALGOSWITCH.DlSchSwitch as dlschswitch,
	  CELLALGOSWITCH.UlSchSwitch as ulschswitch,
	  CELLDLSCHALGO.RBDamageNearPointIblerTh as rbdamagenearpointiblerth,
	  CELLDLSCHALGO.FDUEEnhAperCQITrigPeriod as fdueenhapercqitrigperiod,
	  CELLHOPARACFG.BLINDHOA1A2THDRSRP as blindhoa1a2thdrsrp,
	  CELLPDCCHALGO.PdcchCapacityImproveSwitch as pdcchcapacityimproveswitch,
	  CELLPDCCHALGO.PdcchPowerEnhancedSwitch as pdcchpowerenhancedswitch,
	  CELLPDCCHALGO.PdcchMaxCodeRate as pdcchmaxcoderate,
	  CELLPDCCHALGO.PdcchOutLoopAdjLowerLimit as pdcchoutloopadjlowerlimit,
	  CELLPDCCHALGO.PdcchSymNumSwitch as pdcchsymnumswitch,
	  CELLPDCCHALGO.CceRatioAdjSwitch as cceratioadjswitch,
	  CELLRESEL.MeasBandWidthCfgInd as measbandwidthcfgind,
	  CELLRESEL.CellReselPriority as cellreselpriority,
	  CELLRESEL.ThrshServLowQCfgInd as thrshservlowqcfgind,
	  CELLRESEL.SNonIntraSearch as snonintrasearch,
	  CELLRESEL.ThrshServLow as thrshservlow,
	  CELLSEL.QQualMin as qqualmin,
	  CQIADAPTIVECFG.PucchPeriodicCqiOptSwitch as pucchperiodiccqioptswitch,
	  EUTRANINTERNFREQ.DlEarfcn as dlearfcn,
	  EUTRANINTERNFREQ.MeasBandWidth as measbandwidth,
	  EUTRANINTERNFREQ.THRESHXLOW as threshxlow,
	  EUTRANINTERNFREQ.CELLRESELPRIORITYCFGIND as cellreselprioritycfgind,
	  INTERFREQHOGROUP.INTERFREQHOA1THDRSRP as interfreqhoa1thdrsrp,
	  INTERFREQHOGROUP.INTERFREQHOA2THDRSRP as interfreqhoa2thdrsrp,
	  INTERFREQHOGROUP.INTERFREQHOA4THDRSRP as interfreqhoa4thdrsrp,
	  INTERRATHOCOMM.InterRatHoA1A2TrigQuan as interrathoa1a2trigquan,
	  INTRAFREQHOGROUP.IntraFreqHoA3Hyst as intrafreqhoa3hyst,
	  INTRAFREQHOGROUP.IntraFreqHoA3Offset as intrafreqhoa3offset,
	  INTRARATHOCOMM.IntraFreqHoA3TrigQuan as intrafreqhoa3trigquan,
	  INTRARATHOCOMM.InterFreqHoA1A2TrigQuan as interfreqhoa1a2trigquan,
	  INTRARATHOCOMM.FreqPriInterFreqHoA1TrigQuan as freqpriinterfreqhoa1trigquan,
	  INTRARATHOCOMM.InterFreqHoA4TrigQuan as interfreqhoa4trigquan,
	  INTRARATHOCOMM.A3InterFreqHoA1A2TrigQuan as a3interfreqhoa1a2trigquan,
	  PDSCHCFG.ReferenceSignalPwr as referencesignalpwr,
	  PHICHCFG.PhichDuration as phichduration,
	  PHICHCFG.PhichResource as phichresource,
	  PUCCHCFG.Format3RbNum as format3rbnum,
	  PUCCHCFG.Max2CCAckChNum as max2ccackchnum,
	  RACHCFG.PrachFreqOffsetStrategy as prachfreqoffsetstrategy,
	  SPECTRUMCLOUD.SpectrumCloudEnhSwitch as spectrumcloudenhswitch,
	  SPECTRUMCLOUD.SpectrumCloudSwitch as spectrumcloudswitch,
	  SPECTRUMCLOUD.DlCchSendStrategy as dlcchsendstrategy,
	  SPECTRUMCLOUD.InterfPfmOptSwitch as interfpfmoptswitch,
	  SRSCFG.AnSrsSimuTrans as ansrssimutrans,
	  ULSPECTRUMSHRCONFIG.PrbRsvType as prbrsvtype,
	  ULSPECTRUMSHRCONFIG.LtePrbStartIndex as lteprbstartindex,
	  ULSPECTRUMSHRCONFIG.LtePrbEndIndex as lteprbendindex
	 FROM ((((((((((((((((((CELL
	 INNER JOIN CELLALGOSWITCH ON (CELL.LocalCellId = CELLALGOSWITCH.LocalCellId) AND (CELL.[NE NAME] = CELLALGOSWITCH.[NE NAME]))
	 INNER JOIN CELLDLSCHALGO ON (CELL.LocalCellId = CELLDLSCHALGO.LocalCellId) AND (CELL.[NE NAME] = CELLDLSCHALGO.[NE NAME]))
	 INNER JOIN CELLHOPARACFG ON (CELL.LocalCellId = CELLHOPARACFG.LocalCellId) AND (CELL.[NE NAME] = CELLHOPARACFG.[NE NAME]))
	 INNER JOIN CELLPDCCHALGO ON (CELL.LocalCellId = CELLPDCCHALGO.LocalCellId) AND (CELL.[NE NAME] = CELLPDCCHALGO.[NE NAME]))
	 INNER JOIN CELLRESEL ON (CELL.LocalCellId = CELLRESEL.LocalCellId) AND (CELL.[NE NAME] = CELLRESEL.[NE NAME]))
	 INNER JOIN CELLSEL ON (CELL.LocalCellId = CELLSEL.LocalCellId) AND (CELL.[NE NAME] = CELLSEL.[NE NAME]))
	 INNER JOIN CQIADAPTIVECFG ON CELL.[NE NAME] = CQIADAPTIVECFG.[NE NAME])
	 INNER JOIN EUTRANINTERNFREQ ON (CELL.LocalCellId = EUTRANINTERNFREQ.LocalCellId) AND (CELL.[NE NAME] = EUTRANINTERNFREQ.[NE NAME]))
	 INNER JOIN INTERFREQHOGROUP ON (CELL.LocalCellId = INTERFREQHOGROUP.LocalCellId) AND (CELL.[NE NAME] = INTERFREQHOGROUP.[NE NAME]))
	 INNER JOIN PUCCHCFG ON (CELL.LocalCellId = PUCCHCFG.LocalCellId) AND (CELL.[NE NAME] = PUCCHCFG.[NE NAME]))
	 INNER JOIN RACHCFG ON (CELL.LocalCellId = RACHCFG.LocalCellId) AND (CELL.[NE NAME] = RACHCFG.[NE NAME]))
	 INNER JOIN SPECTRUMCLOUD ON (CELL.LocalCellId = SPECTRUMCLOUD.LocalCellId) AND (CELL.[NE NAME] = SPECTRUMCLOUD.[NE NAME]))
	 INNER JOIN SRSCFG ON (CELL.LocalCellId = SRSCFG.LocalCellId) AND (CELL.[NE NAME] = SRSCFG.[NE NAME]))
	 INNER JOIN ULSPECTRUMSHRCONFIG ON CELL.[NE NAME] = ULSPECTRUMSHRCONFIG.[NE NAME])
	 INNER JOIN INTERRATHOCOMM ON CELL.[NE NAME] = INTERRATHOCOMM.[NE NAME])
	 INNER JOIN INTRAFREQHOGROUP ON (CELL.LocalCellId = INTRAFREQHOGROUP.LocalCellId) AND (CELL.[NE NAME] = INTRAFREQHOGROUP.[NE NAME]))
	 INNER JOIN INTRARATHOCOMM ON CELL.[NE NAME] = INTRARATHOCOMM.[NE NAME])
	 INNER JOIN PDSCHCFG ON (CELL.LocalCellId = PDSCHCFG.LocalCellId) AND (CELL.[NE NAME] = PDSCHCFG.[NE NAME]))
	 INNER JOIN PHICHCFG ON (CELL.LocalCellId = PHICHCFG.LocalCellId) AND (CELL.[NE NAME] = PHICHCFG.[NE NAME])`

	qry.query4gulo = `SELECT DISTINCT 
  CELL.[NE NAME] as ne_name,
  ULOCELL.LOCELLTYPE as locelltype,
  ULOCELL.ULSPECTRUMSHARINGSWITCH as ulspectrumsharingswitch,
  ULOCELL.ULLEFTSHAREDBANDWIDTH as ulleftsharedbandwidth,
  ULOCELL.ULRIGHTSHAREDBANDWIDTH as ulrightsharedbandwidth,
  ULOCELL.DLLEFTSHAREDBANDWIDTH as dlleftsharedbandwidth,
  ULOCELL.DLRIGHTSHAREDBANDWIDTH as dlrightsharedbandwidth,
  ULOCELL.ULSPECTRUMSHARINGPHASE2 as ulspectrumsharingphase2,
  ULOCELLALGPARA.TURBOICENHANCEDSW as turboicenhancedsw,
  ULOCELLMACHSPARA.CQIADJALGOFNONCON as cqiadjalgofnoncon,
  ULOCELLMACHSPARA.CQIADJBYDYNBLERADJSW as cqiadjbydynbleradjsw,
  ULOCELL.ICMODE as icmode,
  ULOCELL.TURBOIC as turboic,
  ULOCELL.TURBOICPHASE2 as turboicphase2,
  NODEBALGPARA.CCPICPHASESWCTRL as ccpicphaseswctrl,
  ULOCELL.NBIS as nbis,
  ULOCELLALGPARA.DYNAMICPCPICHSWITCH as dynamicpcpichswitch
 FROM (CELL INNER JOIN ((ULOCELL INNER JOIN ULOCELLALGPARA ON (ULOCELL.ULOCELLID = ULOCELLALGPARA.ULOCELLID) AND (ULOCELL.[NE NAME] = ULOCELLALGPARA.[NE NAME])) INNER JOIN ULOCELLMACHSPARA ON ULOCELLALGPARA.[NE NAME] = ULOCELLMACHSPARA.[NE NAME]) ON CELL.[NE NAME] = ULOCELL.[NE NAME]) INNER JOIN NODEBALGPARA ON ULOCELL.[NE NAME] = NODEBALGPARA.[NE NAME];`

	if tech == "3g" {
		return qry.query3g, ""
	}

	if tech == "4g" {
		return qry.query4g, qry.query4gulo
	}

	return "", ""
}
