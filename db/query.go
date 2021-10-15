package db

import (
	"errors"
)

func (sks Stocks) FetchAll() []Stocks {
	var stocks []Stocks
	Database.Table("stocks").Find(&stocks)

	return stocks
}

func (history History) FetchHistoryByTsCode(tsCode string, tradeDate string) (*History, error) {
	var ht History
	Database.Table("history").
		Where("ts_code = ? AND trade_date = ?", tsCode, tradeDate).First(&ht)

	if ht.TsCode == "" {
		return &History{}, errors.New("record not found")
	}
	return &ht, nil
}
