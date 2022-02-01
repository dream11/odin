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
like: environment, profile, etc...

Options are merely the flags that are required with the
command.
*/

/*
TODO:
- status & logs verbs for env resource
- status & logs verbs for service resource
- add verbs for profile resource
*/

// CommandsCatalog : initiate commands catalog
func CommandsCatalog() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"configure": func() (cli.Command, error) {
			return &commands.Configure{}, nil
		},

		// Verbs for `env` resource
		"create env": func() (cli.Command, error) {
			return &commands.Env{Create: true}, nil
		},
		"update env": func() (cli.Command, error) {
			return &commands.Env{Update: true}, nil
		},
		"describe env": func() (cli.Command, error) {
			return &commands.Env{Describe: true}, nil
		},
		"list env": func() (cli.Command, error) {
			return &commands.Env{List: true}, nil
		},
		"delete env": func() (cli.Command, error) {
			return &commands.Env{Delete: true}, nil
		},
		"get-history env": func() (cli.Command, error) {
			return &commands.Env{GetHistory: true}, nil
		},
		"describe-history env": func() (cli.Command, error) {
			return &commands.Env{DescribeHistory: true}, nil
		},
		"status env": func() (cli.Command, error) {
			return &commands.Env{Status: true}, nil
		},

		// Verbs for `component-type` resource
		"list component-type": func() (cli.Command, error) {
			return &commands.ComponentType{List: true}, nil
		},

		// Verbs for `component` resource
		"describe component-type": func() (cli.Command, error) {
			return &commands.ComponentType{Describe: true}, nil
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
		"undeploy service": func() (cli.Command, error) {
			return &commands.Service{Undeploy: true}, nil
		},
		"status service": func() (cli.Command, error) {
			return &commands.Service{Status: true}, nil
		},
		/*
			Sample commands -

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

		*/
	}
}
