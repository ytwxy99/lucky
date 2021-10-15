package db

import (
	"errors"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"github.com/ytwxy99/lucky/utils"
)

var Log = utils.Log

func GetDB() (db *gorm.DB, err error){
	path, _ := os.Getwd()
	db, err = gorm.Open(sqlite.Open(path + "test.db"), &gorm.Config{})
	if err != nil {
		errors.New("Open database error!")
	}

	return db, err
}
