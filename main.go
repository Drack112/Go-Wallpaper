package main

import (
    "fmt"
    "strings"

    "github.com/AlecAivazis/survey/v2"
    "github.com/Drack112/Go-Wallpaper/lib"
)

var (
    category string
    link     string
)

func init() {

    category = ""
    link = ""

    promptInput := &survey.Input{
        Message: "Please, type the category of the wallpaper to init the scrapper: ",
        Default: "Vaporwave",
        Help:    "Example: Aesthetic, Attack on titan, Initial D, Cyberpunk 2077",
    }

    err := survey.AskOne(promptInput, &category)

    if err != nil {
        panic(err.Error())
    }

}

func main() {

    category = strings.ToLower(category)
    formatCategory := strings.ReplaceAll(category, " ", "%20")

    link = fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s", formatCategory)

    lib.GetRequest(link)

}
