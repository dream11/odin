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
	providerAccount := flagSet.String("account", "", "account name to provision the environment in")
	filePath := flagSet.String("file", "environment.yaml", "file to read environment config")
	detailed := flagSet.Bool("detailed", false, "get detailed view")
	serviceName := flagSet.String("servicename", "", "name of service")
	componentName := flagSet.String("componentname", "", "name of service")

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

	if e.Status{
		if emptyParameterValidation([]string{*name}) {
			e.Logger.Info("Fetching status for environment: " + *name+", service: "+ *serviceName+", component: "+ *componentName)

			if *componentName!="" && *serviceName==""{
				e.Logger.Error("serviceName cannot be blank when componentName is specified")
				return 1
			}
			envStatus, err := envClient.EnvStatus(*name, *serviceName, *componentName)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"Name", "Version", "Status"}
			var tableData [][]interface{}

			if *componentName!=""{
				
				e.Logger.Success(envStatus.Status)

			}else if *serviceName!=""{

				for _, component := range envStatus.Components {
					tableData = append(tableData, []interface{}{
						component.Name,
						component.Version,
						component.Status,
					})
				}

				err = table.Write(tableHeaders, tableData)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}

			}else{
				
				for _, service := range envStatus.Services {
					tableData = append(tableData, []interface{}{
						service.Name,
						service.Version,
						service.Status,
					})
				}

				err = table.Write(tableHeaders, tableData)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}
			}

			return 0
		}
		e.Logger.Error("Environment name cannot be blank")
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
			envResp, err := envClient.DescribeEnv(*name)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			for _, env := range envResp {
				e.Logger.Info(env.Name + " details!")
				details, err := yaml.Marshal(env)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}

				e.Logger.Output(string(details))
			}

			return 0
		}

		if e.List {
			e.Logger.Info("Listing all environment(s)")
			envList, err := envClient.ListEnv()
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			if *detailed {
				for _, env := range envList {
					e.Logger.Info("Env definition for: " + env.Name)

					envYaml, err := yaml.Marshal(env)
					if err != nil {
						e.Logger.Error("Unable to parse environment definition! " + err.Error())
						return 1
					}

					e.Logger.Output(string(envYaml))
				}
			} else {
				tableHeaders := []string{"Name", "Purpose", "Team", "Env Type", "State", "Account", "Deletion Time"}
				var tableData [][]interface{}

				for _, inf := range envList {
					tableData = append(tableData, []interface{}{
						inf.Name,
						inf.Purpose,
						inf.Team,
						inf.EnvType,
						inf.State,
						inf.Account,
						inf.DeletionTime,
					})
				}

				err = table.Write(tableHeaders, tableData)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}
			}
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
			"--detailed (get a detailed view)",
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

	if e.Status {
		return commandHelper("status", "environment", []string{
			"--name=name of environment",
			"--servicename=name of service",
			"--componentname=name of component",
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

	if e.Status {
		return "Fetch deployment status of the environment"
	}
	return defaultHelper()
}
