package commands

import (
	"fmt"

	"github.com/brownhash/golog"
)

type Profile struct {
	Destroy bool
	Name    string
	Version string
	Env     string
}

func (n *Profile) Run(args []string) int {
	action := "deployed"
	if n.Destroy {
		action = "deleted"
	}
	if len(args) != 3 {
		golog.Error(fmt.Errorf("`env %s` requires exactly one argument `env name`, %d were given.", action, len(args)))
	}
	golog.Success(fmt.Sprintf("Profile/%s-%s-%s %s", args[0], args[1], args[2], action))
	return 0
}

func (n *Profile) Help() string {
	if n.Destroy {
		return "use `profile create <profile-name> <version> <env-name>` to deploy the provided profile in the provided env"
	}
	return "use `profile delete <profile-name> <version> <env-name>` to deploy the provided profile in the provided env"
}

func (n *Profile) Synopsis() string {
	if n.Destroy {
		return "delete a deployed profile in the provided env"
	}
	return "deploy the profile in the provided env"
}
