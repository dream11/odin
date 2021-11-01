package cli

import (
	"os"

	"github.com/dream11/odin/internal/command"
	"github.com/mitchellh/cli"
)

// add commands to hide from help section
var hidenCommands = []string{
	"create test",
	"delete test",
	"list test",
	"describe test",
	"deploy test",
	"destroy test",
}

func Cli(appName, appVersion string) *cli.CLI {
	// initiate cli
	// for more refer https://github.com/mitchellh/cli/blob/master/cli.go#L49
	return &cli.CLI{
		Name: appName,
		Version: appVersion,
		Args: os.Args[1:],
		Commands: command.CommandCatalog(),
		HelpFunc: cli.BasicHelpFunc(appName),
		Autocomplete: true,
		HelpWriter: os.Stdout,
		ErrorWriter: os.Stderr,
		HiddenCommands: hidenCommands,
	}
}