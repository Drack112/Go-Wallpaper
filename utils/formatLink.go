package utils

import (
    "fmt"
    "strings"
)

func FormatLink(category string) string {
    formatCategory := strings.ReplaceAll(category, " ", "+")
    link := fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s", formatCategory)

    return link
}
