package main

import (
	initialize "github.com/ytwxy99/lucky/init"
	"github.com/ytwxy99/lucky/utils"
)

var LOG = utils.Log

func main() {
	utils.Log.Info("start lucky process!")
	initialize.InitDB()
	err := initialize.InitDetailHistory()
	if err != nil {
		LOG.Error("initialize detail history data err:", err)
	}
}
