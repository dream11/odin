package command

import (
	"github.com/dream11/d11-cli/internal/command/commands"
	"github.com/mitchellh/cli"
)

// TODO: Accept parsed flags
func CommandCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		// Env Creation/Deletion
		"env": func() (cli.Command, error) {
			return &commands.Namespace{Create: false, Destroy: false}, nil
		},
		"env create": func() (cli.Command, error) {
			return &commands.Namespace{Create: true, Destroy: false}, nil
		},
		"env delete": func() (cli.Command, error) {
			return &commands.Namespace{Create: false, Destroy: true}, nil
		},
		// Profile Deploy/Destroy
		"profile": func() (cli.Command, error) {
			return &commands.Profile{Deploy: false, Destroy: false}, nil
		},
		"profile deploy": func() (cli.Command, error) {
			return &commands.Profile{Deploy: true, Destroy: false}, nil
		},
		"profile destroy": func() (cli.Command, error) {
			return &commands.Profile{Deploy: false, Destroy: true}, nil
		},
		// Sample command
		"test": func() (cli.Command, error) {
			return &commands.Test{}, nil
		},
	}
}
