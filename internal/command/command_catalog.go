package command

import (
	"github.com/mitchellh/cli"
	"github.com/dream11/d11-cli/internal/command/commands"
)

func CommandCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"test": func() (cli.Command, error) {
			return &commands.Test{}, nil
		},
	}
}
