package main

import (
	"os"

	"github.com/dream11/d11-cli/d11cli"
	"github.com/dream11/d11-cli/internal/cli"
	"github.com/dream11/d11-cli/internal/logger"
	"github.com/brownhash/golog"
)

func main() {
	// handle logging format and levels
	logger.HandleLogging()
	// create/check workdir
	err := d11cli.WorkDir.Create()
	if err != nil {
		golog.Error(err)
	}

	c := cli.Cli(d11cli.App.Name, d11cli.App.Version)
	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint