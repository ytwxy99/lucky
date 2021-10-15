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
	database, err := db.GetDB()
	if err != nil {
		Log.Error("Open database error:", err)
	}
	database.AutoMigrate(&db.History{})
	database.AutoMigrate(&db.Stocks{})
}

func InitDetailHistory() {
	// initialze detail history data.
	stocks := db.Stocks{}.FetchAll()
	for index, stock := range stocks {
		if index == 1 {
			histories, err := api.FetchHisotry(utils.Sohu, stock.TsCode)
			if err != nil {
				Log.Error("Init deital history error:", err)
			}

			for _, history := range histories {
				for _, hq := range history.Hq {
					_, err := db.History{}.FetchHistoryByTsCode(stock.TsCode, strings.Join(strings.Split(hq[0], "-"), ""))
					if err != nil {
						db.AddOne(db.History{
							TradeDate:   strings.Join(strings.Split(hq[0], "-"), ""),
							Open:        hq[1],
							Close:       hq[2],
							Change:      hq[3],
							PctChg:      hq[4],
							Low:         hq[5],
							High:        hq[6],
							Vol:         hq[7],
							Amount:      hq[8],
							VolumeRatio: hq[9],
							TsCode:      stock.TsCode,
						})
					}
				}
			}
		}
	}
}
