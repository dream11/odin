package main

import (
	"os"

	"github.com/dream11/d11-cli/d11cli"
	"github.com/dream11/d11-cli/internal/cli"
	"github.com/brownhash/golog"
)

func main() {
	c := cli.Cli(d11cli.App.Name, d11cli.App.Version)
	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint