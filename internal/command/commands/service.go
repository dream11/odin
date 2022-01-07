package commands

import (
	"encoding/json"
	"flag"
	"strings"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"gopkg.in/yaml.v3"
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
	envName := flagSet.String("env", "", "name of environment to deploy the service in")
	teamName := flagSet.String("team", "", "name of user's team")
	isMature := flagSet.Bool("mature", false, "mark service version as matured")
	detailed := flagSet.Bool("detailed", false, "get detailed view")
	rebuild := flagSet.Bool("rebuild", false, "rebuild executor for creating images")

	err := flagSet.Parse(args)
	if err != nil {
		s.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if s.Create {

		if *rebuild {
			if emptyParameterValidation([]string{*serviceName, *serviceVersion}) {
				serviceClient.RebuildService(*serviceName, *serviceVersion)
				return 0
			}
			s.Logger.Error("service name & version cannot be blank")
			return 1
		}

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
		if emptyParameterValidation([]string{*serviceName}) {
			s.Logger.Info("Describing service: " + *serviceName + "@" + *serviceVersion)
			serviceResp, err := serviceClient.DescribeService(*serviceName, *serviceVersion)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			s.Logger.Info(serviceResp.Name + "@" + serviceResp.Version + " details!")
			details, err := yaml.Marshal(serviceResp)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			s.Logger.Output(string(details))

			return 0
		}

		s.Logger.Error("service name cannot be blank")
		return 1
	}

	if s.List {
		s.Logger.Info("Listing all services")
		serviceList, err := serviceClient.ListServices(*teamName, *serviceVersion, *serviceName, *isMature)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		if *detailed {
			for _, service := range serviceList {
				s.Logger.Info("Service definition for: " + service.Name + "@" + service.Version)

				serviceYaml, err := yaml.Marshal(service)
				if err != nil {
					s.Logger.Error("Unable to parse environment definition! " + err.Error())
					return 1
				}

				s.Logger.Output(string(serviceYaml))
			}
		} else {
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
		}

		return 0
	}

	if s.Label {
		if emptyParameterValidation([]string{*serviceName, *serviceVersion}) {

			// Add more labels to this condition
			if !*isMature {
				s.Logger.Error("No label specified")
				return 1
			}

			if *isMature {
				s.Logger.Info("Marking " + *serviceName + "@" + *serviceVersion + " as mature")
				serviceClient.MarkMature(*serviceName, *serviceVersion)
			}
			return 0
		}

		s.Logger.Error("service name & version cannot be blank")
		return 1
	}

	if s.Deploy {
		if emptyParameterValidation([]string{*serviceName, *serviceVersion, *envName}) {
			s.Logger.Warn("Deploying service: " + *serviceName + "@" + *serviceVersion + " in " + *envName)

			var parsedConfig interface{}

			configData, err := file.Read(*filePath)
			if err == nil {
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
			} else {
				// initialise config as empty json
				parsedConfig = make(map[string]string)
			}

			serviceClient.DeployService(*serviceName, *serviceVersion, *envName, parsedConfig)

			return 0
		}

		s.Logger.Error("service name, version and environment name cannot be blank")
		return 1
	}

	if s.Delete {
		if emptyParameterValidation([]string{*serviceName, *serviceVersion}) {
			s.Logger.Warn("Deleting service: " + *serviceName + "@" + *serviceVersion)
			serviceClient.DeleteService(*serviceName, *serviceVersion)

			return 0
		}

		s.Logger.Error("service name & version cannot be blank")
		return 1
	}

	if s.Status {
		if emptyParameterValidation([]string{*serviceName, *serviceVersion}) {
			s.Logger.Info("Getting status of service: " + *serviceName + "@" + *serviceVersion)
			serviceStatus, err := serviceClient.StatusService(*serviceName, *serviceVersion)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"Component Name", "AMI", "DOCKER IMAGE"}
			var tableData [][]interface{}
			for _, componentStatus := range serviceStatus {
				tableData = append(tableData, []interface{}{
					componentStatus.Name,
					componentStatus.Ec2,
					componentStatus.Docker,
				})
			}

			err = table.Write(tableHeaders, tableData)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}
			return 0
		}

		s.Logger.Error("service name & version cannot be blank")
		return 1
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
			"--detailed (get a detailed view)",
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
			"--env=name of environment to deploy service in",
			"--file=name of config file",
		})
	}

	if s.Delete {
		return commandHelper("delete", "service", []string{
			"--name=name of service to delete",
			"--version=version of service to delete",
		})
	}

	if s.Status {
		return commandHelper("status", "service", []string{
			"--name=name of service",
			"--version=version of service",
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

	if s.Delete {
		return "delete a service version"
	}

	if s.Status {
		return "get status of a service version"
	}

	return defaultHelper()
}
