package commands

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/dream11/odin/api/service"
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
	filePath := flagSet.String("file", "", "file to read service config")
	serviceName := flagSet.String("name", "", "name of service to be used")
	serviceVersion := flagSet.String("version", "", "version of service to be used")
	envName := flagSet.String("env", "", "name of environment to deploy the service in")
	teamName := flagSet.String("team", "", "name of user's team")
	component := flagSet.String("component", "", "name of service component")
	configStoreNamespace := flagSet.String("d11-config-store-namespace", "", "config store branch/tag to use")
	label := flagSet.String("label", "", "name of the label")
	provisioningConfigFile := flagSet.String("provisioning", "", "file to read provisioning config")

	err := flagSet.Parse(args)
	if err != nil {
		s.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if s.Create {

		isFilePresent := len(*filePath) > 0
		isServiceNamePresent := len(*serviceName) > 0
		isServiceVersionPresent := len(*serviceVersion) > 0

		if isFilePresent && (isServiceNamePresent || isServiceVersionPresent) {
			s.Logger.Error("--name and --version should not be provided when --file is provided.")
			return 1
		} else if (!isFilePresent && isServiceNamePresent && !isServiceVersionPresent) ||
			(!isFilePresent && !isServiceNamePresent && isServiceVersionPresent) {
			s.Logger.Error("Please provide both --name and --version.")
			return 1
		} else if !isFilePresent && !isServiceNamePresent && !isServiceVersionPresent {
			*filePath = "service.json"
		}

		emptyCreateParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion})
		if len(emptyCreateParameters) == 0 {
			serviceClient.RebuildServiceStream(*serviceName, *serviceVersion)
			s.Logger.Output("Command to check status of images")
			s.Logger.ItalicEmphasize(fmt.Sprintf("odin status service --name %s --version %s", *serviceName, *serviceVersion))
			return 0
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

		serviceClient.CreateServiceStream(parsedConfig)
		return 0
	}

	if s.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Describing service: " + *serviceName)
			serviceResp, err := serviceClient.DescribeService(*serviceName, *serviceVersion, *component)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			var details []byte
			if len(*component) == 0 {
				s.Logger.Info(serviceResp.Name + "@" + serviceResp.Version + " details!")
				details, err = json.MarshalIndent(serviceResp, "", "  ")
			} else {
				s.Logger.Info(fmt.Sprintf("%s component details for %s@%s", *component, serviceResp.Name, serviceResp.Version))
				details, err = json.MarshalIndent(serviceResp.Components[0], "", "  ")
			}

			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			s.Logger.Output(fmt.Sprintf("\n%s", details))
			if len(*component) == 0 {
				s.Logger.Output("Command to get component details")
				s.Logger.ItalicEmphasize(fmt.Sprintf("odin describe service --name %s --version <serviceVersion> --component <componentName>", *serviceName))
			}
			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.List {
		s.Logger.Info("Listing all services")
		serviceList, err := serviceClient.ListServices(*teamName, *serviceVersion, *serviceName, *label)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		var tableHeaders []string
		if len(*serviceName) == 0 {
			tableHeaders = []string{"Name", "Latest Version", "Description", "Team"}
		} else {
			tableHeaders = []string{"Name", "Version", "Description", "Team"}
		}
		var tableData [][]interface{}

		for _, service := range serviceList {
			tableData = append(tableData, []interface{}{
				service.Name,
				service.Version,
				service.Description,
				service.Team,
			})
		}

		table.Write(tableHeaders, tableData)

		s.Logger.Output("\nCommand to describe service")
		s.Logger.ItalicEmphasize("odin describe service --name <serviceName> --version <serviceVersion>")
		return 0
	}

	if s.Label {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion, "--label": *label})
		if len(emptyParameters) == 0 {
			serviceClient.LabelService(*serviceName, *serviceVersion, *label)
			s.Logger.Success("Successfully labelled " + *serviceName + "@" + *serviceVersion + " with " + *label)
			return 0
		}
		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.Unlabel {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion, "--label": *label})
		if len(emptyParameters) == 0 {
			serviceClient.UnlabelService(*serviceName, *serviceVersion, *label)
			s.Logger.Success("Successfully unlabelled " + *serviceName + "@" + *serviceVersion + " from " + *label)
			return 0
		}
		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.CreateDeploy {
		emptyCreateParameters := emptyParameters(map[string]string{"--env": *envName})
		if len(emptyCreateParameters) == 0 {

			if len(*filePath) != 0 {
				serviceDefinition, err := file.Read(*filePath)
				if err != nil {
					s.Logger.Error("Unable to read from " + *filePath + "\n" + err.Error())
					return 1
				}

				var parsedDefinition interface{}

				if strings.Contains(*filePath, ".yaml") || strings.Contains(*filePath, ".yml") {
					err = yaml.Unmarshal(serviceDefinition, &parsedDefinition)
					if err != nil {
						s.Logger.Error("Unable to parse YAML. " + err.Error())
						return 1
					}
				} else if strings.Contains(*filePath, ".json") {
					err = json.Unmarshal(serviceDefinition, &parsedDefinition)
					if err != nil {
						s.Logger.Error("Unable to parse JSON. " + err.Error())
						return 1
					}
				} else {
					s.Logger.Error("Unrecognized file format")
					return 1
				}
				emptyFileParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion})
				split := strings.Split(emptyFileParameters, ",")
				if len(split) < 2 {
					s.Logger.Error("--name and --version should not be provided when --file is provided.")
					return 1
				}

				serviceClient.BuildAndDeployServiceStream(parsedDefinition, *envName, *configStoreNamespace, *serviceName, *serviceVersion)

			} else {

				emptyCreateParameters = emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion})
				if len(emptyCreateParameters) == 0 {
					serviceClient.BuildAndDeployServiceStream(nil, *envName, *configStoreNamespace, *serviceName, *serviceVersion)

				} else {

					s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyCreateParameters))
					return 1
				}
			}

			return 0
		}
		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyCreateParameters))
		return 1
	}

	if s.Deploy {

		isEnvPresent := len(*envName) > 0
		isFilePresent := len(*filePath) > 0
		isServiceNamePresent := len(*serviceName) > 0
		isServiceVersionPresent := len(*serviceVersion) > 0

		if !isEnvPresent {
			s.Logger.Error("--env is mandatory")
			return 1
		}

		if isFilePresent && (isServiceNamePresent || isServiceVersionPresent) {
			s.Logger.Error("--name and --version should not be provided when --file is provided.")
			return 1
		} else if !isFilePresent && (!isServiceNamePresent || !isServiceVersionPresent) {
			s.Logger.Error("Please provide both --name and --version.")
			return 1
		}

		emptyUnreleasedParameters := emptyParameters(map[string]string{"--env": *envName, "--file": *filePath})
		if len(emptyUnreleasedParameters) == 0 {
			var serviceDefinition map[string]interface{}

			err, parsedConfig := parseFile(*filePath)
			serviceDefinition = parsedConfig.(map[string]interface{})
			if err != nil {
				s.Logger.Error("Error while parsing service file, err: \n" + err.Error())
				return 1
			}

			return s.deployUnreleasedService(envName, serviceDefinition, provisioningConfigFile, configStoreNamespace)
		}

		emptyReleasedParameters := emptyParameters(map[string]string{"--env": *envName, "--name": *serviceName, "--version": *serviceVersion})
		if len(emptyReleasedParameters) == 0 {
			return s.deployReleasedService(envName, serviceName, serviceVersion, provisioningConfigFile, configStoreNamespace)
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyUnreleasedParameters))
		return 1
	}

	if s.Undeploy {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName, "--env": *envName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Initiating service un-deploy: " + *serviceName + " from environment " + *envName)
			serviceClient.UnDeployServiceStream(*serviceName, *envName)

			return 0
		}
		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.Status {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion})
		if len(emptyParameters) == 0 {
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
					componentStatus.VM,
					componentStatus.Container,
				})
			}

			table.Write(tableHeaders, tableData)

			s.Logger.Output("\nCommand to deploy service")
			s.Logger.ItalicEmphasize(fmt.Sprintf("odin deploy service --name %s --version %s --env <envName>", *serviceName, *serviceVersion))
			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	s.Logger.Error("Not a valid command")
	return 127
}

func (s *Service) deployUnreleasedService(envName *string, serviceDefinition map[string]interface{}, provisioningConfigFile *string, configStoreNamespace *string) int {
	serviceName := serviceDefinition["name"].(string)
	serviceVersion := serviceDefinition["version"].(string)

	emptyParameters := emptyParameters(map[string]string{"name": serviceName, "version": serviceVersion})
	if len(emptyParameters) != 0 {
		s.Logger.Error(fmt.Sprintf("%s mandatory in the service definition file.", emptyParameters))
		return 1
	}

	rebuildService, forceService, parsedProvisioningConfig, i, done := s.validateDeployService(envName, serviceName, serviceVersion, provisioningConfigFile)
	if done {
		return i
	}

	s.Logger.Debug(fmt.Sprintf("%s: %s : %s: %s: %t: %t", serviceName, serviceVersion, *envName, *configStoreNamespace, forceService, rebuildService))
	s.Logger.Info("Initiating service deployment: " + serviceName + "@" + serviceVersion + " in " + *envName)
	serviceClient.DeployUnreleasedServiceStream(serviceDefinition, parsedProvisioningConfig, *envName, *configStoreNamespace, forceService, rebuildService)

	return 0
}

func (s *Service) deployReleasedService(envName *string, serviceName *string, serviceVersion *string,
	provisioningConfigFile *string, configStoreNamespace *string) int {

	rebuildService, forceService, parsedProvisioningConfig, i, done := s.validateDeployService(envName, *serviceName, *serviceVersion, provisioningConfigFile)
	if done {
		return i
	}

	s.Logger.Debug(fmt.Sprintf("%s: %s : %s: %s: %t: %t", *serviceName, *serviceVersion, *envName, *configStoreNamespace, forceService, rebuildService))
	s.Logger.Info("Initiating service deployment: " + *serviceName + "@" + *serviceVersion + " in " + *envName)
	serviceClient.DeployReleasedServiceStream(*serviceName, *serviceVersion, *envName, *configStoreNamespace, forceService, rebuildService, parsedProvisioningConfig)

	return 0
}

/*
validateDeployService
	returns rebuildService: bool, forceService: bool, parsedProvisioningConfig: interface{}, exitCode int, toExit bool
*/
func (s *Service) validateDeployService(envName *string, serviceName string, serviceVersion string, provisioningConfigFile *string) (bool, bool, interface{}, int, bool) {
	envServices, err := envClient.DescribeEnv(*envName, "", "")

	if err != nil {
		s.Logger.Error(err.Error())
		return false, false, nil, 1, true
	}

	envService := service.Service{}

	rebuildService := false
	forceService := false

	for _, curService := range envServices.Services {
		if curService.Name == serviceName && curService.Version == serviceVersion {
			rebuildService = true
			envService = curService
			break
		}
		if curService.Name == serviceName && curService.Version != serviceVersion {
			forceService = true
			envService = curService
			break
		}
	}

	if forceService {
		s.Logger.Info(fmt.Sprintf("service: %s already exists in the env with different version: %s", serviceName, envService.Version))
		s.Logger.Output("Press [Y] to force deploy service or press [n] to skip service deploy.\n")
		message := fmt.Sprintf("Update version of Service %s : %s -> %s[Y/n]: ", serviceName, envService.Version, serviceVersion)

		allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
		val, err := s.Input.AskWithConstraints(message, allowedInputs)

		if err != nil {
			s.Logger.Error(err.Error())
			return false, false, nil, 1, true
		}

		if val != "Y" {
			s.Logger.Info("Skipping force service deploy")
			return false, false, nil, 0, true
		}
	}

	var parsedProvisioningConfig interface{}

	if len(*provisioningConfigFile) > 0 {
		err, parsedConfig := parseFile(*provisioningConfigFile)
		parsedProvisioningConfig = parsedConfig

		if err != nil {
			s.Logger.Error("Error while parsing provisioning file, err: \n" + err.Error())
			return false, false, nil, 1, true
		}
	}
	return rebuildService, forceService, parsedProvisioningConfig, 0, false
}

func parseFile(filePath string) (error, interface{}) {
	if len(filePath) != 0 {
		var parsedDefinition interface{}

		fileDefinition, err := file.Read(filePath)
		if err != nil {
			return errors.New("Unable to read from " + filePath + "\n" + err.Error()), parsedDefinition
		}

		if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
			err = yaml.Unmarshal(fileDefinition, &parsedDefinition)
			if err != nil {
				return errors.New("Unable to parse YAML. " + err.Error()), parsedDefinition
			}
		} else if strings.Contains(filePath, ".json") {
			err = json.Unmarshal(fileDefinition, &parsedDefinition)
			if err != nil {
				return errors.New("Unable to parse JSON. " + err.Error()), parsedDefinition
			}
		} else {
			return errors.New("unrecognized file format"), parsedDefinition
		}
		return nil, parsedDefinition
	}
	return nil, errors.New("filepath cannot be empty")
}

// Help : returns an explanatory string
func (s *Service) Help() string {
	if s.Create {
		return commandHelper("create", "service", "", []Options{
			{Flag: "--file", Description: "json file to read service definition"},
			{Flag: "--name", Description: "name of the service (required if using --rebuild)"},
			{Flag: "--version", Description: "version of the service (required if using --rebuild)"},
		})
	}

	if s.Describe {
		return commandHelper("describe", "service", "", []Options{
			{Flag: "--name", Description: "name of service to describe"},
			{Flag: "--version", Description: "version of service to describe"},
			{Flag: "--component", Description: "name of component to describe"},
		})
	}

	if s.List {
		return commandHelper("list", "service", "", []Options{
			{Flag: "--name", Description: "name of service"},
			{Flag: "--version", Description: "version of services to be listed"},
			{Flag: "--team", Description: "name of team"},
			{Flag: "--label", Description: "name of label"},
		})
	}

	if s.Label {
		return commandHelper("label", "service", "", []Options{
			{Flag: "--name", Description: "name of service to label"},
			{Flag: "--version", Description: "version of service to label"},
			{Flag: "--label", Description: "name of the label"},
		})
	}

	if s.Deploy {
		return commandHelper("deploy", "service", "", []Options{
			{Flag: "--name", Description: "name of service to deploy"},
			{Flag: "--version", Description: "version of service to deploy"},
			{Flag: "--env", Description: "name of environment to deploy service in"},
			{Flag: "--d11-config-store-namespace", Description: "config store branch/tag to use"},
			{Flag: "--provisioning", Description: "file to read provisioning config."},
		})
	}

	if s.Undeploy {
		return commandHelper("deploy", "service", "", []Options{
			{Flag: "--name", Description: "name of service to undeploy"},
			{Flag: "--env", Description: "name of environment to undeploy service in"},
		})
	}

	if s.Status {
		return commandHelper("status", "service", "", []Options{
			{Flag: "--name", Description: "name of service"},
			{Flag: "--version", Description: "version of service"},
		})
	}

	if s.CreateDeploy {
		return commandHelper("create and deploy", "service", "", []Options{
			{Flag: "--file", Description: "service definition file"},
			{Flag: "--env", Description: "name of environment to deploy service in"},
			{Flag: "--name", Description: "name of an already created service"},
			{Flag: "--version", Description: "version of the already created service"},
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

	if s.Undeploy {
		return "undeploy a service"
	}

	if s.Status {
		return "get status of a service version"
	}

	if s.CreateDeploy {
		return "create and deploy a service in an env"
	}

	return defaultHelper()
}
