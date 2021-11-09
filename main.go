package main

import (
	"os"

	"github.com/dream11/odin/internal/cli"
	"github.com/dream11/odin/internal/commandline"
	"github.com/dream11/odin/odin"
)

func main() {
	c := cli.Cli(odin.App.Name, odin.App.Version)
	exitStatus, err := c.Run()
	if err != nil {
		commandline.Interface.Error(err.Error())
		os.Exit(1)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint
