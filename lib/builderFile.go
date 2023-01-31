package lib

import (
    "net/url"
    "strings"
)

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
