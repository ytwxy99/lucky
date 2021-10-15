package db

type Stocks struct {
	//gorm.Model

	StockId  string `gorm:"primary_key" json:"stock_id"`
	TsCode   string `json:"ts_code"`
	Mame     string `json:"name"`
	Classify string `json:"classify"`
	Region   string `json:"region"`
}

func (s Stocks) TableName() string {
	return "stocks"
}

type History struct {
	//gorm.Model

	TsCode        string `sql:"index"`
	TradeDate     string
	Open          string
	High          string
	Low           string
	Close         string
	PreClose      string
	Change        string // 涨跌额
	PctChg        string // 涨跌幅（未复权）
	Vol           string // 成交量
	Amount        string // 成交额
	TurnoverRate  string
	TurnoverRateF string // 流通换手率
	VolumeRatio   string // 量比
}

func (h History) TableName() string {
	return "history"
}
