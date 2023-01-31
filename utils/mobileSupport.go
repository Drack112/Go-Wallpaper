package utils

import "fmt"

func VerifyMobileSupport(link string, mobileConfirm bool) string {
    if mobileConfirm {
        return fmt.Sprintf("%s&%s", link, "mobile=ok")
    } else {
        return link
    }
}
