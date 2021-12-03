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
	- update
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

/*
TODO:
- add verbs for env resource
- status & logs verbs for infra resource
- status & logs verbs for service resource
- add verbs for profile resource
- add verbs for env resource
*/

// CommandsCatalog : initiate commands catalog
func CommandsCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"configure": func() (cli.Command, error) {
			return &commands.Configure{}, nil
		},

		// Verbs for `infra` resource
		"create infra": func() (cli.Command, error) {
			return &commands.Infra{Create: true}, nil
		},
		"update infra": func() (cli.Command, error) {
			return &commands.Infra{Update: true}, nil
		},
		"describe infra": func() (cli.Command, error) {
			return &commands.Infra{Describe: true}, nil
		},
		"list infra": func() (cli.Command, error) {
			return &commands.Infra{List: true}, nil
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
		"delete service": func() (cli.Command, error) {
			return &commands.Service{Delete: true}, nil
		},

		// Sample commands
		"create test": func() (cli.Command, error) {
			return &commands.Test{Create: true}, nil
		},
		"update test": func() (cli.Command, error) {
			return &commands.Test{Update: true}, nil
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
		"label test": func() (cli.Command, error) {
			return &commands.Test{Label: true}, nil
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
