package d11cli

import (
	"github.com/dream11/d11-cli/internal/logger"
)

type application struct {
	Name       string
	Version    string
}

var App application = application{
	Name: "d11-cli",
	Version: "1.0.0-beta",
}

// initiate logger on app initialisation
func init() {
	logger.HandleLogging()
}