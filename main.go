package main

import (
	"os"

	"github.com/dream11/d11-cli/internal/cli"
	"github.com/dream11/d11-cli/internal/logger"
	"github.com/brownhash/golog"
)

const (
	appName = "d11-cli"
	appVersion = "1.0.0-beta"
)

func main() {
	// handle logging format and levels
	logger.HandleLogging()

	c := cli.Cli(appName, appVersion)
	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint