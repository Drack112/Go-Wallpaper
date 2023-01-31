package main

import (
    "fmt"
    "strings"

    "github.com/AlecAivazis/survey/v2"
    "github.com/Drack112/Go-Wallpaper/lib"
)

var (
    link     string
    category string
    confirm  bool
)

func main() {

    link = ""
    category = ""

    promptInput := &survey.Input{
        Message: "Por favor digite uma tag de wallpaper que deseja: ",
        Default: "Vaporwave",
        Help:    "Exemplos: Aesthetic, Attack on titan, Initial D, Cyberpunk 2077",
    }

    promptConfirm := &survey.Confirm{
        Message: "Você quer wallpapers para Mobile?",
        Default: false,
    }

    err := survey.AskOne(promptInput, &category)

    if err != nil {
        fmt.Print("Operação cancelada! Erro Abaixo: \n")
        panic(err)
    }

    err = survey.AskOne(promptConfirm, &confirm)

    if err != nil {
        fmt.Print("Operação cancelada! Erro Abaixo: \n")
        panic(err)
    }

    category = strings.ToLower(category)
    formatCategory := strings.ReplaceAll(category, " ", "%20")

    if confirm {
        link = fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s&mobile=ok", formatCategory)
    } else {
        link = fmt.Sprintf("https://www.wallpaperflare.com/search?wallpaper=%s", formatCategory)
    }

    lib.GetRequest(link)

}
