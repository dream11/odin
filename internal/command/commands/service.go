package commands

import (
	"encoding/json"
	"errors"
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

const VOLATILE = "volatile"

// Run : implements the actual functionality of the command
func (s *Service) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	filePath := flagSet.String("file", "", "file to read service config or to provide options for service operations")
	serviceName := flagSet.String("name", "", "name of service to be used")
	serviceVersion := flagSet.String("version", "", "version of service to be used")
	envName := flagSet.String("env", "", "name of environment to deploy the service in")
	teamName := flagSet.String("team", "", "name of user's team")
	component := flagSet.String("component", "", "name of service component")
	configStoreNamespace := flagSet.String("d11-config-store-namespace", "", "config store branch/tag to use")
	label := flagSet.String("label", "", "name of the label")
	provisioningConfigFile := flagSet.String("provisioning", "", "file to read provisioning config")
	directoryPath := flagSet.String("path", "", "path to directory containing service definition and provisioning config")
	operation := flagSet.String("operation", "", "name of the operation to performed on the component")
	options := flagSet.String("options", "", "options for service operations")

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
			return 1
		}

		emptyCreateParameters := emptyParameters(map[string]string{"--name": *serviceName, "--version": *serviceVersion})
		if len(emptyCreateParameters) == 0 {
			serviceClient.RebuildServiceStream(*serviceName, *serviceVersion)
			s.Logger.Output("Command to check status of images")
			s.Logger.ItalicEmphasize(fmt.Sprintf("odin status service --name %s --version %s", *serviceName, *serviceVersion))
			return 0
		}

		if exists, err := dir.IsDir(*directoryPath); !exists || err != nil {
			s.Logger.Error("Provided directory path : " + *directoryPath + " , is not valid")
			return 1
		}

		serviceDefinition, serviceDefinitionPath, err := file.FindAndReadAllAllowedFormat(*directoryPath+"/definition", []string{".json", ".yml", ".yaml"})
		if err != nil {
			s.Logger.Error("Unable to read from " + *directoryPath + "/definition.json\n")
			return 1
		}

		// Throw error on empty service def file
		if len(serviceDefinition) == 0 {
			s.Logger.Error("service definition file(definition.json) cannot be empty")
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
		r, _ := regexp.Compile(`provisioning-([a-zA-Z]*)\.json`)
		provisioningConfigMap := make(map[string]interface{})

		for _, fileName := range allFiles {
			envType := r.FindStringSubmatch(fileName)
			if len(envType) == 0 {
				if fileName != "definition.json" {
					s.Logger.Warn(fmt.Sprintf("Ignoring %s. Provisioning config files should be in following format provisioning-<env_type>.json", fileName))
				}
			} else {
				if !utils.Contains(append(envTypes.EnvTypes, "default"), envType[1]) {
					s.Logger.Warn(fmt.Sprintf("Ignoring %s as env type %s does not exist", fileName, envType[1]))
					continue
				}
				f := filepath.Join(*directoryPath, utils.GetProvisioningFileName(envType[1]))
				data, provisioningFilePath, err := file.FindAndReadAllAllowedFormat(f, []string{".json", ".yml", ".yaml"})
				// Ignore empty provisioning files
				if len(data) == 0 {
					continue
				}
				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}
				parsedProvisioningConfig, err := utils.ParserYmlOrJson(provisioningFilePath, data)
				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}
				provisioningConfigMap[envType[1]] = parsedProvisioningConfig
			}
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
			tableHeaders = []string{"Name", "Latest Version", "Label", "Description", "Team"}
		} else {
			tableHeaders = []string{"Name", "Version", "Label", "Description", "Team"}
		}
		var tableData [][]interface{}

		for _, service := range serviceList {
			var labelList []string
			for _, label := range service.Labels {
				labelList = append(labelList, label.Name)
			}
			tableData = append(tableData, []interface{}{
				service.Name,
				service.Version,
				strings.Join(labelList, ", "),
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

	if s.Deploy {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}
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
			parsedConfig, err := parseFile(*filePath)
			if err != nil {
				s.Logger.Error("Error while parsing service file " + *filePath + " : " + err.Error())
				return 1
			}
			consent := s.askForConsent(envName)
			if consent == 1 {
				return 1
			}
			serviceDefinition := parsedConfig.(map[string]interface{})

			return s.deployUnreleasedService(envName, serviceDefinition, provisioningConfigFile, configStoreNamespace)
		}

		emptyReleasedParameters := emptyParameters(map[string]string{"--env": *envName, "--name": *serviceName, "--version": *serviceVersion})
		if len(emptyReleasedParameters) == 0 {
			consent := s.askForConsent(envName)
			if consent == 1 {
				return 1
			}
			return s.deployReleasedService(envName, serviceName, serviceVersion, provisioningConfigFile, configStoreNamespace)
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyUnreleasedParameters))
		return 1
	}

	if s.Undeploy {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName, "--env": *envName})
		if len(emptyParameters) == 0 {
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
					componentStatus.AWS_EC2,
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

	if s.Operate {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}

		isNamePresent := len(*serviceName) > 0
		isOptionsPresent := len(*options) > 0
		isFilePresent := len(*filePath) > 0
		isOperationPresnt := len(*operation) > 0
		isEnvNamePresent := len(*envName) > 0

		if !isNamePresent {
			s.Logger.Error("--name cannot be blank")
			return 1
		}

		if !isOperationPresnt {
			s.Logger.Error("--operation cannot be blank")
			return 1
		}

		if !isEnvNamePresent {
			s.Logger.Error("--env cannot be blank")
			return 1
		}

		if isOptionsPresent && isFilePresent {
			s.Logger.Error("You can provide either --options or --file but not both")
			return 1
		}

		if !isOptionsPresent && !isFilePresent {
			s.Logger.Error("You should provide either --options or --file")
			return 1
		}

		var optionsData map[string]interface{}

		if isFilePresent {
			parsedConfig, err := parseFile(*filePath)
			if err != nil {
				s.Logger.Error("Error while parsing service file " + *filePath + " : " + err.Error())
				return 1
			}
			optionsData = parsedConfig.(map[string]interface{})
		} else if isOptionsPresent {
			err = json.Unmarshal([]byte(*options), &optionsData)
			if err != nil {
				s.Logger.Error("Unable to parse JSON data " + err.Error())
				return 1
			}
		}

		if len(optionsData) == 0 {
			s.Logger.Error("You can't send an empty JSON data")
			return 1
		}

		consent := s.askForConsent(envName)

		if consent == 1 {
			return 1
		}

		dataForScalingConsent := map[string]interface{}{
			"env_name": *envName,
			"action":   *operation,
			"config":   optionsData,
		}
		scalingConsent := s.askForScalingConsent(serviceName, dataForScalingConsent)
		if scalingConsent == 1 {
			return 1
		}

		data := service.OperationRequest{
			EnvName: *envName,
			Operations: []service.Operation{
				{
					Name: *operation,
					Data: optionsData,
				},
			},
		}

		s.Logger.Info("Validating the operation: " + *operation + " on the service: " + *serviceName)

		validateOperateResponse, err := serviceClient.ValidateOperation(*serviceName, data)

		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		for _, operation := range validateOperateResponse.Response.Operations {
			isFeedbackRequired := operation.IsFeedbackRequired
			message := operation.Message

			if isFeedbackRequired {
				consentMessage := fmt.Sprintf("\n%s", message)

				allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
				val, err := s.Input.AskWithConstraints(consentMessage, allowedInputs)

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if val != "Y" {
					s.Logger.Info("Aborting the operation")
					return 1
				}
			}
		}

		serviceClient.OperateService(*serviceName, data)
		return 0
	}

	s.Logger.Error("Not a valid command")
	return 127
}

func (s *Service) askForConsent(envName *string) int {
	envTypeResp, err := envTypeClient.GetEnvType(*envName)
	if err != nil {
		s.Logger.Error(err.Error())
		return 1
	}
	if envTypeResp.Strict {
		consentMessage := fmt.Sprintf("\nYou are executing the above command on a restricted environment. Are you sure? Enter \033[1m%s\033[0m to continue:", *envName)
		val, err := s.Input.Ask(consentMessage)

		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		if val != *envName {
			s.Logger.Info("Aborting the operation")
			return 1
		}
	}
	return 0
}

func (s *Service) askForScalingConsent(serviceName *string, data map[string]interface{}) int {
	componentListResponse, err := serviceClient.ScalingServiceConsent(*serviceName, data)
	if err != nil {
		s.Logger.Error(err.Error())
		return 1
	}
	for _, component := range componentListResponse.Response {
		consentMessage := fmt.Sprintf("\nAs you have enabled ASG auto scaling policy for %s, It will no longer be scaled using Scaler post this operation. Do you wish to continue? [Y/n]:", component)
		allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
		val, err := s.Input.AskWithConstraints(consentMessage, allowedInputs)

		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		if val != "Y" {
			s.Logger.Info("\nAborting...")
			return 1
		}
	}
	return 0
}

func (s *Service) deployUnreleasedService(envName *string, serviceDefinition map[string]interface{}, provisioningConfigFile *string, configStoreNamespace *string) int {

	if serviceDefinition["name"] == nil || len(serviceDefinition["name"].(string)) == 0 {
		s.Logger.Error("Service name cannot be blank. Please provide service name")
		return 1
	}

	serviceName := serviceDefinition["name"].(string)
	serviceVersion := ""

	parsedProvisioningConfig, i, done := s.validateDeployService(envName, serviceName, serviceVersion, serviceDefinition, provisioningConfigFile, configStoreNamespace)
	if done {
		return i
	}

	s.Logger.Debug(fmt.Sprintf("%s: %s : %s: %s:", serviceName, serviceVersion, *envName, *configStoreNamespace))
	s.Logger.Info("Initiating service deployment: " + serviceName + "@" + serviceVersion + " in " + *envName)
	serviceClient.DeployUnreleasedServiceStream(serviceDefinition, parsedProvisioningConfig, *envName, *configStoreNamespace)

	return 0
}

func (s *Service) deployReleasedService(envName *string, serviceName *string, serviceVersion *string,
	provisioningConfigFile *string, configStoreNamespace *string) int {

	parsedProvisioningConfig, i, done := s.validateDeployService(envName, *serviceName, *serviceVersion, nil, provisioningConfigFile, configStoreNamespace)
	if done {
		return i
	}
	dataForScalingConsent := map[string]interface{}{
		"env_name":        *envName,
		"service_version": *serviceVersion,
		"action":          "released_service_deploy",
		"config":          parsedProvisioningConfig,
	}
	scalingConsent := s.askForScalingConsent(serviceName, dataForScalingConsent)
	if scalingConsent == 1 {
		return 1
	}
	s.Logger.Debug(fmt.Sprintf("%s: %s : %s: %s:", *serviceName, *serviceVersion, *envName, *configStoreNamespace))
	s.Logger.Info("Initiating service deployment: " + *serviceName + "@" + *serviceVersion + " in " + *envName)
	serviceClient.DeployReleasedServiceStream(*serviceName, *serviceVersion, *envName, *configStoreNamespace, parsedProvisioningConfig)

	return 0
}

/*
validateDeployService

	returns parsedProvisioningConfig: interface{}, exitCode int, toExit bool
*/
func (s *Service) validateDeployService(envName *string, serviceName string, serviceVersion string, serviceDefinition map[string]interface{}, provisioningConfigFile *string, configStoreNamespace *string) (interface{}, int, bool) {
	envServices, err := envClient.DescribeEnv(*envName, "", "")

	if err != nil {
		s.Logger.Error(err.Error())
		return nil, 1, true
	}

	envService := service.Service{}

	forceService := false

	for _, curService := range envServices.Services {
		if curService.Name == serviceName {
			forceService = true
			envService = curService
			break
		}
	}

	var parsedProvisioningConfig interface{}

	if len(*provisioningConfigFile) > 0 {
		parsedConfig, err := parseFile(*provisioningConfigFile)
		if err != nil {
			s.Logger.Error("Error while parsing provisioning file " + *provisioningConfigFile + " : " + err.Error())
			return nil, 1, true
		}
		parsedProvisioningConfig = parsedConfig
	}

	if forceService {
		diff, err := serviceClient.CompareService(envName, serviceName, serviceVersion, serviceDefinition, parsedProvisioningConfig, configStoreNamespace)
		if err != nil {
			s.Logger.Error(err.Error())
			return nil, 1, true
		}
		if diff != "" {
			s.Logger.Info(fmt.Sprintf("service: %s already exists in the env with version: %s\n", serviceName, envService.Version))
			message := fmt.Sprintf("Below changes will happen after this deployement\n\n%s\nDo you Accept? [Y/n]: ", diff)

			allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
			val, err := s.Input.AskWithConstraints(message, allowedInputs)

			if err != nil {
				s.Logger.Error(err.Error())
				return nil, 1, true
			}

			if val != "Y" {
				s.Logger.Info("Skipping force service deploy")
				return nil, 0, true
			}
		}
	}

	return parsedProvisioningConfig, 0, false
}

func parseFile(filePath string) (interface{}, error) {
	if len(filePath) != 0 {
		var parsedDefinition interface{}

		fileDefinition, err := file.Read(filePath)
		if err != nil {
			return parsedDefinition, errors.New("file does not exist")
		}

		if strings.Contains(filePath, ".yaml") || strings.Contains(filePath, ".yml") {
			err = yaml.Unmarshal(fileDefinition, &parsedDefinition)
			if err != nil {
				return parsedDefinition, errors.New("Unable to parse YAML. " + err.Error())
			}
		} else if strings.Contains(filePath, ".json") {
			err = json.Unmarshal(fileDefinition, &parsedDefinition)
			if err != nil {
				return parsedDefinition, errors.New("Unable to parse JSON. " + err.Error())
			}
		} else {
			return parsedDefinition, errors.New("unrecognized file format")
		}
		return parsedDefinition, nil
	}
	return errors.New("filepath cannot be empty"), nil
}

// Help : returns an explanatory string
func (s *Service) Help() string {
	if s.Release {
		return commandHelper("release", "service", "", []Options{
			{Flag: "--path", Description: "path to directory containing service definition and provisioning config"},
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

	if s.Unlabel {
		return commandHelper("unlabel", "service", "", []Options{
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
			{Flag: "--file", Description: "file to read service definition."},
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

	if s.Operate {
		return commandHelper("operate", "service", "", []Options{
			{Flag: "--name", Description: "name of service"},
			{Flag: "--env", Description: "name of environment"},
			{Flag: "--operation", Description: "name of the operation to be performed on the service"},
			{Flag: "--file", Description: "path of the file which contains the options for the operation in JSON format"},
			{Flag: "--options", Description: "options for the operation in JSON format"},
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

	if s.Unlabel {
		return "unlabel a service version"
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

	if s.Operate {
		return "perform operations on a service"
	}

	return defaultHelper()
}
