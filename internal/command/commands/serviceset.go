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
	filePath := flagSet.String("file", "service-set.yaml", "file to read service-set config")
	serviceSetName := flagSet.String("name", "", "name of service-set to be used")
	serviceName := flagSet.String("service", "", "name of service in service-set")
	envName := flagSet.String("env", "", "name of environment to deploy the service-set in")
	platform := flagSet.String("platform", "", "platform of environment to deploy the service-set in")
	force := flagSet.Bool("force", false, "forcefully deploy service-set into the Env")

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

		serviceSetClient.CreateServiceSet(parsedConfig)
		s.Logger.Info(fmt.Sprintf("ServiceSet: %s created Successfully.", serviceDataMap["name"].(string)))
		return 0
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

		err = table.Write(tableHeaders, tableData)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		s.Logger.Output("\nCommand to describe serviceset")
		s.Logger.ItalicEmphasize("odin describe serviceset --name <serviceSetName>")
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
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceSetName, "--env": *envName})
		if len(emptyParameters) == 0 {
			var forceDeployServices []serviceset.ListEnvService
			if !*force {
				//get list of env services
				s.Logger.Debug(fmt.Sprintf("Env Services of service-set %s and env %s", *serviceSetName, *envName))
				serviceSetList, err := serviceSetClient.ListEnvServices(*serviceSetName, *envName, "conflictedVersion")

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if len(serviceSetList) > 0 {
					s.Logger.Output("Following services have conflicting versions in the Env: " + *envName)
					s.Logger.Output("Press [Y] to update the service version or press [n] to skip service.\n")
					allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
					for _, serviceSet := range serviceSetList {
						message := fmt.Sprintf("Update version of Service %s : %s -> %s[Y/n]: ", serviceSet.Name, serviceSet.EnvVersion, serviceSet.Version)
						//s.Logger.Output(message)

						val, err := s.Input.AskWithConstraints(message, allowedInputs)

						if err != nil {
							s.Logger.Error(err.Error())
							return 1
						}

						s.Logger.Output(val)
						if val == "Y" {
							forceDeployServices = append(forceDeployServices, serviceSet)
						}
					}
				}
			}

			/*deploy service-set*/
			s.Logger.Debug("Deploying service-set: " + *serviceSetName + " in " + *envName)
			serviceSetList, err := serviceSetClient.DeployServiceSet(*serviceSetName, *envName, *platform, forceDeployServices, *force)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"Name", "Version", "ExecutorUrl", "Error"}
			var tableData [][]interface{}

			for _, serviceSet := range serviceSetList {
				tableData = append(tableData, []interface{}{
					serviceSet.Name,
					serviceSet.Version,
					serviceSet.ExecutorUrl,
					serviceSet.Error,
				})
			}

			s.Logger.Success(fmt.Sprintf("Deployment of service-set %s is started on env %s\n", *serviceSetName, *envName))
			err = table.Write(tableHeaders, tableData)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.Undeploy {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceSetName, "--env": *envName})
		if len(emptyParameters) == 0 {
			var forceUndeployServices []serviceset.ListEnvService
			if !*force {
				serviceSetList, err := serviceSetClient.ListEnvServices(*serviceSetName, *envName, "conflictedVersion")

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if len(serviceSetList) > 0 {
					s.Logger.Output("Following services have conflicting versions in the Env: " + *envName)
					s.Logger.Output("Press [Y] to undeploy the service with the conflicting version or press [n] to skip service.\n")
					allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
					for _, serviceSet := range serviceSetList {
						message := fmt.Sprintf("undeploy Service: %s with version: %s[Y/n]: ", serviceSet.Name, serviceSet.EnvVersion)
						val, err := s.Input.AskWithConstraints(message, allowedInputs)

						if err != nil {
							s.Logger.Error(err.Error())
							return 1
						}

						if val == "Y" {
							forceUndeployServices = append(forceUndeployServices, serviceSet)
						}
					}
				}

				fmt.Println(forceUndeployServices)
			}

			/*deploy service-set*/
			s.Logger.Info("Undeploying service-set: " + *serviceSetName + " in Env:" + *envName)
			serviceSetList, err := serviceSetClient.UndeployServiceSet(*serviceSetName, *envName, forceUndeployServices, *force)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"Name", "Version", "ExecutorUrl", "Error"}
			var tableData [][]interface{}

			for _, serviceSet := range serviceSetList {
				tableData = append(tableData, []interface{}{
					serviceSet.Name,
					serviceSet.Version,
					serviceSet.ExecutorUrl,
					serviceSet.Error,
				})
			}

			s.Logger.Success(fmt.Sprintf("Undeployment of service-set %s is started on env %s\n", *serviceSetName, *envName))
			err = table.Write(tableHeaders, tableData)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			return 0
		}
		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.Update {
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

		serviceSetClient.UpdateServiceSet(serviceDataMap["name"].(string), parsedConfig)
		s.Logger.Info(fmt.Sprintf("ServiceSet: %s updated Successfully.", serviceDataMap["name"].(string)))

		return 0
	}

	s.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (s *ServiceSet) Help() string {
	if s.Create {
		return commandHelper("create", "service-set", []string{
			"--file=yaml file to read service-set definition",
		})
	}

	if s.List {
		return commandHelper("list", "service-set", []string{
			"--name=name of the service-set",
			"--service=name of service in the service-set",
		})
	}

	if s.Describe {
		return commandHelper("describe", "service", []string{
			"--name=name of the service-set to describe",
		})
	}

	if s.Delete {
		return commandHelper("delete", "service-set", []string{
			"--name=name of service-set to delete",
		})
	}

	if s.Deploy {
		return commandHelper("deploy", "service-set", []string{
			"--name=name of service-set to deploy",
			"--env=name of environment to deploy service-set in",
			"--platform=platform of environment to deploy service-set in",
			"--force=forcefully deploy service-set into the Env",
		})
	}

	if s.Undeploy {
		return commandHelper("deploy", "service-set", []string{
			"--name=name of service-set to deploy",
			"--env=name of environment to deploy service-set in",
			"--force=forcefully undeploy service-set from the Env",
		})
	}

	if s.Update {
		return commandHelper("update", "service-set", []string{
			"--file=yaml file to read service-set definition",
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

	if s.Deploy {
		return "undeploy a service-set"
	}

	if s.Update {
		return "update a service-set"
	}

	return defaultHelper()
}
