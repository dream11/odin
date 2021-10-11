package cli

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/dream11/d11-cli/internal/command"
)

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
		HiddenCommands: []string{"test"},
	}
}