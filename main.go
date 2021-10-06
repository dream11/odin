package main

import (
	"os"

	"github.com/dream11/d11-cli/internal/command"
	"github.com/mitchellh/cli"
	"github.com/brownhash/golog"
)

const (
	appName = "d11-cli"
	appVersion = "1.0.0-beta"
)

func main() {
	c := cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]
	c.Commands = command.CommandCatalog()

	exitStatus, err := c.Run()
	if err != nil {
		golog.Error(err)
	}

	os.Exit(exitStatus)
}