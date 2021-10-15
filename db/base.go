package db

import (
	"github.com/ytwxy99/lucky/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var Log = utils.Log
var Database = GetDB()

func GetDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("lucky.db"), &gorm.Config{})
	if err != nil {
		panic("Open database error!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("connect db server failed.")
	}

	sqlDB.SetMaxIdleConns(10)                   // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)                  // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * 600) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

	return db
}
