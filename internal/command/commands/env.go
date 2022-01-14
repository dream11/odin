package commands

import (
	"encoding/json"
	"flag"
	"strings"

	"github.com/dream11/odin/api/environment"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"gopkg.in/yaml.v3"
)

// initiate backend client for environment
var envClient backend.Env

// Env : command declaration
type Env command

// Run : implements the actual functionality of the command
func (e *Env) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	name := flagSet.String("name", "", "name of environment")
	team := flagSet.String("team", "", "display environments created by a team")
	purpose := flagSet.String("purpose", "", "reason to create environment")
	env := flagSet.String("env-type", "kube", "environment to attach with environment")
	service := flagSet.String("service", "", "service name to filter out describe infra")
	component := flagSet.String("component", "", "component name to filter out describe infra")
	providerAccount := flagSet.String("account", "", "account name to provision the environment in")
	filePath := flagSet.String("file", "environment.yaml", "file to read environment config")

	err := flagSet.Parse(args)
	if err != nil {
		e.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if e.Create {
		if emptyParameterValidation([]string{*env}) {
			e.Logger.Warn("Creating environment for team: " + *team)
			envConfig := environment.Env{
				Team:    *team,
				Purpose: *purpose,
				EnvType: *env,
				Account: *providerAccount,
			}

			response, err := envClient.CreateEnv(envConfig)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			e.Logger.Success("Env: " + response.Name + " created!")

			return 0
		}

		e.Logger.Error("env-type is a required parameter")
		return 1
	}

	if e.Update {
		if emptyParameterValidation([]string{*name}) {
			e.Logger.Warn("Updating environment: " + *name)

			configData, err := file.Read(*filePath)
			if err != nil {
				e.Logger.Error("Unable to read from " + *filePath + "\n" + err.Error())
				return 1
			}

			var parsedConfig interface{}

			if strings.Contains(*filePath, ".yaml") || strings.Contains(*filePath, ".yml") {
				err = yaml.Unmarshal(configData, &parsedConfig)
				if err != nil {
					e.Logger.Error("Unable to parse YAML. " + err.Error())
					return 1
				}
			} else if strings.Contains(*filePath, ".json") {
				err = json.Unmarshal(configData, &parsedConfig)
				if err != nil {
					e.Logger.Error("Unable to parse JSON. " + err.Error())
					return 1
				}
			} else {
				e.Logger.Error("Unrecognized file format")
				return 1
			}

			envClient.UpdateEnv(*name, parsedConfig)

			return 0
		}

		e.Logger.Error("environment name cannot be blank")
		return 1
	}

	if e.Describe {
		if emptyParameterValidation([]string{*name}) {
			e.Logger.Info("Describing " + *name)
			envResp, err := envClient.DescribeEnv(*name, *service, *component)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			details, err := yaml.Marshal(envResp)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			e.Logger.Output(string(details))
			if *service == "" && *component == "" {
				e.Logger.Info("NEXT USEFUL COMMAND:- odin describe env --name " + *name + " --service <serviceName> --component <componentName>")
			}
			return 0
		}
		e.Logger.Error("name is a required parameter")
		return 1
	}

	if e.List {
		e.Logger.Info("Listing all environment(s)")
		envList, err := envClient.ListEnv(*name, *team, *env, *providerAccount)
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name", "Team", "Env Type", "State", "Account", "Deletion Time", "Purpose"}
		var tableData [][]interface{}

		for _, inf := range envList {
			tableData = append(tableData, []interface{}{
				inf.Name,
				inf.Team,
				inf.EnvType,
				inf.State,
				inf.Account,
				inf.DeletionTime,
				inf.Purpose,
			})
		}

		err = table.Write(tableHeaders, tableData)
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}
		return 0
	}

	if e.Delete {
		if emptyParameterValidation([]string{*name}) {
			e.Logger.Warn("Deleting environment:" + *name)
			envClient.DeleteEnv(*name)

			return 0
		}

		e.Logger.Error("environment name cannot be blank")
		return 1
	}

	e.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (e *Env) Help() string {
	if e.Create {
		return commandHelper("create", "environment", []string{
			"--team=team name to associate the environment with",
			"--purpose=reason to create environment",
			"--env-type=type of environment",
			"--account=account name to provision the environment in (optional)",
		})
	}

	if e.Update {
		return commandHelper("update", "environment", []string{
			"--name=name of environment to update",
			"--file=file path to pick update config",
		})
	}

	if e.List {
		return commandHelper("list", "environment", []string{
			"--name=name of env",
			"--team=name of team",
			"--env-type=env type of the environment",
			"--account=cloud provider account name",
		})
	}

	if e.Describe {
		return commandHelper("describe", "environment", []string{
			"--name=name of environment to describe",
		})
	}

	if e.Delete {
		return commandHelper("delete", "environment", []string{
			"--name=name of environment to delete",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (e *Env) Synopsis() string {
	if e.Create {
		return "create an environment"
	}

	if e.Update {
		return "update an environment"
	}

	if e.List {
		return "list all active environment"
	}

	if e.Describe {
		return "describe an environment"
	}

	if e.Delete {
		return "delete an environment"
	}
	return defaultHelper()
}
