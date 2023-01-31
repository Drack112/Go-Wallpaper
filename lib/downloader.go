package lib

import (
    "net/http"
    "os"

    "github.com/gocolly/colly"
)

var (
    fileName    string
    fullUrlFile string
)

func Downloader(link string) {

    fileName = ""
    fullUrlFile = ""

    c := colly.NewCollector()

    c.OnHTML("section", func(h *colly.HTMLElement) {
        h.ForEach("#show_img", func(i int, img *colly.HTMLElement) {
            url := img.Attr("src")

            fullUrlFile = url

            response, err := http.Get(fullUrlFile)
            if err != nil {
                panic(err)
            }
            defer response.Body.Close()

            buildFileName()

            file, err2 := os.Create(fileName) // "m1UIjW1.jpg"
            if err2 != nil {
                panic(err2)
            }
            defer file.Close()

            _, e := file.ReadFrom(response.Body)
            if e != nil {
                panic(e)
            }

        })
    })

    c.Visit(link + "/download")
}
