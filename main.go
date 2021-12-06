package main

import (
	"os"

	odin "github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/cli"
	"github.com/dream11/odin/internal/ui"
)

var logger ui.Logger

func main() {
	c := cli.Cli(odin.App.Name, odin.App.Version)
	exitStatus, err := c.Run()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	os.Exit(exitStatus)
}

// TODO: https://github.com/mitchellh/go-glint
// TODO: https://github.com/charmbracelet/bubbletea for advanced interactions with user
