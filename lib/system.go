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
        panic(err.Error())
    }

    return string(out)
}
