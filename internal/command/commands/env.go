package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/brownhash/golog"
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
		golog.Success("Listing all envs")
		return shell.Exec("kubectl get ns")
	}

	if e.Describe {
		golog.Success(fmt.Sprintf("Describing env %s", *name))
		return shell.Exec(fmt.Sprintf("kubectl describe ns %s", *name))
	}

	if e.Deploy {
		golog.Warn(fmt.Sprintf("Deploying env %s", *name))
		return shell.Exec(fmt.Sprintf("kubectl create ns %s", *name))
	}

	if e.Destroy {
		golog.Warn(fmt.Sprintf("Destroying env %s", *name))
		return shell.Exec(fmt.Sprintf("kubectl delete ns %s", *name))
	}

	
	golog.Error("Not a valid command")

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