package lib

import (
    "fmt"
    "os/exec"
)

var (
    system  string
    command string
)

func init() {
    system = ""
    command = ""
}

func systemScrapper(system string) string {

    switch system {
    case "windows":
        system = "windows"
        command = "cls"
    case "linux":
        system = "linux"
        command = "clear"
    case "darwin":
        system = "darwin"
        command = "clear"
    }

    fmt.Print(cleanTerminal(command))

    return system

}

func cleanTerminal(command string) string {
    out, err := exec.Command(command).Output()

    if err != nil {
        fmt.Print("Operação cancelada! Aparentemente, o comando pata limpar o terminal deu problema.\n")
        fmt.Print("Reveja se sua env GOOS está correta ao seu sistema operacional.\n")
        fmt.Print("Se o bug continuar abra uma issue em: https://github.com/Drack112/Go-Wallpaper/issues\n")
        fmt.Print(" ")

        fmt.Print("Log Error: \n")
        panic(err)
    }

    return string(out)
}
