package lib

import (
    "github.com/sirupsen/logrus"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func CreateLogger() {

    formatter := new(prefixed.TextFormatter)

    formatter.FullTimestamp = true
    formatter.ForceFormatting = true

    log.Level = logrus.DebugLevel
}
