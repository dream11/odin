package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/dream11/odin/internal/commandline"
	"github.com/dream11/odin/pkg/shell"
)

type Env command

func (e *Env) Run(args []string) int {
	// Define flagset
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	name := flagSet.String("name", "demo", "name of environment")
	// positional parse flags from [3:]
	flagSet.Parse(os.Args[3:])

	if e.List {
		commandline.Interface.Info("Listing all envs")
		return shell.Exec("kubectl get ns")
	}

	if e.Describe {
		commandline.Interface.Info("Describing env:" + *name)
		return shell.Exec(fmt.Sprintf("kubectl describe ns %s", *name))
	}

	if e.Status {
		commandline.Interface.Info("Fetching status for:" + *name)
		return 0
	}

	if e.Logs {
		commandline.Interface.Info("Fetching execution logs for:" + *name)
		return 0
	}

	if e.Create {
		commandline.Interface.Warn("Creating env:" + *name)
		return shell.Exec(fmt.Sprintf("kubectl create ns %s", *name))
	}

	if e.Delete {
		commandline.Interface.Warn("Deleting env:" + *name)
		return shell.Exec(fmt.Sprintf("kubectl delete ns %s", *name))
	}

	commandline.Interface.Error("Not a valid command")
	return 1
}

func (e *Env) Help() string {
	if e.List {
		return commandHelper("list", "env", []string{})
	}
	if e.Describe {
		return commandHelper("describe", "env", []string{
			"--name=name of environment to describe",
		})
	}
	if e.Create {
		return commandHelper("create", "env", []string{
			"--name=name of environment to create",
		})
	}
	if e.Delete {
		return commandHelper("delete", "env", []string{
			"--name=name of environment to delete",
		})
	}

	return defaultHelper()
}

func (e *Env) Synopsis() string {
	if e.List {
		return "list all active envs"
	}
	if e.Describe {
		return "describe an env"
	}
	if e.Create {
		return "create an env"
	}
	if e.Delete {
		return "delete an env"
	}

	return defaultHelper()
}
