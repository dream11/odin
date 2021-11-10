package command

import (
	"github.com/dream11/odin/internal/command/commands"
	"github.com/mitchellh/cli"
)

/*
Command Structure:

	odin <verb> <resource> <options>

Verbs are essentially the actions that will be performed,
like: create, list, delete, etc...

Verb convention:
	- create
	- delete
	- describe
	- list
	- status
	- logs
	- deploy
	- destroy

Resources are the entities on with the verbs will run,
like: env, profile, etc...

Options are merely the flags that are required with the
command.
*/
func CommandCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		// Verbs for `env` resource
		"list env": func() (cli.Command, error) {
			return &commands.Env{List: true}, nil
		},
		"describe env": func() (cli.Command, error) {
			return &commands.Env{Describe: true}, nil
		},
		"create env": func() (cli.Command, error) {
			return &commands.Env{Create: true}, nil
		},
		"delete env": func() (cli.Command, error) {
			return &commands.Env{Delete: true}, nil
		},

		// Verbs for `profile` resource
		"create profile": func() (cli.Command, error) {
			return &commands.Profile{Create: true}, nil
		},
		"delete profile": func() (cli.Command, error) {
			return &commands.Profile{Delete: true}, nil
		},
		"list profile": func() (cli.Command, error) {
			return &commands.Profile{List: true}, nil
		},
		"describe profile": func() (cli.Command, error) {
			return &commands.Profile{Describe: true}, nil
		},
		"deploy profile": func() (cli.Command, error) {
			return &commands.Profile{Deploy: true}, nil
		},
		"destroy profile": func() (cli.Command, error) {
			return &commands.Profile{Destroy: true}, nil
		},

		// Sample commands
		"create test": func() (cli.Command, error) {
			return &commands.Test{Create: true}, nil
		},
		"delete test": func() (cli.Command, error) {
			return &commands.Test{Delete: true}, nil
		},
		"list test": func() (cli.Command, error) {
			return &commands.Test{List: true}, nil
		},
		"describe test": func() (cli.Command, error) {
			return &commands.Test{Describe: true}, nil
		},
		"status test": func() (cli.Command, error) {
			return &commands.Test{Status: true}, nil
		},
		"logs test": func() (cli.Command, error) {
			return &commands.Test{Logs: true}, nil
		},
		"deploy test": func() (cli.Command, error) {
			return &commands.Test{Deploy: true}, nil
		},
		"destroy test": func() (cli.Command, error) {
			return &commands.Test{Destroy: true}, nil
		},
	}
}
