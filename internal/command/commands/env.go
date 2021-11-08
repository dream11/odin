package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/dream11/odin/internal/ui"
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
		ui.Interface().Info("Listing all envs")
		return shell.Exec("kubectl get ns")
	}

	if e.Describe {
		ui.Interface().Info("Describing env:" + *name)
		return shell.Exec(fmt.Sprintf("kubectl describe ns %s", *name))
	}

	if e.Deploy {
		ui.Interface().Warn("Deploying env:" + *name)
		return shell.Exec(fmt.Sprintf("kubectl create ns %s", *name))
	}

	if e.Destroy {
		ui.Interface().Warn("Destroying env:" + *name)
		return shell.Exec(fmt.Sprintf("kubectl delete ns %s", *name))
	}

	ui.Interface().Error("Not a valid command")
	return 1
}

func (e *Env) Help() string {
	if e.List {
		return commandHelper("list", "env", []string{})
	}
	if e.Describe {
		return commandHelper("describe", "env", []string{
			"--name=name of environemnt to describe",
		})
	}
	if e.Deploy {
		return commandHelper("deploy", "env", []string{
			"--name=name of environemnt to deploy",
		})
	}
	if e.Destroy {
		return commandHelper("destroy", "env", []string{
			"--name=name of environemnt to destroy",
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
	if e.Deploy {
		return "deploy an env"
	}
	if e.Destroy {
		return "destroy an env"
	}

	return defaultHelper()
}
