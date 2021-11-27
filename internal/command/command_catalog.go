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

// CommandsCatalog : initiate commands catalog
func CommandsCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"configure": func() (cli.Command, error) {
			return &commands.Configure{}, nil
		},

		// Verbs for `env` resource
		"list env": func() (cli.Command, error) {
			return &commands.Env{List: true}, nil
		},

		// Verbs for `infra` resource
		"create infra": func() (cli.Command, error) {
			return &commands.Infra{Create: true}, nil
		},
		"describe infra": func() (cli.Command, error) {
			return &commands.Infra{Describe: true}, nil
		},
		"list infra": func() (cli.Command, error) {
			return &commands.Infra{List: true}, nil
		},
		"status infra": func() (cli.Command, error) {
			return &commands.Infra{Status: true}, nil
		},
		"logs infra": func() (cli.Command, error) {
			return &commands.Infra{Logs: true}, nil
		},
		"delete infra": func() (cli.Command, error) {
			return &commands.Infra{Delete: true}, nil
		},

		// Verbs for `component` resource
		"list component": func() (cli.Command, error) {
			return &commands.Component{List: true}, nil
		},

		// Verbs for `service` resource
		"create service": func() (cli.Command, error) {
			return &commands.Service{Create: true}, nil
		},
		"describe service": func() (cli.Command, error) {
			return &commands.Service{Describe: true}, nil
		},
		"list service": func() (cli.Command, error) {
			return &commands.Service{List: true}, nil
		},
		"label service": func() (cli.Command, error) {
			return &commands.Service{Label: true}, nil
		},
		"deploy service": func() (cli.Command, error) {
			return &commands.Service{Deploy: true}, nil
		},
		"destroy service": func() (cli.Command, error) {
			return &commands.Service{Destroy: true}, nil
		},
		"status service": func() (cli.Command, error) {
			return &commands.Service{Status: true}, nil
		},
		"logs service": func() (cli.Command, error) {
			return &commands.Service{Logs: true}, nil
		},
		"delete service": func() (cli.Command, error) {
			return &commands.Service{Delete: true}, nil
		},

		// Verbs for `profile` resource
		"create profile": func() (cli.Command, error) {
			return &commands.Profile{Create: true}, nil
		},
		"describe profile": func() (cli.Command, error) {
			return &commands.Profile{Describe: true}, nil
		},
		"list profile": func() (cli.Command, error) {
			return &commands.Profile{List: true}, nil
		},
		"deploy profile": func() (cli.Command, error) {
			return &commands.Profile{Deploy: true}, nil
		},
		"destroy profile": func() (cli.Command, error) {
			return &commands.Profile{Destroy: true}, nil
		},
		"status profile": func() (cli.Command, error) {
			return &commands.Profile{Status: true}, nil
		},
		"logs profile": func() (cli.Command, error) {
			return &commands.Profile{Logs: true}, nil
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
