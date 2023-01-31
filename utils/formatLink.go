package utils

import (
    "fmt"
    "strings"
)

func FormatLink(category string) string {
    lowerCategory := strings.ToLower(category)
    formatCategory := strings.ReplaceAll(lowerCategory, " ", "%20")
    link := fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s", formatCategory)

    return link
}
