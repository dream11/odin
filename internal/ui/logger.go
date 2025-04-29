package ui

import (
	"os"

	"github.com/dream11/odin/pkg/constant"
	log "github.com/sirupsen/logrus"
)

// CustomTextFormatter custom text formatter for logging
type CustomTextFormatter struct {
	BaseFormatter *log.TextFormatter
}

// Format Configure log format
func (f *CustomTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	var colorStart, colorEnd string
	colorEnd = "\033[0m"

	switch entry.Level {
	case log.WarnLevel:
		colorStart = "\033[33m" // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		colorStart = "\033[31m" // red
	case log.DebugLevel:
		colorStart = "\033[34m" // blue
	default:
		colorStart = "" // no color
		colorEnd = ""
	}

	return []byte(colorStart + entry.Message + colorEnd + "\n"), nil
}

func init() {

	log.SetFormatter(&CustomTextFormatter{
		BaseFormatter: &log.TextFormatter{
			ForceColors:            true,
			DisableColors:          false,
			DisableTimestamp:       true,
			DisableLevelTruncation: true,
		},
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
