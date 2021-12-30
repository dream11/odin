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
func (i *Env) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	team := flagSet.String("team", "", "display environments created by a team")
	purpose := flagSet.String("purpose", "", "reason to create env")
	envType := flagSet.String("envType", "", "envType to attach with env")
	providerAccount := flagSet.String("account", "", "account name to provision the env in")

	err := flagSet.Parse(args)
	if err != nil {
		i.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if i.Create {
		if emptyParameterValidation([]string{*envType}) {
			i.Logger.Warn("Creating env in  " + *envType)

			envConfig := environment.Env{
				Team:    *team,
				Purpose: *purpose,
				EnvType: *envType,
				Account: *providerAccount,
			}

			response, err := envClient.CreateEnv(envConfig)
			if err != nil {
				i.Logger.Error(err.Error())
				return 1
			}

			i.Logger.Success("Env: " + response.Name + " created!")

			return 0
		}

		i.Logger.Error("envType cannot be blank")
		return 1
	}

	i.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (i *Env) Help() string {
	if i.Create {
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
func (i *Env) Synopsis() string {
	if i.Create {
		return "create an env"
	}
	return defaultHelper()
}
