package commands

import (
	"os"
	"fmt"
	"flag"

	"github.com/brownhash/golog"
	"github.com/dream11/d11-cli/pkg/shell"
)

type Namespace struct {
	Create     bool
	Destroy    bool
}

func (n *Namespace) Run(args []string) int {
	action := "" // initiate empty action
	if n.Create {
		action = "create"
	} else if n.Destroy {
		action = "delete"
	}

	// Define flagset
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)

	// create flags
	name := flagSet.String("name", "demo", "name of environment")

	if action == "" {
		// positional parse flags from [2:] due to no subcommands
		flagSet.Parse(os.Args[2:])

		if flagSet.NFlag() != 0 {
			golog.Error(fmt.Errorf("`env` requires no flag, %d were given.", flagSet.NFlag()))
		}

		golog.Debug("Listing all envs")
		return shell.Exec("kubectl get ns")
	}

	if len(args) > 1 {
		golog.Error(fmt.Errorf("`env %s` requires exactly one argument `env name`, %d were given.", action, len(args)))
	}

	// positional parse flags from [3:] due to subcommands
	flagSet.Parse(os.Args[3:])

	if flagSet.NFlag() != 1 {
		golog.Error(fmt.Errorf("`env %s` requires exactly one flag `-name=string`, %d were given.", action, flagSet.NFlag()))
	}

	command := fmt.Sprintf("kubectl %s ns %s", action, *name)

	return shell.Exec(command)
}

func (n *Namespace) Help() string {
	if n.Create {
		return "use `env create <env name>` to create/delete the provided env name"
	} else if n.Destroy {
		return "use `env delete <env name>` to delete the provided env name"
	}

	return "use `env` to list all the created envs"
}

func (n *Namespace) Synopsis() string {
	if n.Create {
		return "create env"
	} else if n.Destroy {
		return "delete env"
	}
	
	return "list envs"
}