package lib

import (
    "net/http"
    "net/url"
    "os"
    "strings"

    "github.com/gocolly/colly"
    "github.com/sirupsen/logrus"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
    fileName    string
    fullUrlFile string
)

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

func Downloader(link string) {
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

            n, e := file.ReadFrom(response.Body)
            if e != nil {
                panic(e)
            }
            log.Warn("File size: ", n)
        })
    })

    c.Visit(link + "/download")
}

func buildFileName() {
    fileUrl, err := url.Parse(fullUrlFile)
    if err != nil {
        panic(err)
    }

    path := fileUrl.Path
    segments := strings.Split(path, "/")

    fileName = segments[len(segments)-1]
    log.Print(fileName)
}
