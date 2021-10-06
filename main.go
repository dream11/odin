package main

import (
	"os"
	"log"

	"github.com/mitchellh/cli"
	"github.com/dream11/d11-cli/internal/command"
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
		log.Println(err)
	}

	os.Exit(exitStatus)
}