package odin

import (
	"github.com/dream11/odin/internal/logger"
)

type application struct {
	Name    string
	Version string
}

var App application = application{
	Name:    "odin",
	Version: "1.0.0-beta",
}

// initiate logger on app initialisation
func init() {
	logger.HandleLogging()
}
