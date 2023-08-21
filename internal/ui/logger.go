package ui

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
}
