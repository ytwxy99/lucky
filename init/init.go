package init

import (
	"strings"

	"github.com/ytwxy99/lucky/api"
	"github.com/ytwxy99/lucky/db"
	"github.com/ytwxy99/lucky/utils"
)

var Log = utils.Log

func InitDB() {
	// init database
	db.Database.AutoMigrate(&db.History{})
	db.Database.AutoMigrate(&db.Stocks{})
}

func InitDetailHistory() error {
	// initialze detail history data.
	stocks := db.Stocks{}.FetchAll()
	for _, stockSplit := range transTsCodes(stocks) {
		histories, err := api.FetchHisotry(stockSplit)
		if err != nil {
			Log.Error("Get history date from public source error:", err, histories)
			return err
		}

		for _, history := range histories {
			for _, hq := range history.Hq {
				stock, err := db.Stocks{}.FetchStockByCode(strings.Split(history.Code, "_")[1])
				if err != nil {
					Log.Error("Get stock recode error:", err, history.Code)
				}
				_, err = db.History{}.FetchHistoryByTsCode(stock.TsCode, strings.Join(strings.Split(hq[0], "-"), ""))
				if err != nil {
					db.AddOne(db.History{
						TradeDate:   strings.Join(strings.Split(hq[0], "-"), ""),
						TsCode:      stock.TsCode,
						Open:        hq[1],
						Close:       hq[2],
						Change:      hq[3],
						PctChg:      hq[4],
						Low:         hq[5],
						High:        hq[6],
						Vol:         hq[7],
						Amount:      hq[8],
						VolumeRatio: hq[9],
					})
				}
			}
		}
	}
	return nil
}

func transTsCodes(stocks []db.Stocks) []string {
	totalLen := len(stocks)
	page := totalLen / utils.PageSize
	tsCodes := make([]string, page, page)

	for i := 0; i < page; i++ {
		tsCodesSplit := ""
		stockSplit := stocks[i*utils.PageSize : (i+1)*utils.PageSize]
		for index, stock := range stockSplit {
			if len(stockSplit) == (index + 1) {
				tsCodesSplit = tsCodesSplit + "cn_" + strings.Split(stock.TsCode, ".")[0]
			} else {
				tsCodesSplit = tsCodesSplit + "cn_" + strings.Split(stock.TsCode, ".")[0] + ","
			}
		}
		tsCodes[i] = tsCodesSplit
	}

	return tsCodes
}
