package command

import (
	"github.com/mitchellh/cli"
	"github.com/dream11/d11-cli/internal/command/commands"
)

// TODO: Accept parsed flags
func CommandCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"env": func() (cli.Command, error) {
			return &commands.Namespace{Destroy: false}, nil
		},
		"env create": func() (cli.Command, error) {
			return &commands.Namespace{Destroy: false}, nil
		},
		"env delete": func() (cli.Command, error) {
			return &commands.Namespace{Destroy: true}, nil
		},
		// Sample command
		"test": func() (cli.Command, error) {
			return &commands.Test{}, nil
		},
	}
}
