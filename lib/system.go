package lib

import (
    "fmt"

    "github.com/Drack112/Go-Wallpaper/utils"
)

func SystemCleanTerminal(system string) {
    commandOS := utils.GetSystem(system)
    fmt.Println(utils.CleanTerminal(commandOS))
}
