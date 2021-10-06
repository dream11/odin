package logger

import (
	"os"
	"strings"

	"github.com/brownhash/golog"
)

func HandleLogging() {
	setLogFormat := strings.ToLower(os.Getenv("D11_LOG_FORMAT"))
	if setLogFormat == "yes" {
		golog.SetLogFormat()
	} else {
		golog.UnsetLogFormat()
	}

	golog.SetLogLevel(os.Getenv("D11_LOG_LEVEL"))
}