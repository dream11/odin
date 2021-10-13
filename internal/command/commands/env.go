package commands

import (
	"fmt"

	"github.com/brownhash/golog"
	"github.com/dream11/d11-cli/pkg/shell"
)

type Namespace struct {
	Destroy    bool
}

func (n *Namespace) Run(args []string) int {
	action := "create"
	if n.Destroy {
		action = "delete"
	}

	if len(args) > 1 {
		golog.Error(fmt.Errorf("`env %s` requires exactly one argument `env name`, %d were given.", action, len(args)))
	}

	command := fmt.Sprintf("kubectl %s ns %s", action, args[0])

	return shell.Exec(command)
}

func (n *Namespace) Help() string {
	if n.Destroy {
		return "use `env delete <env name>` to delete the provided env name"
	}
	return "use `env create <env name>` to create/delete the provided env name"
}

func (n *Namespace) Synopsis() string {
	if n.Destroy {
		return "delete env"
	}
	return "create env"
}