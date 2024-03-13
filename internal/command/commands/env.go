package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/dream11/odin/api/environment"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/internal/constant"
	"github.com/dream11/odin/pkg/datetime"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/utils"
	"gopkg.in/yaml.v3"
)

// initiate backend client for environment
var envClient backend.Env

// Env : command declaration
type Env command

const ENV_NAME_KEY = "EnvName"

func splitProviderAccount(providerAccount string) []string {
	if providerAccount == "" {
		return nil
	}
	return strings.Split(providerAccount, ",")
}

// Run : implements the actual functionality of the command
func (e *Env) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	name := flagSet.String("name", "", "name of environment")
	team := flagSet.String("team", "", "display environments created by a team")
	env := flagSet.String("env-type", "", "environment to attach with environment")
	service := flagSet.String("service", "", "service name to filter out describe environment")
	component := flagSet.String("component", "", "component name to filter out describe environment")
	providerAccount := flagSet.String("account", "", "account name to provision the environment in")
	id := flagSet.Int("id", 0, "unique id of a changelog of an env")
	filePath := flagSet.String("file", "", "file to update env or provide options for environment operations")
	data := flagSet.String("data", "", "data for updating the env")
	displayAll := flagSet.Bool("all", false, "whether to display all environments")
	operation := flagSet.String("operation", "", "name of the operation to performed on the environment")
	options := flagSet.String("options", "", "options for environment operations")

	err := flagSet.Parse(args)
	if err != nil {
		e.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if e.Create {
		if *env == "" {
			*env = "dev"
		}
		emptyParameters := emptyParameters(map[string]string{"--env-type": *env, "--name": *name})
		if len(emptyParameters) == 0 {
			if len(*name) > 9 {
				e.Logger.Error("Env Name should not be of length more than 9")
				return 1
			}
			if (utils.SearchString(*name, "^[a-z]([a-z0-9-]*[a-z0-9])?$")) == "nil" {
				e.Logger.Error("Env name only allows lower case alphabets, numbers and '-'. Should start with an alphabet and not end with a hyphen.")
				return 1
			}
			envConfig := environment.Env{
				EnvType: *env,
				Account: splitProviderAccount(*providerAccount),
				Name:    *name,
			}

			envClient.CreateEnvStream(envConfig)
			e.Logger.Output("Command to set default Env. Once you do this, no need to pass --env everytime.")
			e.Logger.ItalicEmphasize("odin set env --name <envName>")
			return 0
		}

		e.Logger.Error(constant.ENV_NAME_OPTION)

		return 1
	}

	if e.Status {
		if *name == "" {
			*name = utils.FetchKey(ENV_NAME_KEY)
		}
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

				tableHeaders := []string{"Name", "Version", "Status", "Address"}
				var tableData [][]interface{}

				for _, component := range envServiceStatus.Components {
					tableData = append(tableData, []interface{}{
						component.Name,
						component.Version,
						component.Status,
						strings.Join(component.Address, ", "),
					})
				}

				table.Write(tableHeaders, tableData)

			} else {
				e.Logger.Info(fmt.Sprintf("Fetching status for environment: %s", *name))
				envDetail, err := envClient.DescribeEnv(*name, *service, *component)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}
				e.Logger.Output(fmt.Sprintf("Environment Status: %s\n", envDetail.State))
				tableHeaders := []string{"Name", "Version", "Status", "Last deployed"}
				var tableData []interface{}
				e.Logger.Output("Services:")
				colWidth := utils.GetColumnWidth(envDetail.Services)
				table.PrintHeader(tableHeaders, colWidth)

				c := make(chan serviceStatus)
				for _, service := range envDetail.Services {
					serviceName := service.Name
					go e.GetServiceStatus(*name, serviceName, c)
				}

				for range envDetail.Services {
					result := <-c
					if result.err != nil {
						e.Logger.Error(fmt.Sprintf("Error fetching status of service %s: %s", result.serviceName, result.err.Error()))
					} else {
						relativeDeployedSinceTime := datetime.DateTimeFromNow(result.response.LastDeployedAt)
						tableData = []interface{}{
							result.serviceName,
							result.response.Version,
							result.response.Status,
							relativeDeployedSinceTime,
						}
						table.AppendRow(tableData, colWidth)
					}
				}

				close(c)
			}

			return 0
		}
		e.Logger.Error(constant.ENV_NAME_OPTION)
		return 1
	}

	if e.Describe {
		if *name == "" {
			*name = utils.FetchKey(ENV_NAME_KEY)
		}
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			e.Logger.Info("Describing Env: " + *name)
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
		e.Logger.Error(constant.ENV_NAME_OPTION)
		return 1
	}

	if e.List {
		e.Logger.Info("Listing all environment(s)")
		envList, err := envClient.ListEnv(*name, *team, *env, *providerAccount, *displayAll)
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name", "Team", "Created By", "Env Type", "State", "Account"}
		var tableData [][]interface{}

		for _, env := range envList {
			tableData = append(tableData, []interface{}{
				env.Name,
				env.Team,
				env.CreatedBy,
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

			i, done := UserInput(name, e)
			if done {
				return i
			}

			e.Logger.Info(fmt.Sprintf("Initialising Env: [%s] deletion", *name))
			response, err := envClient.DeleteEnv(*name)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}
			e.Logger.Output(fmt.Sprintf("Deletion request accepted for Env [%s], You can track the progress here: %s", response.EnvResponse.Name, response.EnvResponse.ExecutorUrl))
			return 0
		}

		e.Logger.Error(constant.ENV_NAME_OPTION)
		return 1
	}

	if e.DescribeHistory {
		if *name == "" {
			*name = utils.FetchKey(ENV_NAME_KEY)
		}
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
		e.Logger.Error(constant.ENV_NAME_OPTION)
		return 1
	}

	if e.Set {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			e.Logger.Info(fmt.Sprintf("Setting default environment to %s", *name))
			err := utils.SetEnv(*name)
			if err != nil {
				e.Logger.Error(err.Error())
				return 1
			}
			e.Logger.Output(fmt.Sprintf("Default environment has been set to %s", *name))
			return 0
		}
		e.Logger.Error(constant.ENV_NAME_OPTION)
		return 1
	}

	if e.Update {
		if *name == "" {
			*name = utils.FetchKey(ENV_NAME_KEY)
		}

		isNamePresent := len(*name) > 0
		isDataPresent := len(*data) > 0
		isFilePresent := len(*filePath) > 0

		if !isNamePresent {
			e.Logger.Error("--name cannot be blank")
			return 1
		}

		if isDataPresent && isFilePresent {
			e.Logger.Error("You can provide either --data or --file but not both")
			return 1
		}

		if !isDataPresent && !isFilePresent {
			e.Logger.Error("You should provide either --data or --file")
			return 1
		}

		var updationData map[string]interface{}

		if isFilePresent {
			parsedConfig, err := parseFile(*filePath)
			if err != nil {
				e.Logger.Error("Error while parsing service file " + *filePath + " : " + err.Error())
				return 1
			}
			updationData = parsedConfig.(map[string]interface{})
		} else if isDataPresent {
			err = json.Unmarshal([]byte(*data), &updationData)
			if err != nil {
				e.Logger.Error("Unable to parse JSON data " + err.Error())
				return 1
			}
		}

		if len(updationData) == 0 {
			e.Logger.Error("You can't send an empty JSON data")
			return 1
		}

		e.Logger.Info("Updating Env: " + *name)

		envResp, err := envClient.UpdateEnv(*name, updationData)

		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		e.Logger.Output(fmt.Sprintf("The new autoDeletionTime is [%s]", envResp.DeletionTime))
		return 0
	}

	if e.Operate {
		if *name == "" {
			*name = utils.FetchKey(ENV_NAME_KEY)
		}

		isNamePresent := len(*name) > 0
		isOptionsPresent := len(*options) > 0
		isFilePresent := len(*filePath) > 0
		isOperationPresnt := len(*operation) > 0

		if !isNamePresent {
			e.Logger.Error("--name cannot be blank")
			return 1
		}
		if !isOperationPresnt {
			e.Logger.Error("--operation cannot be blank")
			return 1
		}
		if isOptionsPresent && isFilePresent {
			e.Logger.Error("You can provide either --options or --file but not both")
			return 1
		}

		var optionsData map[string]interface{}

		if isFilePresent {
			parsedConfig, err := parseFile(*filePath)
			if err != nil {
				e.Logger.Error("Error while parsing service file " + *filePath + " : " + err.Error())
				return 1
			}
			optionsData = parsedConfig.(map[string]interface{})
		} else if isOptionsPresent {
			err = json.Unmarshal([]byte(*options), &optionsData)
			if err != nil {
				e.Logger.Error("Unable to parse JSON data " + err.Error())
				return 1
			}
		}

		data := environment.OperationRequest{
			Operations: []environment.Operation{
				{
					Name: *operation,
					Data: optionsData,
				},
			},
		}

		e.Logger.Info("Validating the operation: " + *operation + " on the environment: " + *name)

		validateOperateResponse, err := envClient.ValidateOperation(*name, data)
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		for _, operation := range validateOperateResponse.Response.Operations {
			if operation.IsFeedbackRequired {
				consentMessage := fmt.Sprintf("\n%s", operation.Message)
				allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
				val, err := e.Input.AskWithConstraints(consentMessage, allowedInputs)
				if err != nil {
					e.Logger.Error(err.Error())
					return 1
				}
				if val != "Y" {
					e.Logger.Info("Aborting the operation")
					return 1
				}
			} else {
				e.Logger.Info("Validations succeeded. Proceeding...")
			}
		}

		operateResponse, err := envClient.OperateService(*name, data)
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		errored := false
		for _, operation := range operateResponse.Response.Operations {
			e.Logger.Output(operation.Message)
		}

		if errored {
			return 1
		}
		return 0
	}

	e.Logger.Error("Not a valid command")
	return 127
}

func UserInput(name *string, e *Env) (int, bool) {
	message := fmt.Sprintf("Are you sure you want to delete env: %s? [Y/n]: ", *name)
	allowedInputs := map[string]struct{}{"Y": {}, "n": {}}

	val, err := e.Input.AskWithConstraints(message, allowedInputs)

	if err != nil {
		e.Logger.Error(err.Error())
		return 1, true
	}

	if val != "Y" {
		e.Logger.Info("Skipping env deletion")
		return 0, true
	}
	return 0, false
}

// Help : returns an explanatory string
func (e *Env) Help() string {
	if e.Create {
		return commandHelper("create", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment"},
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
			{Flag: "--all", Description: "list all environments"},
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

	if e.Set {
		return commandHelper("set", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment"},
		})
	}

	if e.Update {
		return commandHelper("update", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment"},
			{Flag: "--data", Description: "JSON data which has values for the fields that should be updated in the env"},
			{Flag: "--file", Description: "JSON file which has values for the fields that should be updated in the env"},
		})
	}

	if e.Operate {
		return commandHelper("operate", "environment", "", []Options{
			{Flag: "--name", Description: "name of environment"},
			{Flag: "--operation", Description: "name of the operation to be performed on the environment"},
			{Flag: "--options", Description: "options for the operation in JSON format"},
			{Flag: "--file", Description: "path of the file which contains the options for the operation in JSON format"},
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

	if e.Set {
		return "Set a default env"
	}

	if e.Update {
		return "update an env"
	}

	if e.Operate {
		return "operate an environment"
	}

	return defaultHelper()
}

func (e *Env) GetServiceStatus(name string, serviceName string, c chan serviceStatus) {
	envServiceStatus, err := envClient.EnvServiceStatus(name, serviceName)
	c <- serviceStatus{serviceName, envServiceStatus, err}
}

type serviceStatus struct {
	serviceName string
	response    environment.EnvServiceStatus
	err         error
}
