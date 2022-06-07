package lib

import (
    "fmt"

    "github.com/gocolly/colly"
    "github.com/sirupsen/logrus"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
    formatter := new(prefixed.TextFormatter)

    formatter.FullTimestamp = true
    formatter.SetColorScheme(&prefixed.ColorScheme{
        PrefixStyle:    "blue+b",
        TimestampStyle: "white+h",
    })
    formatter.ForceFormatting = true

    log.Level = logrus.DebugLevel
}

func GetRequest(link string) {
    c := colly.NewCollector()

    c.OnError(func(r *colly.Response, err error) {
        log.Fatal("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Print("\033[2J")
        log.Print("URL atual -> ", r.Request.URL.String(), "\n")
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
