package init

import (
	"fmt"
	"github.com/ytwxy99/lucky/db"
	"github.com/ytwxy99/lucky/utils"
	"github.com/ytwxy99/lucky/api"
)

var Log = utils.Log

func InitDB() {
	database, err := db.GetDB()
	if err != nil {
		Log.Error("Open database error:", err)
	}
	database.AutoMigrate(&db.History{})
}

func InitDetailHistory() {
	api.FetchHisotry(utils.Sohu, "000403.SZ")
}