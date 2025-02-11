package ui

import log "github.com/sirupsen/logrus"

type CustomTextFormatter struct {
	BaseFormatter *log.TextFormatter
}

func (f *CustomTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Return only the log message, ignoring the log level and timestamp
	return []byte(entry.Message + "\n"), nil
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

	level, err := log.ParseLevel("info")
	if err != nil {
		log.Warning("Invalid log level. Allowed values are: panic, fatal, error, warn, info, debug, trace")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}
