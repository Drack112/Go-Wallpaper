package lib

import (
    "runtime"

    "github.com/gocolly/colly"
    "github.com/sirupsen/logrus"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

var GOOS string

func init() {

    GOOS = runtime.GOOS

    formatter := new(prefixed.TextFormatter)

    formatter.FullTimestamp = true
    formatter.ForceFormatting = true

    log.Level = logrus.DebugLevel
}

func GetRequest(link string) {
    c := colly.NewCollector()

    c.OnError(func(r *colly.Response, err error) {
        log.Fatal("Request URL: ", r.Request.URL, "failed with response: ", r, "\nError: ", err)
    })

    c.OnResponse(func(r *colly.Response) {

        systemScrapper(GOOS)
        log.Print("Página dos wallpapers -> ", r.Request.URL.String(), "\n")
    })

    c.OnScraped(func(r *colly.Response) {
        log.Print("Scrapper has finished the job.")
    })

    c.OnHTML("li", func(e *colly.HTMLElement) {
        e.ForEach("a", func(i int, h *colly.HTMLElement) {
            links := h.Attr("href")
            Downloader(links)
        })
    })

    c.Visit(link)
}
