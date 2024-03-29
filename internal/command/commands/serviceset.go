package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/dream11/odin/api/serviceset"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/utils"
	"gopkg.in/yaml.v3"
)

// initiate backend client for service-set
var serviceSetClient backend.ServiceSet

// ServiceSet : command declaration
type ServiceSet command

// Run : implements the actual functionality of the command
func (s *ServiceSet) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	filePath := flagSet.String("file", "", "file to read service-set config")
	serviceSetName := flagSet.String("name", "", "name of service-set to be used")
	serviceName := flagSet.String("service", "", "name of service in service-set")
	envName := flagSet.String("env", "", "name of environment to deploy the service-set in")
	force := flagSet.Bool("force", false, "forcefully deploy service-set into the Env")
	configStoreNamespace := flagSet.String("d11-config-store-namespace", "", "config store branch/tag to use")

	err := flagSet.Parse(args)
	if err != nil {
		s.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if s.Create {
		commandStatus := s.CreateServiceSet(*filePath, false)
		return commandStatus.(int)
	}

	if s.List {
		s.Logger.Info("Listing all service-sets")
		serviceSetList, err := serviceSetClient.ListServiceSet(*serviceSetName, *serviceName)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name"}
		var tableData [][]interface{}

		for _, serviceSet := range serviceSetList {
			tableData = append(tableData, []interface{}{
				serviceSet.Name,
			})
		}

		table.Write(tableHeaders, tableData)

		s.Logger.Output("\nCommand to describe serviceset")
		s.Logger.ItalicEmphasize("odin describe service-set --name <serviceSetName>")
		return 0
	}

	if s.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceSetName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Describing service-set: " + *serviceSetName)
			serviceSetResp, err := serviceSetClient.DescribeServiceSet(*serviceSetName)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			var details []byte
			s.Logger.Info(serviceSetResp.Name + " details!")
			details, err = yaml.Marshal(serviceSetResp)

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

	if s.Delete {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceSetName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Deleting service-set: " + *serviceSetName)
			serviceSetClient.DeleteServiceSet(*serviceSetName)

			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.Deploy {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}
		isEnvPresent := len(*envName) > 0
		isFilePresent := len(*filePath) > 0
		isServiceSetNamePresent := len(*serviceSetName) > 0

		if !isEnvPresent {
			s.Logger.Error("--env is mandatory")
			return 1
		}
		if isFilePresent && isServiceSetNamePresent {
			s.Logger.Error("--name should not be provided when --file is provided.")
			return 1
		} else if !isFilePresent && !isServiceSetNamePresent {
			s.Logger.Error("Please provide either --name or --file.")
			return 1
		}

		if isFilePresent {
			parsedConfig := s.CreateServiceSet(*filePath, true)
			serviceDataMap := parsedConfig.(map[string]interface{})
			exitStatus := s.DeployUnDeployServiceSet(serviceDataMap["name"].(string), *envName, *configStoreNamespace, *force, true, "deploy")
			return exitStatus
		}

		if isServiceSetNamePresent {
			exitStatus := s.DeployUnDeployServiceSet(*serviceSetName, *envName, *configStoreNamespace, *force, false, "deploy")
			return exitStatus
		}
	}

	if s.Undeploy {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceSetName, "--env": *envName})
		if len(emptyParameters) == 0 {
			serviceSetUsingFile := serviceSetClient.IdentifyServiceSetType(*serviceSetName)
			exitStatus := s.DeployUnDeployServiceSet(*serviceSetName, *envName, *configStoreNamespace, *force, serviceSetUsingFile, "undeploy")
			return exitStatus
		}
		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	s.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (s *ServiceSet) Help() string {
	if s.Create {
		return commandHelper("create", "service-set", "", []Options{
			{Flag: "--file", Description: "yaml/json file to read service-set definition"},
		})
	}

	if s.List {
		return commandHelper("list", "service-set", "", []Options{
			{Flag: "--name", Description: "name of the service-set"},
			{Flag: "--service", Description: "name of service in the service-set"},
		})
	}

	if s.Describe {
		return commandHelper("describe", "service", "", []Options{
			{Flag: "--name", Description: "name of the service-set to describe"},
		})
	}

	if s.Delete {
		return commandHelper("delete", "service-set", "", []Options{
			{Flag: "--name", Description: "name of service-set to delete"},
		})
	}

	if s.Deploy {
		return commandHelper("deploy", "service-set", "", []Options{
			{Flag: "--name", Description: "name of service-set to deploy"},
			{Flag: "--env", Description: "name of environment to deploy service-set in"},
			{Flag: "--force", Description: "forcefully deploy service-set into the Env"},
			{Flag: "--d11-config-store-namespace", Description: "config store branch/tag to use"},
			{Flag: "--file", Description: "json file to read temporary service-set definition "},
		})
	}

	if s.Undeploy {
		return commandHelper("undeploy", "service-set", "", []Options{
			{Flag: "--name", Description: "name of service-set to undeploy"},
			{Flag: "--env", Description: "name of environment to undeploy service-set in"},
			{Flag: "--force", Description: "forcefully undeploy service-set from the Env"},
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (s *ServiceSet) Synopsis() string {
	if s.Create {
		return "create a service-set"
	}

	if s.List {
		return "list all service-sets"
	}

	if s.Describe {
		return "describe a service-set"
	}

	if s.Delete {
		return "delete a service-set"
	}

	if s.Deploy {
		return "deploy a service-set"
	}

	if s.Undeploy {
		return "undeploy a service-set"
	}

	return defaultHelper()
}

func (s *ServiceSet) DeployUnDeployServiceSet(serviceSetName string, envName string, configStoreNamespace string, force bool, serviceSetUsingFile bool, serviceSetAction string) int {
	var forceDeployUnDeployServices []serviceset.ListEnvService
	if !force {
		//get list of env services
		s.Logger.Debug(fmt.Sprintf("Env Services of service-set %s and env %s", serviceSetName, envName+"\n\n"))
		serviceSetList, err := serviceSetClient.ListEnvServices(serviceSetName, envName, "conflictedVersion", serviceSetUsingFile)

		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		if len(serviceSetList) > 0 {
			s.Logger.Output("Following services have conflicting versions in the Env: " + envName)
			//s.Logger.Output("Press [Y] to update the service version or press [n] to skip service.\n")
			s.Logger.Output(fmt.Sprintf("Press [Y] to %s the service with the conflicting version or press [n] to skip service.\n", serviceSetAction))
			allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
			var message string
			for _, serviceSet := range serviceSetList {
				if serviceSetAction == "deploy" {
					message = fmt.Sprintf("Update version of Service %s : %s -> %s[Y/n]: ", serviceSet.Name, serviceSet.EnvVersion, serviceSet.Version)
				} else {
					message = fmt.Sprintf("undeploy Service: %s with version: %s[Y/n]: ", serviceSet.Name, serviceSet.EnvVersion)
				}
				val, err := s.Input.AskWithConstraints(message, allowedInputs)

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if val == "Y" {
					forceDeployUnDeployServices = append(forceDeployUnDeployServices, serviceSet)
				}
			}
		}
	}

	/*deploy service-set*/
	if serviceSetAction == "deploy" {
		s.Logger.ItalicEmphasize("Deploying service-set: " + serviceSetName + " in " + envName + "\n\n")
		serviceSetClient.DeployServiceSet(serviceSetName, envName, configStoreNamespace, forceDeployUnDeployServices, force, serviceSetUsingFile)
		return 0
	} else {
		s.Logger.ItalicEmphasize("Undeploying service-set:  " + serviceSetName + " in " + envName + "\n\n")
		serviceSetClient.UndeployServiceSet(serviceSetName, envName, forceDeployUnDeployServices, force, serviceSetUsingFile)
		return 0
	}
}

func (s *ServiceSet) CreateServiceSet(filePath string, serviceSetUsingFile bool) interface{} {

	configData, err := file.Read(filePath)
	if err != nil {
		s.Logger.Error("Unable to read from " + filePath + "\n" + err.Error())
		return 1
	}

	var parsedConfig interface{}

	if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
		err = yaml.Unmarshal(configData, &parsedConfig)
		if err != nil {
			s.Logger.Error("Unable to parse YAML. " + err.Error())
			return 1
		}
	} else if strings.Contains(filePath, ".json") {
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
	if serviceSetUsingFile {
		serviceSetClient.CreateUpdateTempServiceSet(parsedConfig)
		s.Logger.Success(fmt.Sprintf("Temporary ServiceSet: %s initialised Successfully.", serviceDataMap["name"].(string)))
		return parsedConfig
	} else {
		serviceSetClient.CreateServiceSet(parsedConfig)
		s.Logger.Success(fmt.Sprintf("ServiceSet: %s created Successfully.", serviceDataMap["name"].(string)))
		return 0
	}

}
