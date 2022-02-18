package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/dream11/odin/api/environment"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/datetime"
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
	env := flagSet.String("env-type", "dev", "environment to attach with environment")
	service := flagSet.String("service", "", "service name to filter out describe environment")
	component := flagSet.String("component", "", "component name to filter out describe environment")
	providerAccount := flagSet.String("account", "", "account name to provision the environment in")
	filePath := flagSet.String("file", "environment.yaml", "file to read environment config")
	id := flagSet.Int("id", 0, "unique id of a changelog of an env")

	err := flagSet.Parse(args)
	if err != nil {
		e.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if e.Create {
		emptyParameters := emptyParameters(map[string]string{"--env-type": *env})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Creating environment for team: " + *team)
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

		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))

		return 1
	}

	if e.Status {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Fetching status for environment: " + *name + ", service: " + *service)

			if *service != "" {

				envServiceStatus, err := envClient.EnvServiceStatus(*name, *service)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}

				relativeDeployedSinceTime := datetime.DateTimeFromNow(envServiceStatus.LastDeployedAt)
				e.Logger.Output("Service version: " + string(envServiceStatus.Version))
				e.Logger.Output("Last deployed: " + relativeDeployedSinceTime)
				e.Logger.Output("Component details: ")

				tableHeaders := []string{"Name", "Version", "Status"}
				var tableData [][]interface{}

				for _, component := range envServiceStatus.Components {
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

			} else {

				envStatus, err := envClient.EnvStatus(*name)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}

				tableHeaders := []string{"Name", "Version", "Status", "Last deployed"}
				var tableData [][]interface{}

				for _, serviceStatus := range envStatus.ServiceStatus {
					relativeDeployedSinceTime := datetime.DateTimeFromNow(serviceStatus.LastDeployedAt)
					tableData = append(tableData, []interface{}{
						serviceStatus.Name,
						serviceStatus.Version,
						serviceStatus.Status,
						relativeDeployedSinceTime,
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
		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if e.Update {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
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

		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if e.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
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

			e.Logger.Output(fmt.Sprintf("\n%s", details))
			if *service == "" && *component == "" {
				e.Logger.Output("\nCommand to descibe env")
				e.Logger.ItalicEmphasize(fmt.Sprintf("odin describe env --name %s --service <serviceName> --component <componentName>", *name))
			}
			return 0
		}
		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if e.List {
		e.Logger.Info("Listing all environment(s)")
		envList, err := envClient.ListEnv(*name, *team, *env, *providerAccount)
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name", "Team", "Env Type", "State", "Account", "Deletion Time", "Purpose", "CreatedAt", "CreatedBy", "UpdatedAt", "UpdatedBy"}
		var tableData [][]interface{}

		for _, env := range envList {
			relativeDeletionTimestamp := datetime.DateTimeFromNow(env.DeletionTime)
			relativeCreatedAtTimestamp := datetime.DateTimeFromNow(env.CreatedAt)
			relativeUpdatedAtTimestamp := datetime.DateTimeFromNow(env.UpdatedAt)
			tableData = append(tableData, []interface{}{
				env.Name,
				env.Team,
				env.EnvType,
				env.State,
				env.Account,
				relativeDeletionTimestamp,
				env.Purpose,
				relativeCreatedAtTimestamp,
				env.CreatedBy,
				relativeUpdatedAtTimestamp,
				env.UpdatedBy,
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
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Deleting environment: " + *name)
			envClient.DeleteEnv(*name)
			e.Logger.Success(fmt.Sprintf("Deletion started: %s", *name))
			return 0
		}

		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if e.GetHistory {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Fetching changelog for env: " + *name)
			envResp, err := envClient.GetHistoryEnv(*name)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"ID", "State", "Action", "Resource Details", "Modified by", "Last Modified"}
			var tableData [][]interface{}

			for _, env := range envResp {
				relativeCreationTimestamp := datetime.DateTimeFromNow(env.CreatedAt)
				tableData = append(tableData, []interface{}{
					env.ID,
					env.State,
					env.Action,
					env.ResourceDetails,
					env.CreatedBy,
					relativeCreationTimestamp,
				})
			}
			err = table.Write(tableHeaders, tableData)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			e.Logger.Output("\nCommand to describe a changelog in detail")
			e.Logger.ItalicEmphasize("odin describe-history env --name <envName> --id <changelogId>")
			return 0
		}
		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if e.DescribeHistory {
		s := ""
		if *id > 0 {
			s = strconv.Itoa(*id)
		}

		emptyParameters := emptyParameters(map[string]string{"--name": *name, "--id": s})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Detailed description of a changelog for env: " + *name + " with ID: " + s)
			envResp, err := envClient.DescribeHistoryEnv(*name, s)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			if len(envResp) == 0 {
				e.Logger.Output("\nCommand to get the correct ID of the changelog")
				e.Logger.ItalicEmphasize("odin get-history env --name " + *name)
				return 1
			}

			details, err := yaml.Marshal(envResp[0])
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}

			e.Logger.Output(string(details))

			return 0
		}
		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
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
			"--service service config that is deployed on env",
			"--component component config that is deployed on env",
		})
	}

	if e.Delete {
		return commandHelper("delete", "environment", []string{
			"--name=name of environment to delete",
		})
	}

	if e.GetHistory {
		return commandHelper("get-history", "environment", []string{
			"--name=name of environment fetch changelog for",
		})
	}

	if e.DescribeHistory {
		return commandHelper("describe-history", "environment", []string{
			"--name=name of environment to fetch changelog for",
			"--id=unique id of a changelog for the specified env to get details for (positive integer)",
		})
	}

	if e.Status {
		return commandHelper("status", "environment", []string{
			"--name=name of environment",
			"--service=name of service",
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

	if e.GetHistory {
		return "get changelog of an environment"
	}

	if e.DescribeHistory {
		return "get env config details for a changelog of an environment"
	}

	if e.Status {
		return "Fetch deployment status of the environment"
	}
	return defaultHelper()
}
