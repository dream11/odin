package cli

import (
	"os"

	"github.com/dream11/odin/internal/command"
	"github.com/mitchellh/cli"
)

// add commands to hide from help section
var hiddenCommands = []string{
	"create test",
	"delete test",
	"list test",
	"describe test",
	"status test",
	"logs test",
	"deploy test",
	"destroy test",
}

// Cli : initiate the cli framework
func Cli(appName, appVersion string) *cli.CLI {
	// initiate cli
	// for more refer https://github.com/mitchellh/cli/blob/master/cli.go#L49
	return &cli.CLI{
		Name:           appName,
		Version:        appVersion,
		Args:           os.Args[1:],
		Commands:       command.CommandsCatalog(),
		HelpFunc:       cli.BasicHelpFunc(appName),
		Autocomplete:   true,
		HelpWriter:     os.Stdout,
		ErrorWriter:    os.Stderr,
		HiddenCommands: hiddenCommands,
	}
}
