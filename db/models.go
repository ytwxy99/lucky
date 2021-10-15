package db

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model

	TsCode string
	TradeDate string
	Open string
	High string
	Low string
	Close string
	PreClose string
	Change string // 涨跌额
	PctChg string // 涨跌幅（未复权）
	Vol string // 成交量
	Amount string // 成交额
	TurnoverRate string
	TurnoverRateF string // 流通换手率
	VolumeRatio string // 量比
}
