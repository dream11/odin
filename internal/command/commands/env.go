package commands

import (
	"flag"

	"github.com/dream11/odin/api/environment"
	"github.com/dream11/odin/internal/backend"
)

var envClient backend.Env

// Env : command declaration
type Env command

// Run : implements the actual functionality of the command
func (e *Env) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	team := flagSet.String("team", "", "display environments created by a team")
	purpose := flagSet.String("purpose", "", "reason to create env")
	envType := flagSet.String("envType", "", "envType to attach with env")
	providerAccount := flagSet.String("account", "", "account name to provision the env in")

	err := flagSet.Parse(args)
	if err != nil {
		e.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if e.Create {
		if emptyParameterValidation([]string{*envType}) {
			e.Logger.Warn("Creating env in  " + *envType)

			envConfig := environment.Env{
				Team:    *team,
				Purpose: *purpose,
				EnvType: *envType,
				Account: *providerAccount,
			}

			response, err := envClient.CreateEnv(envConfig)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			e.Logger.Success("Env: " + response.Name + " creation is in progress!")

			return 0
		}

		e.Logger.Error("envType cannot be blank")
		return 1
	}

	e.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (e *Env) Help() string {
	if e.Create {
		return commandHelper("create", "env", []string{
			"--team=team name to associate the env with(optional)",
			"--purpose=reason to create env(optional)",
			"--envType=env_type to create env in",
			"--account=account name to provision the env in (optional)",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (e *Env) Synopsis() string {
	if e.Create {
		return "create an env"
	}
	return defaultHelper()
}
