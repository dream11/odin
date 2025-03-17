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
	var logLevel string
	if value, ok := os.LookupEnv(constant.LogLevelKey); ok {
		logLevel = value
	} else {
		logLevel = "info"
	}
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Warning("Invalid log level. Allowed values are: panic, fatal, error, warn, info, debug, trace")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}
