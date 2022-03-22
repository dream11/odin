package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
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

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (s *ServiceGroup) Synopsis() string {
	if s.Create {
		return "create a service-group"
	}

	return defaultHelper()
}
