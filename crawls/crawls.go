package crawls

import (
	"github.com/ytwxy99/lucky/crawls/pages"
	"github.com/ytwxy99/lucky/db"
)

func DoCrawls() {
	stocks := db.Stocks{}.FetchAll()
	// data source from baidu inc.
	ba := pages.BaiduAladdin{
		UrlPrefix: "https://quote.eastmoney.com/",
		UrlSuffix: ".html?from=BaiduAladdin",
	}
	ba.Crawls(stocks)
}
