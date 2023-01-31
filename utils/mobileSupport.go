package utils

import "fmt"

func VerifyMobileSupport(link string, mobileConfirm bool) string {
    if mobileConfirm {
        return fmt.Sprintf(link+"%s", "&mobile=ok")
    } else {
        return link
    }
}
