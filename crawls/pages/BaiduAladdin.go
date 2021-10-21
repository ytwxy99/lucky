package pages

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/ytwxy99/lucky/crawls/controller"
	"github.com/ytwxy99/lucky/db"
	"github.com/ytwxy99/lucky/utils"
)

var Log = utils.Log

type BaiduAladdin struct {
	UrlPrefix string
	UrlSuffix string
}

func (ba *BaiduAladdin) Crawls(stocks []db.Stocks) {
	for index, stock := range stocks {
		code := strings.Split(stock.TsCode, ".")
		crawlUrl := ba.UrlPrefix + code[1] + code[0] + ba.UrlSuffix
		if index == 0 {
			c := controller.GetController()

			// Set error handler
			c.OnError(func(r *colly.Response, err error) {
				Log.Info("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
			})

			c.OnHTML("*", func(e *colly.HTMLElement) {
				return
			})

			c.OnRequest(func(r *colly.Request) {
				Log.Info("Visiting", r.URL.String())
			})

			// Start scraping
			err := c.Visit(crawlUrl)
			if err != nil {
				Log.Error("crawl page failed: ", crawlUrl)
			}

			return
		}
	}
}
