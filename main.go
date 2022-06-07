package main

import (
    "fmt"
    "strings"

    "github.com/AlecAivazis/survey/v2"
    "github.com/Drack112/Go-Wallpaper/lib"
)

var category string

func init() {

    category = ""

    promptInput := &survey.Input{
        Message: "Please, type the category of the wallpaper to init the scrapper: ",
        Default: "Vaporwave",
        Help:    "Example: Vaporwave, Attack on titan, Initial D, Cyberpunk 2077",
    }

    survey.AskOne(promptInput, &category)

}

func main() {

    lowerCategory := strings.ToLower(category)
    removeSpaces := strings.ReplaceAll(lowerCategory, " ", "%20")

    link := fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s", removeSpaces)

    lib.GetRequest(link)

}
