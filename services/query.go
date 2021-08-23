package services

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aksafarand/dump/structs"
	"github.com/jmoiron/sqlx"
	"github.com/joho/sqltocsv"
)

func QueryExport3g(f string, src string, qry string) {

	db, err := sqlx.Open("odbc", f)
	if err != nil {
		log.Println("open db error ", err.Error())
		return
	}
	defer db.Close()

	timeStr := time.Now()
	log.Println("Start Query")
	rows, err := db.Queryx(qry)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Done Query In", fmt.Sprintf("%v", time.Since(timeStr)))
	csvFile, err := os.Create("./" + src + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	timeStr = time.Now()
	csvwriter := csv.NewWriter(csvFile)

	csvwriter.Write([]string{`cell_name`, `ne_name`, `cell_id`, `site_id`, `hspaplusswitch_fdpch_capability_invalid`, `loadbalanceratio`, `maxhsupa2msttiusernum`, `cellcoalgoenhswitch_spec_user_hsdpa_disabled_switch`, `earfcn`, `npriority`, `thdtohigh`, `thdtolow`, `emeasbw`, `eqrxlevmin`, `supcnopgrpindex`, `blacklstcellnumber`, `rsrqswitch`, `perfenhanceswitch3_perfenh_cqi_0_h2d_switch`})
	data := structs.Data3g{}

	for rows.Next() {
		err = rows.StructScan(&data)
		if err != nil {
			panic(err.Error())
		}
		out, err := json.Marshal(data)
		if err != nil {
			panic(err.Error())
		}

		dt := structs.Data3g{}
		_ = json.Unmarshal(out, &dt)

		var str []string
		str = append(str, dt.CellName.String, dt.NeName.String, dt.CellId.String, dt.SiteId.String, dt.HSPAPLUSSWITCH_FDPCH_CAPABILITY_INVALID.String, dt.LOADBALANCERATIO.String, dt.MAXHSUPA2MSTTIUSERNUM.String, dt.CELLCOALGOENHSWITCH_SPEC_USER_HSDPA_DISABLED_SWITCH.String, dt.EARFCN.String, dt.NPRIORITY.String, dt.THDTOHIGH.String, dt.THDTOLOW.String, dt.EMEASBW.String, dt.EQRXLEVMIN.String, dt.SUPCNOPGRPINDEX.String, dt.BLACKLSTCELLNUMBER.String, dt.RSRQSWITCH.String, dt.PERFENHANCESWITCH3_PERFENH_CQI_0_H2D_SWITCH.String)
		csvwriter.Write(str)

	}

	csvwriter.Flush()
	csvFile.Close()
	rows.Close()
	log.Println("Done Export In", fmt.Sprintf("%v", time.Since(timeStr)))

}

func QueryTables(f string, src string, tables []string) {
	db, err := sql.Open("odbc", f)
	if err != nil {
		log.Println("open db error ", err.Error())
		return
	}
	defer db.Close()

	for _, t := range tables {
		timeStr := time.Now()
		log.Println("Start Fetching", t)
		rows, err := db.Query(fmt.Sprintf("select * from [%s]", t))
		if err != nil {
			if strings.Contains(err.Error(), t) {
				log.Println("Tables", t, "Not Found In", src)
			}

			continue
		}
		err = sqltocsv.WriteFile(fmt.Sprintf("./%s_%s.csv", src, t), rows)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("Done Export In", fmt.Sprintf("%v", time.Since(timeStr)))
	}

}

func QueryExport4g(f string, src string, qry1 string, qry2 string) {

	db, err := sqlx.Open("odbc", f)
	if err != nil {
		log.Println("open db error ", err.Error())
		return
	}
	defer db.Close()

	timeStr := time.Now()
	log.Println("Start Query")
	rows, err := db.Queryx(qry1)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Done Query In", fmt.Sprintf("%v", time.Since(timeStr)))
	csvFile, err := os.Create("./" + src + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	timeStr = time.Now()
	csvwriter := csv.NewWriter(csvFile)

	csvwriter.Write([]string{`ne_name`, `dlbandwidth`, `ulbandwidth`, `pucchalgoswitch`, `dlpcalgoswitch`, `nonstandardbwalgosw`, `cqiadjalgoswitch`, `dlschswitch`, `ulschswitch`, `rbdamagenearpointiblerth`, `fdueenhapercqitrigperiod`, `blindhoa1a2thdrsrp`, `pdcchcapacityimproveswitch`, `pdcchpowerenhancedswitch`, `pdcchmaxcoderate`, `pdcchoutloopadjlowerlimit`, `pdcchsymnumswitch`, `cceratioadjswitch`, `measbandwidthcfgind`, `cellreselpriority`, `thrshservlowqcfgind`, `snonintrasearch`, `thrshservlow`, `qqualmin`, `pucchperiodiccqioptswitch`, `dlearfcn`, `measbandwidth`, `threshxlow`, `cellreselprioritycfgind`, `interfreqhoa1thdrsrp`, `interfreqhoa2thdrsrp`, `interfreqhoa4thdrsrp`, `interrathoa1a2trigquan`, `intrafreqhoa3hyst`, `intrafreqhoa3offset`, `intrafreqhoa3trigquan`, `interfreqhoa1a2trigquan`, `freqpriinterfreqhoa1trigquan`, `interfreqhoa4trigquan`, `a3interfreqhoa1a2trigquan`, `referencesignalpwr`, `phichduration`, `phichresource`, `format3rbnum`, `max2ccackchnum`, `prachfreqoffsetstrategy`, `spectrumcloudenhswitch`, `spectrumcloudswitch`, `dlcchsendstrategy`, `interfpfmoptswitch`, `ansrssimutrans`, `prbrsvtype`, `lteprbstartindex`, `lteprbendindex`})
	data := structs.Data4g{}

	for rows.Next() {
		err = rows.StructScan(&data)
		if err != nil {
			panic(err.Error())
		}
		out, err := json.Marshal(data)
		if err != nil {
			panic(err.Error())
		}

		dt := structs.Data4g{}
		_ = json.Unmarshal(out, &dt)
		var str []string
		str = append(str, dt.NE_NAME.String, dt.DLBANDWIDTH.String, dt.ULBANDWIDTH.String, dt.PUCCHALGOSWITCH.String, dt.DLPCALGOSWITCH.String, dt.NONSTANDARDBWALGOSW.String, dt.CQIADJALGOSWITCH.String, dt.DLSCHSWITCH.String, dt.ULSCHSWITCH.String, dt.RBDAMAGENEARPOINTIBLERTH.String, dt.FDUEENHAPERCQITRIGPERIOD.String, dt.BLINDHOA1A2THDRSRP.String, dt.PDCCHCAPACITYIMPROVESWITCH.String, dt.PDCCHPOWERENHANCEDSWITCH.String, dt.PDCCHMAXCODERATE.String, dt.PDCCHOUTLOOPADJLOWERLIMIT.String, dt.PDCCHSYMNUMSWITCH.String, dt.CCERATIOADJSWITCH.String, dt.MEASBANDWIDTHCFGIND.String, dt.CELLRESELPRIORITY.String, dt.THRSHSERVLOWQCFGIND.String, dt.SNONINTRASEARCH.String, dt.THRSHSERVLOW.String, dt.QQUALMIN.String, dt.PUCCHPERIODICCQIOPTSWITCH.String, dt.DLEARFCN.String, dt.MEASBANDWIDTH.String, dt.THRESHXLOW.String, dt.CELLRESELPRIORITYCFGIND.String, dt.INTERFREQHOA1THDRSRP.String, dt.INTERFREQHOA2THDRSRP.String, dt.INTERFREQHOA4THDRSRP.String, dt.INTERRATHOA1A2TRIGQUAN.String, dt.INTRAFREQHOA3HYST.String, dt.INTRAFREQHOA3OFFSET.String, dt.INTRAFREQHOA3TRIGQUAN.String, dt.INTERFREQHOA1A2TRIGQUAN.String, dt.FREQPRIINTERFREQHOA1TRIGQUAN.String, dt.INTERFREQHOA4TRIGQUAN.String, dt.A3INTERFREQHOA1A2TRIGQUAN.String, dt.REFERENCESIGNALPWR.String, dt.PHICHDURATION.String, dt.PHICHRESOURCE.String, dt.FORMAT3RBNUM.String, dt.MAX2CCACKCHNUM.String, dt.PRACHFREQOFFSETSTRATEGY.String, dt.SPECTRUMCLOUDENHSWITCH.String, dt.SPECTRUMCLOUDSWITCH.String, dt.DLCCHSENDSTRATEGY.String, dt.INTERFPFMOPTSWITCH.String, dt.ANSRSSIMUTRANS.String, dt.PRBRSVTYPE.String, dt.LTEPRBSTARTINDEX.String, dt.LTEPRBENDINDEX.String)
		csvwriter.Write(str)

	}

	csvwriter.Flush()
	csvFile.Close()
	rows.Close()
	log.Println("Done Export In", fmt.Sprintf("%v", time.Since(timeStr)))

	// ULO
	timeStr = time.Now()
	log.Println("Start Query - Ulo")
	rows, err = db.Queryx(qry2)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Done Query In", fmt.Sprintf("%v", time.Since(timeStr)))
	csvFile, err = os.Create("./" + src + "_ULO.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	timeStr = time.Now()
	csvwriter = csv.NewWriter(csvFile)

	csvwriter.Write([]string{`ne_name`, `locelltype`, `ulspectrumsharingswitch`, `ulleftsharedbandwidth`, `ulrightsharedbandwidth`, `dlleftsharedbandwidth`, `dlrightsharedbandwidth`, `ulspectrumsharingphase2`, `turboicenhancedsw`, `cqiadjalgofnoncon`, `cqiadjbydynbleradjsw`, `icmode`, `turboic`, `turboicphase2`, `ccpicphaseswctrl`, `nbis`, `dynamicpcpichswitch`})
	dataUlo := structs.Data4gUlo{}

	for rows.Next() {
		err = rows.StructScan(&dataUlo)
		if err != nil {
			panic(err.Error())
		}
		out, err := json.Marshal(dataUlo)
		if err != nil {
			panic(err.Error())
		}

		dtUlo := structs.Data4gUlo{}
		_ = json.Unmarshal(out, &dtUlo)
		var str []string
		str = append(str, dtUlo.NE_NAME.String, dtUlo.LOCELLTYPE.String, dtUlo.ULSPECTRUMSHARINGSWITCH.String, dtUlo.ULLEFTSHAREDBANDWIDTH.String, dtUlo.ULRIGHTSHAREDBANDWIDTH.String, dtUlo.DLLEFTSHAREDBANDWIDTH.String, dtUlo.DLRIGHTSHAREDBANDWIDTH.String, dtUlo.ULSPECTRUMSHARINGPHASE2.String, dtUlo.TURBOICENHANCEDSW.String, dtUlo.CQIADJALGOFNONCON.String, dtUlo.CQIADJBYDYNBLERADJSW.String, dtUlo.ICMODE.String, dtUlo.TURBOIC.String, dtUlo.TURBOICPHASE2.String, dtUlo.CCPICPHASESWCTRL.String, dtUlo.NBIS.String, dtUlo.DYNAMICPCPICHSWITCH.String)
		csvwriter.Write(str)

	}

	csvwriter.Flush()
	csvFile.Close()
	rows.Close()
	log.Println("Done Export In", fmt.Sprintf("%v", time.Since(timeStr)))

}

func QueryExport4gUlo(f string, src string, qry2 string) {

	db, err := sqlx.Open("odbc", f)
	if err != nil {
		log.Println("open db error ", err.Error())
		return
	}
	defer db.Close()

	timeStr := time.Now()
	log.Println("Start Query - Ulo")
	rows, err := db.Queryx(qry2)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Done Query In", fmt.Sprintf("%v", time.Since(timeStr)))
	csvFile, err := os.Create("./" + src + "_ULO.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	timeStr = time.Now()
	csvwriter := csv.NewWriter(csvFile)

	csvwriter.Write([]string{`ne_name`, `locelltype`, `ulspectrumsharingswitch`, `ulleftsharedbandwidth`, `ulrightsharedbandwidth`, `dlleftsharedbandwidth`, `dlrightsharedbandwidth`, `ulspectrumsharingphase2`, `turboicenhancedsw`, `cqiadjalgofnoncon`, `cqiadjbydynbleradjsw`, `icmode`, `turboic`, `turboicphase2`, `ccpicphaseswctrl`, `nbis`, `dynamicpcpichswitch`})
	dataUlo := structs.Data4gUlo{}

	for rows.Next() {
		err = rows.StructScan(&dataUlo)
		if err != nil {
			panic(err.Error())
		}
		out, err := json.Marshal(dataUlo)
		if err != nil {
			panic(err.Error())
		}

		dtUlo := structs.Data4gUlo{}
		_ = json.Unmarshal(out, &dtUlo)
		var str []string
		str = append(str, dtUlo.NE_NAME.String, dtUlo.LOCELLTYPE.String, dtUlo.ULSPECTRUMSHARINGSWITCH.String, dtUlo.ULLEFTSHAREDBANDWIDTH.String, dtUlo.ULRIGHTSHAREDBANDWIDTH.String, dtUlo.DLLEFTSHAREDBANDWIDTH.String, dtUlo.DLRIGHTSHAREDBANDWIDTH.String, dtUlo.ULSPECTRUMSHARINGPHASE2.String, dtUlo.TURBOICENHANCEDSW.String, dtUlo.CQIADJALGOFNONCON.String, dtUlo.CQIADJBYDYNBLERADJSW.String, dtUlo.ICMODE.String, dtUlo.TURBOIC.String, dtUlo.TURBOICPHASE2.String, dtUlo.CCPICPHASESWCTRL.String, dtUlo.NBIS.String, dtUlo.DYNAMICPCPICHSWITCH.String)
		csvwriter.Write(str)

	}

	csvwriter.Flush()
	csvFile.Close()
	rows.Close()
	log.Println("Done Export In", fmt.Sprintf("%v", time.Since(timeStr)))

}
