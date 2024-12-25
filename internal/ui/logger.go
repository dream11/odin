package ui

import (
	"github.com/dream11/odin/pkg/constant"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05", // Custom format
		FullTimestamp:   true,
	})
	if os.Getenv(constant.LogLevelKey) == "yes" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
