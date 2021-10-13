package commands

import (
	"fmt"

	"github.com/brownhash/golog"
	"github.com/dream11/d11-cli/pkg/shell"
)

type Namespace struct {
<<<<<<< HEAD
	Destroy bool
=======
	Create     bool
	Destroy    bool
>>>>>>> a96fe98786c8cd2a762922dad8548c6b76336f9c
}

func (n *Namespace) Run(args []string) int {
	action := "" // initiate empty action
	if n.Create {
		action = "create"
	} else if n.Destroy {
		action = "delete"
	}

	if action == "" {
		if len(args) > 0 {
			golog.Error(fmt.Errorf("`env` requires no argument, %d were given.", len(args)))
		}

		golog.Debug("Listing all envs")
		return shell.Exec("kubectl get ns")
	}

	if len(args) > 1 {
		golog.Error(fmt.Errorf("`env %s` requires exactly one argument `env name`, %d were given.", action, len(args)))
	}

	command := fmt.Sprintf("kubectl %s ns %s", action, args[0])

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
<<<<<<< HEAD
	return "create env"
}
=======
	
	return "list envs"
}
>>>>>>> a96fe98786c8cd2a762922dad8548c6b76336f9c
