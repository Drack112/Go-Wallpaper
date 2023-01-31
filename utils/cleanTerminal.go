package utils

import (
    "fmt"
    "os/exec"
)

func GetSystem(system string) string {

    command := ""

    switch system {
    case "windows":
        command = "cls"
    case "linux":
        command = "clear"
    case "darwin":
        command = "clear"
    }

    return command

}

func CleanTerminal(command string) string {
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
