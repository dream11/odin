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
	level, err := log.ParseLevel(os.Getenv(constant.LogLevelKey))
	if err != nil {
		log.Warning("Invalid log level. Allowed values are: panic, fatal, error, warn, info, debug, trace")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}
