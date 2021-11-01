package main

import (
	"os"

	"github.com/brownhash/golog"
	"github.com/dream11/odin/internal/cli"
	"github.com/dream11/odin/odin"
)

func main() {
	c := cli.Cli(odin.App.Name, odin.App.Version)
	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint
