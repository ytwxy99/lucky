package db

import (
	"errors"
	"github.com/ytwxy99/lucky/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Log = utils.Log

func GetDB() (db *gorm.DB, err error){
	db, err = gorm.Open(sqlite.Open("lucky.db"), &gorm.Config{})
	if err != nil {
		errors.New("Open database error!")
	}

	return db, err
}