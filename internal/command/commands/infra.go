package commands

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/dream11/odin/api/infra"
	"github.com/dream11/odin/internal/backend"
)

// initiate backend client for infra
var infraClient backend.Infra

// Infra : command declaration
type Infra command

// Run : implements the actual functionality of the command
func (i *Infra) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	name := flagSet.String("name", "", "name of environment")
	//detail := flagSet.Bool("detail", false, "detailed view of environments")
	//all := flagSet.Bool("all", false, "display all environments (active & inactive)")
	team := flagSet.String("team", "", "display environments created by a team")
	reason := flagSet.String("reason", "", "reason to create infra")
	env := flagSet.String("env", "", "env to attach with infra")
	state := flagSet.String("state", "", "state of infras to fetch")

	// positional parse flags from [3:]
	err := flagSet.Parse(os.Args[3:])
	if err != nil {
		i.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if i.Create {
		i.Logger.Warn("Creating infra: " + *name + " for team: " + *team)

		infraConfig := infra.Infra{
			Name:   *name,
			Team:   *team,
			Reason: *reason,
			Env:    *env,
		}

		infraConfigJson, err := json.Marshal(infraConfig)
		if err != nil {
			i.Logger.Error("Unable to generate infra config! " + err.Error())
			return 1
		}

		// TODO: validate request
		infraClient.CreateInfra(*name, infraConfigJson)

		return 0
	}

	if i.Describe {
		i.Logger.Info("Describing infra: " + *name)
		// TODO: validate request & receive parsed input to display
		infraClient.DescribeInfra(*name)

		return 0
	}

	if i.List {
		i.Logger.Info("Listing all infra(s)")
		// TODO: validate request & receive parsed input to display
		infraClient.ListInfra(*team, *state, *env)

		return 0
	}

	if i.Delete {
		i.Logger.Warn("Deleting infra:" + *name)
		// TODO: validate request
		infraClient.DeleteInfra(*name)

		return 0
	}

	if i.Status {
		i.Logger.Info("Fetching status for infra: " + *name)
		// TODO: call PG api that will fetch the status of the given infra
		// GET /infraStatus?name=<infra name>
		return 0
	}

	if i.Logs {
		i.Logger.Info("Fetching logs for infra: " + *name)
		// TODO: call PG api that will fetch the logs of the given infra
		// GET /infraLogs?name=<infra name>
		return 0
	}

	i.Logger.Error("Not a valid command")
	return 1
}

// Help : returns an explanatory string
func (i *Infra) Help() string {
	if i.List {
		return commandHelper("list", "infra", []string{
			"--team=team name",
			"--state=current state of infra",
			"--env=env of infra",
		})
	}

	if i.Describe {
		return commandHelper("describe", "infra", []string{
			"--name=name of infra to describe",
		})
	}

	if i.Status {
		return commandHelper("status", "infra", []string{
			"--name=name of infra to get status",
		})
	}

	if i.Logs {
		return commandHelper("logs", "infra", []string{
			"--name=name of infra to get execution logs",
		})
	}

	if i.Create {
		return commandHelper("create", "infra", []string{
			"--name=name of infra to create",
			"--team=team name to associate the infra with",
			"--reason=reason to create infra",
			"--env=env to create infra in",
		})
	}

	if i.Delete {
		return commandHelper("delete", "infra", []string{
			"--name=name of environment to delete",
		})
	}
	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (i *Infra) Synopsis() string {
	if i.List {
		return "list all active infra"
	}

	if i.Describe {
		return "describe an infra"
	}

	if i.Status {
		return "current status of an infra"
	}

	if i.Logs {
		return "execution logs for infra"
	}

	if i.Create {
		return "create an infra"
	}

	if i.Delete {
		return "delete an infra"
	}

	return defaultHelper()
}
