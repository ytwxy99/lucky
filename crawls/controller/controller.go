package controller

import (
	"github.com/gocolly/colly"
	"net/http"
)

func GetController() *colly.Collector {
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		// colly.AllowedDomains(webSites),
		// set User-Agent
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3100.0 Safari/537.36"),
	)

	// disable http KeepAlive attirbute.
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	return c
}
