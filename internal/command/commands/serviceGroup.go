package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"gopkg.in/yaml.v3"
	"strings"
)

// initiate backend client for service
var serviceGroupClient backend.ServiceGroup

// ServiceGroup : command declaration
type ServiceGroup command

// Run : implements the actual functionality of the command
func (s *ServiceGroup) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	filePath := flagSet.String("file", "service.yaml", "file to read service config")
	serviceGroupName := flagSet.String("name", "", "name of service-group to be used")
	serviceName := flagSet.String("service", "", "name of service in service-group")

	err := flagSet.Parse(args)
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
		serviceDataMap := parsedConfig.(map[string]interface{})

		s.Logger.Info(fmt.Sprintf("Service-group creation started for %s  ", serviceDataMap["services"]))
		serviceGroupResp, err := serviceGroupClient.CreateServiceGroup(parsedConfig)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		s.Logger.Success(fmt.Sprintf("%s", serviceGroupResp))
		return 0
	}

	if s.List {
		s.Logger.Info("Listing all service-groups")
		serviceList, err := serviceGroupClient.ListServiceGroups(*serviceGroupName, *serviceName)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name"}
		var tableData [][]interface{}

		for _, service := range serviceList {
			tableData = append(tableData, []interface{}{
				service.Name,
			})
		}

		err = table.Write(tableHeaders, tableData)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		s.Logger.Output("\nCommand to describe service-group")
		s.Logger.ItalicEmphasize("odin describe service-group --name <serviceName>")
		return 0
	}

	if s.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceGroupName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Describing service-group: " + *serviceGroupName)
			serviceGroupResp, err := serviceGroupClient.DescribeService(*serviceGroupName)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			var details []byte
			s.Logger.Info(serviceGroupResp.Name + " details!")
			details, err = yaml.Marshal(serviceGroupResp)

			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			s.Logger.Output(fmt.Sprintf("\n%s", details))
			s.Logger.Output("Command to get service details")
			s.Logger.ItalicEmphasize("odin describe service --name <serviceName> --version <serviceVersion>")
			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	s.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (s *ServiceGroup) Help() string {
	if s.Create {
		return commandHelper("create", "service-group", []string{
			"--file=yaml file to read service-group definition",
		})
	}

	if s.List {
		return commandHelper("list", "service-group", []string{
			"--name=name of the service-group",
			"--service=name of service in the service-group",
		})
	}

	if s.Describe {
		return commandHelper("describe", "service", []string{
			"--name=name of the service-group to describe",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (s *ServiceGroup) Synopsis() string {
	if s.Create {
		return "create a service-group"
	}

	if s.List {
		return "list all service-groups"
	}

	if s.Describe {
		return "describe a service-group"
	}

	return defaultHelper()
}
