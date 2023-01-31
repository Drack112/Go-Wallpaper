package cmd

import (
    "fmt"

    "github.com/AlecAivazis/survey/v2"
)

var (
    category string
    confirm  bool
)

func InitPrompts() (*survey.Input, *survey.Confirm) {
    promptInput := &survey.Input{
        Message: "Por favor digite uma tag de wallpaper que deseja: ",
        Default: "Vaporwave",
        Help:    "Exemplos: Aesthetic, Attack on titan, Initial D, Cyberpunk 2077",
    }

    promptConfirm := &survey.Confirm{
        Message: "Você quer wallpapers para Mobile?",
        Default: false,
    }

    return promptInput, promptConfirm
}

func ReceiveData() (string, bool) {

    dataInput, dataConfirm := InitPrompts()

    err := survey.AskOne(dataInput, &category)

    if err != nil {
        fmt.Print("Operação cancelada! Erro Abaixo: \n")
        panic(err)
    }

    err = survey.AskOne(dataConfirm, &confirm)

    if err != nil {
        fmt.Print("Operação cancelada! Erro Abaixo: \n")
        panic(err)
    }

    return category, confirm
}
