package init

import (
	"github.com/ytwxy99/lucky/api"
	"github.com/ytwxy99/lucky/db"
	"github.com/ytwxy99/lucky/utils"
	"strings"
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
	for _, stock := range stocks {
		histories, err := api.FetchHisotry(utils.Sohu, stock.TsCode)
		if err != nil {
			Log.Error("Init deital history error:", err)
			return err
		}

		for _, history := range histories {
			for _, hq := range history.Hq {
				_, err := db.History{}.FetchHistoryByTsCode(stock.TsCode, strings.Join(strings.Split(hq[0], "-"), ""))
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
