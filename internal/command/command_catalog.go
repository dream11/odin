package command

import (
	"github.com/mitchellh/cli"
	"github.com/dream11/d11-cli/internal/command/commands"
)

// TODO: Accept parsed flags
func CommandCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"publish": func() (cli.Command, error) {
			return &commands.Publish{}, nil // TODO: Send flags here
		},
		// Sample command
		"test": func() (cli.Command, error) {
			return &commands.Test{}, nil
		},
	}
}
