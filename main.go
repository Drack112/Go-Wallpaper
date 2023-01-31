package main

import (
    "fmt"

    "github.com/Drack112/Go-Wallpaper/cmd"
    "github.com/Drack112/Go-Wallpaper/lib"
    "github.com/Drack112/Go-Wallpaper/utils"
)

func main() {

    data, dataConfirm := cmd.ReceiveData()
    baseURI := utils.FormatLink(data)
    link := utils.VerifyMobileSupport(baseURI, dataConfirm)
    fmt.Println(link)

    lib.GetRequest(link)

}
