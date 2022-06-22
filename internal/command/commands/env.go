package commands

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/dream11/odin/api/environment"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/datetime"
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
	env := flagSet.String("env-type", "dev", "environment to attach with environment")
	service := flagSet.String("service", "", "service name to filter out describe environment")
	component := flagSet.String("component", "", "component name to filter out describe environment")
	providerAccount := flagSet.String("account", "", "account name to provision the environment in")
	id := flagSet.Int("id", 0, "unique id of a changelog of an env")

	err := flagSet.Parse(args)
	if err != nil {
		e.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if e.Create {
		emptyParameters := emptyParameters(map[string]string{"--env-type": *env})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Initiating environment creation")
			envConfig := environment.Env{
				EnvType: *env,
				Account: *providerAccount,
			}

			envClient.CreateEnvStream(envConfig)

			return 0
		}

		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))

		return 1
	}

	if e.Status {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {

			if *service != "" {
				e.Logger.Info(fmt.Sprintf("Fetching status for service: %s in environment: %s", *service, *name))
				envServiceStatus, err := envClient.EnvServiceStatus(*name, *service)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}

				relativeDeployedSinceTime := datetime.DateTimeFromNow(envServiceStatus.LastDeployedAt)
				e.Logger.Output("Service version: " + string(envServiceStatus.Version))
				e.Logger.Output("Service Status: " + string(envServiceStatus.Status))
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

				table.Write(tableHeaders, tableData)

			} else {
				e.Logger.Info(fmt.Sprintf("Fetching status for environment: %s", *name))
				envStatus, err := envClient.EnvStatus(*name)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}
				e.Logger.Output(fmt.Sprintf("Environment Status: %s\n", envStatus.Status))
				tableHeaders := []string{"Name", "Version", "Status", "Last deployed"}
				var tableData [][]interface{}
				e.Logger.Output("Services:")
				for _, serviceStatus := range envStatus.ServiceStatus {
					relativeDeployedSinceTime := datetime.DateTimeFromNow(serviceStatus.LastDeployedAt)
					tableData = append(tableData, []interface{}{
						serviceStatus.Name,
						serviceStatus.Version,
						serviceStatus.Status,
						relativeDeployedSinceTime,
					})
				}

				table.Write(tableHeaders, tableData)

			}

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
				e.Logger.Output("\nCommand to describe a service component deployed in an env")
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

		tableHeaders := []string{"Name", "Team", "Env Type", "State", "Account"}
		var tableData [][]interface{}

		for _, env := range envList {
			tableData = append(tableData, []interface{}{
				env.Name,
				env.Team,
				env.EnvType,
				env.State,
				env.Account,
			})
		}

		table.Write(tableHeaders, tableData)

		return 0
	}

	if e.Delete {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Environment(" + *name + ") deletion initiated")
			response, err := envClient.DeleteEnv(*name)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}
			e.Logger.Output(fmt.Sprintf("Deletion request accepted. Env [%s] will be deleted.", response.EnvResponse.Name))
			return 0
		}

		e.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if e.DescribeHistory {

		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			if *id == 0 {
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
				table.Write(tableHeaders, tableData)

				e.Logger.Output("\nCommand to describe a changelog in detail")
				e.Logger.ItalicEmphasize("odin history env --name <envName> --id <changelogId>")
				return 0

			} else {
				s := ""
				if *id > 0 {
					s = strconv.Itoa(*id)
				}
				e.Logger.Info("Detailed description of a changelog for env: " + *name + " with ID: " + s)
				envResp, err := envClient.DescribeHistoryEnv(*name, s)
				if err != nil {
					e.Logger.Error(err.Error())
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
		return commandHelper("create", "environment", "", []Options{
			{Flag: "--env-type", Description: "type of environment"},
			{Flag: "--account", Description: "account name to provision the environment in (optional)"},
		})
	}

	if e.List {
		return commandHelper("list", "environment", "", []Options{
			{Flag: "--name", Description: "name of env"},
			{Flag: "--team", Description: "name of team"},
			{Flag: "--env-type", Description: "env type of the environment"},
			{Flag: "--account", Description: "cloud provider account name"},
		})
	}

	if e.Describe {
		return commandHelper("describe", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment to describe"},
			{Flag: "--service", Description: "service config that is deployed on env"},
			{Flag: "--component", Description: "component config that is deployed on env"},
		})
	}

	if e.Delete {
		return commandHelper("delete", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment to delete"},
		})
	}

	if e.DescribeHistory {
		return commandHelper("history", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment to fetch changelog for"},
			{Flag: "--id", Description: "unique id of a changelog for the specified env to get details for (positive integer)"},
		})
	}

	if e.Status {
		return commandHelper("status", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment"},
			{Flag: "--service", Description: "name of service"},
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (e *Env) Synopsis() string {
	if e.Create {
		return "create an environment"
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

	if e.DescribeHistory {
		return "get env config details for a changelog of an environment"
	}

	if e.Status {
		return "Fetch deployment status of the environment"
	}
	return defaultHelper()
}
