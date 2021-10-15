package main

import (
	"github.com/ytwxy99/lucky/utils"
	initialize "github.com/ytwxy99/lucky/init"
)

func main() {
	utils.Log.Info("start lucky process!")
	initialize.InitDB()
	initialize.InitDetailHistory()
}