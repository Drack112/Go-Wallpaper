package main

import (
    "fmt"
    "strings"

    "github.com/AlecAivazis/survey/v2"
    "github.com/Drack112/Go-Wallpaper/lib"
)

var (
    link string
)

func main() {

    link = ""
    category := ""

    promptInput := &survey.Input{
        Message: "Please, type the category of the wallpaper to init the scrapper: ",
        Default: "Vaporwave",
        Help:    "Example: Aesthetic, Attack on titan, Initial D, Cyberpunk 2077",
    }

    err := survey.AskOne(promptInput, &category)

    if err != nil {
        fmt.Print("Operação cancelada! Erro Abaixo: \n")
        panic(err)
    }

    category = strings.ToLower(category)
    formatCategory := strings.ReplaceAll(category, " ", "%20")

    link = fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s", formatCategory)

    lib.GetRequest(link)

}
