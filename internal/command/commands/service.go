package commands

import (
	"encoding/json"
	"flag"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
)

// initiate backend client for service
var serviceClient backend.Service

// Service : command declaration
type Service command

// Run : implements the actual functionality of the command
func (s *Service) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	filePath := flagSet.String("file", "service.yaml", "file to read service config")
	serviceName := flagSet.String("name", "", "name of service to be used")
	serviceVersion := flagSet.String("version", "", "version of service to be used")
	envName := flagSet.String("env", "", "name of environment to use")
	infraName := flagSet.String("infra", "", "name of infra to deploy the service in")
	teamName := flagSet.String("team", "", "name of user's team")
	isMature := flagSet.Bool("mature", false, "mark service version as matured")

	// positional parse flags from [3:]
	err := flagSet.Parse(os.Args[3:])
	if err != nil {
		s.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if s.Create {
		configData, err := file.Read(*filePath)
		if err != nil {
			s.Logger.Error("Unable to read from " + *filePath + "\n" + err.Error())
			return 1
		}

		var parsedConfig interface{}

		if strings.Contains(*filePath, ".yaml") || strings.Contains(*filePath, ".yml") {
			err = yaml.Unmarshal(configData, &parsedConfig)
			if err != nil {
				s.Logger.Error("Unable to parse YAML. " + err.Error())
				return 1
			}
		} else if strings.Contains(*filePath, ".json") {
			err = json.Unmarshal(configData, &parsedConfig)
			if err != nil {
				s.Logger.Error("Unable to parse JSON. " + err.Error())
				return 1
			}
		} else {
			s.Logger.Error("Unrecognized file format")
			return 1
		}

		serviceClient.CreateService(parsedConfig)

		return 0
	}

	if s.Describe {
		s.Logger.Info("Describing service: " + *serviceName + "@" + *serviceVersion)
		// TODO: validate request & receive parsed input to display
		serviceClient.DescribeService(*serviceName, *serviceVersion)

		return 0
	}

	if s.List {
		s.Logger.Info("Listing all services")
		serviceList, err := serviceClient.ListServices(*teamName, *serviceVersion, *isMature)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name", "Version", "Description", "Team", "Mature"}
		var tableData [][]interface{}

		for _, service := range serviceList {
			tableData = append(tableData, []interface{}{
				service.Name,
				service.Version,
				service.Description,
				strings.Join(service.Team, ","),
				service.Mature,
			})
		}

		err = table.Write(tableHeaders, tableData)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		return 0
	}

	if s.Label {
		if *isMature {
			s.Logger.Warn("Marking " + *serviceName + "@" + *serviceVersion + " as mature")
			// TODO: validate request
			serviceClient.MarkMature(*serviceName, *serviceVersion)
		}

		return 0
	}

	if s.Deploy {
		s.Logger.Warn("Deploying service: " + *serviceName + "@" + *serviceVersion + " in " + *envName + "/" + *infraName)
		// TODO: validate request
		serviceClient.DeployService(*serviceName, *serviceVersion, *infraName, *envName)

		return 0
	}

	if s.Destroy {
		s.Logger.Info("Destroying service: " + *serviceName + "@" + *serviceVersion + " in " + *envName + "/" + *infraName)
		// TODO: call PG api that destroys a service version from given env
		// DELETE /deploy?service=<v>&version=<version>&env=<env>&infra=<infra>

		return 0

	}

	if s.Status {
		s.Logger.Info("Fetching status for service: " + *serviceName + " in " + *envName + "/" + *infraName)
		// TODO: call PG api that returns status of service in env
		// GET /profileStatus?service=<service>&env=<env>&infra=<infra>

		return 0
	}

	if s.Logs {
		s.Logger.Info("Fetching logs for service: " + *serviceName + " in " + *envName + "/" + *infraName)
		// TODO: call PG api that returns execution logs of service in env
		// GET /serviceLogs?v=<service>&env=<env>&infra=<infra>

		return 0
	}

	if s.Delete {
		s.Logger.Warn("Deleting service: " + *serviceName + "@" + *serviceVersion)
		serviceClient.DeleteService(*serviceName, *serviceVersion)

		return 0
	}

	s.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (s *Service) Help() string {
	if s.Create {
		return commandHelper("create", "service", []string{
			"--file=yaml file to read service definition",
		})
	}

	if s.Describe {
		return commandHelper("describe", "service", []string{
			"--name=name of service to describe",
			"--version=version of service to describe",
		})
	}

	if s.List {
		return commandHelper("list", "service", []string{
			"--team=name of team",
			"--version=version of services to be listed",
			"--mature (mature marked service versions)",
		})
	}

	if s.Label {
		return commandHelper("label", "service", []string{
			"--name=name of service to label",
			"--version=version of service to label",
			"--mature (mark service version as mature)",
		})
	}

	if s.Deploy {
		return commandHelper("deploy", "service", []string{
			"--name=name of service to deploy",
			"--version=version of service to deploy",
			"--env=name of env to use",
			"--infra=name of infra to deploy service in",
		})
	}

	if s.Destroy {
		return commandHelper("destroy", "service", []string{
			"--name=name of service to destroy",
			"--version=version of service to destroy",
			"--env=name of env to use",
			"--infra=name of infra to destroy service from",
		})
	}

	if s.Delete {
		return commandHelper("delete", "service", []string{
			"--name=name of service to delete",
			"--version=version of service to delete",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (s *Service) Synopsis() string {
	if s.Create {
		return "create a service"
	}

	if s.Describe {
		return "describe a service version"
	}

	if s.List {
		return "list all services"
	}

	if s.Label {
		return "label a service version"
	}

	if s.Deploy {
		return "deploy a service"
	}

	if s.Destroy {
		return "destroy a deployed service"
	}

	if s.Delete {
		return "delete a service version"
	}

	return defaultHelper()
}
