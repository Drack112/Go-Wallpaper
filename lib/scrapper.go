package lib

import (
    "runtime"

    "github.com/gocolly/colly"
)

func GetRequest(link string) {
    c := colly.NewCollector()

    c.OnError(func(r *colly.Response, err error) {
        log.Fatal("URL: ", r.Request.URL, "falhou com a response: ", r, "\nError: ", err)
    })

    c.OnResponse(func(r *colly.Response) {

        SystemCleanTerminal(runtime.GOOS)
        log.Print("Página dos wallpapers -> ", r.Request.URL.String(), "\n")
    })

    c.OnScraped(func(r *colly.Response) {
        log.Print("Scrapper finalizado.")
    })

    c.OnHTML("li", func(e *colly.HTMLElement) {
        e.ForEach("a", func(i int, h *colly.HTMLElement) {
            links := h.Attr("href")
            Downloader(links)
        })
    })

    c.Visit(link)
}
