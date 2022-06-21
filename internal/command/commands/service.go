package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dream11/odin/api/service"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/dir"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/utils"
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
	directoryPath := flagSet.String("path", "", "path to directory containing service definition and provisioning config")

	err := flagSet.Parse(args)
	if err != nil {
		s.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if s.Release {

		isDirectoryPathPresent := len(*directoryPath) > 0
		isServiceNamePresent := len(*serviceName) > 0
		isServiceVersionPresent := len(*serviceVersion) > 0

		if isDirectoryPathPresent && (isServiceNamePresent || isServiceVersionPresent) {
			s.Logger.Error("--name and --version should not be provided when --path is provided.")
			return 1
		} else if (!isDirectoryPathPresent && isServiceNamePresent && !isServiceVersionPresent) ||
			(!isDirectoryPathPresent && !isServiceNamePresent && isServiceVersionPresent) {
			s.Logger.Error("Please provide both --name and --version.")
			return 1
		} else if !isDirectoryPathPresent && !isServiceNamePresent && !isServiceVersionPresent {
			s.Logger.Error("Please provide --path or --name and --version both.")
		}

		emptyCreateParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion})
		if len(emptyCreateParameters) == 0 {
			serviceClient.RebuildServiceStream(*serviceName, *serviceVersion)
			s.Logger.Output("Command to check status of images")
			s.Logger.ItalicEmphasize(fmt.Sprintf("odin status service --name %s --version %s", *serviceName, *serviceVersion))
			return 0
		}

		if exists, err := dir.Exists(*directoryPath); !exists || err != nil {
			s.Logger.Error("Provided directory path : " + *directoryPath + " , is not valid\n" + err.Error())
			return 1
		}

		serviceDefinition, serviceDefinitionPath, err := file.FindAndReadAllAllowedFormat(*directoryPath+"/definition", []string{".json", ".yml", ".yaml"})
		if err != nil {
			s.Logger.Error("Unable to read from " + *directoryPath + "/definition.json\n")
			return 1
		}

		parsedServiceDefinition, err := utils.ParserYmlOrJson(serviceDefinitionPath, serviceDefinition)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		envTypes, err := envClient.EnvTypes()
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		allFiles, err := dir.SubDirs(*directoryPath)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		envFileMap := make(map[string]string)
		r, _ := regexp.Compile(`provisioning-([a-zA-Z]*)\.json`)
		for _, file := range allFiles {
			envType := r.FindStringSubmatch(file)
			if len(envType) > 1 {
				envFileMap[envType[1]] = file
			}
		}

		provisioningConfigMap := make(map[string]interface{})
		for _, envType := range envTypes.EnvTypes {
			if envFileMap[envType] == "" {
				continue
			}
			f := filepath.Join(*directoryPath, utils.GetProvisioningFileName(envType))

			data, provisioningFilePath, err := file.FindAndReadAllAllowedFormat(f, []string{".json", ".yml", ".yaml"})
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}
			parsedProvisioningConfig, err := utils.ParserYmlOrJson(provisioningFilePath, data)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}
			provisioningConfigMap[envType] = parsedProvisioningConfig
		}
		serviceClient.CreateServiceStream(parsedServiceDefinition, provisioningConfigMap)
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
				s.Logger.Output("\nCommand to get component details")
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
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion, "--env": *envName})
		if len(emptyParameters) == 0 {
			envServices, err := envClient.DescribeEnv(*envName, "", "")

			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			envService := service.Service{}

			rebuildService := false
			forceService := false

			for _, curService := range envServices.Services {
				if curService.Name == *serviceName && curService.Version == *serviceVersion {
					rebuildService = true
					envService = curService
					break
				}
				if curService.Name == *serviceName && curService.Version != *serviceVersion {
					forceService = true
					envService = curService
					break
				}
			}

			if forceService {
				s.Logger.Info(fmt.Sprintf("service: %s already exists in the env with different version: %s", *serviceName, envService.Version))
				s.Logger.Output("Press [Y] to force deploy service or press [n] to skip service deploy.\n")
				message := fmt.Sprintf("Update version of Service %s : %s -> %s[Y/n]: ", *serviceName, envService.Version, *serviceVersion)

				allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
				val, err := s.Input.AskWithConstraints(message, allowedInputs)

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if val != "Y" {
					s.Logger.Info("Skipping force service deploy")
					return 0
				}
			}

			s.Logger.Debug(fmt.Sprintf("%s: %s : %s: %s: %t: %t", *serviceName, *serviceVersion, *envName, *configStoreNamespace, forceService, rebuildService))
			s.Logger.Info("Initiating service deployment: " + *serviceName + "@" + *serviceVersion + " in " + *envName)
			serviceClient.DeployServiceStream(*serviceName, *serviceVersion, *envName, *configStoreNamespace, forceService, rebuildService)

			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
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

// Help : returns an explanatory string
func (s *Service) Help() string {
	if s.Release {
		return commandHelper("create", "service", "", []Options{
			{Flag: "--path, -p", Description: "path to directory containing service definition and provisioning config"},
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
	if s.Release {
		return "release a service"
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
