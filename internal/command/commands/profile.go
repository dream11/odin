package commands

import (
	"fmt"

	"github.com/brownhash/golog"
)

type Profile struct {
	Deploy     bool
	Destroy    bool
}

func (n *Profile) Run(args []string) int {
	action := "" // initiate empty action
	if n.Deploy {
		action = "install"
	} else if n.Destroy {
		action = "uninstall"
	}

	if action == "" {
		if len(args) != 1 {
			golog.Error(fmt.Errorf("`profile` requires exactly one argument `profile name`, %d were given.", len(args)))
		}

		golog.Success("Listing all envs") // TODO: convert this log to debug type
		return 0
	}

	if len(args) != 3 {
		golog.Error(fmt.Errorf("`profile %s` requires exactly three arguments `profile name, version, env name`, %d were given.", action, len(args)))
	}

	golog.Success(fmt.Sprintf("Profile/%s@%s %sed in %s", args[0], args[1], action, args[2]))
	return 0
}

func (n *Profile) Help() string {
	if n.Deploy {
		return "use `profile deploy <profile-name> <version> <env-name>` to deploy the provided profile in the provided env"
	} else if n.Destroy {
		return "use `profile destroy <profile-name> <version> <env-name>` to destroy the provided profile in the provided env"
	}

	return "use `profile <name>` to list the created versions for the mentioned profile"
}

func (n *Profile) Synopsis() string {
	if n.Deploy {
		return "deploy the profile"
	} else if n.Destroy {
		return "destroy the deployed profile"
	}
	
	return "list profile versions"
}
