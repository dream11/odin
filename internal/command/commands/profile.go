package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/dream11/odin/api/profile"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"gopkg.in/yaml.v3"
)

// initiate backend client for profile
var profileClient backend.Profile

// Profile : command declaration
type Profile command

// Run : implements the actual functionality of the command
func (s *Profile) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	filePath := flagSet.String("file", "profile.yaml", "file to read profile config")
	profileName := flagSet.String("name", "", "name of profile to be used")
	serviceName := flagSet.String("service", "", "name of service in profile")
	envName := flagSet.String("env", "", "name of environment to deploy the profile in")
	platform := flagSet.String("platform", "", "platform of environment to deploy the profile in")
	force := flagSet.Bool("force", false, "forcefully deploy the new version of the service")

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

		s.Logger.Info(fmt.Sprintf("Profile creation started for %s  ", serviceDataMap["services"]))
		profileResp, err := profileClient.CreateProfile(parsedConfig)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		s.Logger.Success(profileResp)
		return 0
	}

	if s.List {
		s.Logger.Info("Listing all profiles")
		profileList, err := profileClient.ListProfiles(*profileName, *serviceName)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name"}
		var tableData [][]interface{}

		for _, profile := range profileList {
			tableData = append(tableData, []interface{}{
				profile.Name,
			})
		}

		err = table.Write(tableHeaders, tableData)
		if err != nil {
			s.Logger.Error(err.Error())
			return 1
		}
		s.Logger.Output("\nCommand to describe profile")
		s.Logger.ItalicEmphasize("odin describe profile --name <profileName>")
		return 0
	}

	if s.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *profileName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Describing profile: " + *profileName)
			profileResp, err := profileClient.DescribeProfile(*profileName)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			var details []byte
			s.Logger.Info(profileResp.Name + " details!")
			details, err = yaml.Marshal(profileResp)

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
		emptyParameters := emptyParameters(map[string]string{"--name": *profileName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Deleting profile: " + *profileName)
			profileClient.DeleteProfile(*profileName)

			return 0
		}

		s.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if s.Deploy {
		emptyParameters := emptyParameters(map[string]string{"--name": *profileName, "--env": *envName})
		if len(emptyParameters) == 0 {
			var forceDeployServices []profile.ListEnvService
			if !*force {
				//get list of env services
				s.Logger.Info(fmt.Sprintf("Env Services of profile %s and env %s", *profileName, *envName))
				profileList, err := profileClient.ListEnvServices(*profileName, *envName, "conflictedVersion")

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if len(profileList) > 0 {
					s.Logger.Output("Following services have conflicting versions in the Env: " + *envName)
					s.Logger.Output("Press [Y] to update the service version or press [n] to skip service.\n")
					allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
					for _, profile := range profileList {
						message := fmt.Sprintf("Update version of Service %s : %s -> %s[Y/n]: ", profile.Name, profile.EnvVersion, profile.Version)

						val, err := s.Input.AskWithConstraints(message, allowedInputs)

						if err != nil {
							s.Logger.Error(err.Error())
							return 1
						}

						s.Logger.Output(val)
						if val == "Y" {
							forceDeployServices = append(forceDeployServices, profile)
						}
					}
				}

				fmt.Println(forceDeployServices)
			}

			/*deploy profile*/
			s.Logger.Info("Deploying profile: " + *profileName + " in " + *envName)
			profileList, err := profileClient.DeployProfile(*profileName, *envName, *platform, forceDeployServices, *force)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"Name", "Version", "ExecutorUrl", "Error"}
			var tableData [][]interface{}

			for _, profile := range profileList {
				tableData = append(tableData, []interface{}{
					profile.Name,
					profile.Version,
					profile.ExecutorUrl,
					profile.Error,
				})
			}

			s.Logger.Success(fmt.Sprintf("Deployment of profile %s is started on env %s\n", *profileName, *envName))
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
		emptyParameters := emptyParameters(map[string]string{"--name": *profileName, "--env": *envName})
		if len(emptyParameters) == 0 {
			s.Logger.Info("Undeploying profile: " + *profileName + " from environment " + *envName)
			var forceUndeployServices []profile.ListEnvService
			if !*force {
				s.Logger.Info(fmt.Sprintf("Profile: %s services present in the Env: %s are", *profileName, *envName))
				profileList, err := profileClient.ListEnvServices(*profileName, *envName, "conflictedVersion")

				if err != nil {
					s.Logger.Error(err.Error())
					return 1
				}

				if len(profileList) > 0 {
					s.Logger.Output("Following services have conflicting versions in the Env: " + *envName)
					s.Logger.Output("Press [Y] to undeploy the service with the conflicting version or press [n] to skip service.\n")
					allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
					for _, profile := range profileList {
						message := fmt.Sprintf("undeploy Service: %s with version: %s[Y/n]: ", profile.Name, profile.EnvVersion)
						val, err := s.Input.AskWithConstraints(message, allowedInputs)

						if err != nil {
							s.Logger.Error(err.Error())
							return 1
						}

						if val == "Y" {
							forceUndeployServices = append(forceUndeployServices, profile)
						}
					}
				}

				fmt.Println(forceUndeployServices)
			}

			/*deploy profile*/
			s.Logger.Info("Undeploying profile: " + *profileName + " in Env:" + *envName)
			profileList, err := profileClient.UndeployProfile(*profileName, *envName, forceUndeployServices, *force)
			if err != nil {
				s.Logger.Error(err.Error())
				return 1
			}

			tableHeaders := []string{"Name", "Version", "ExecutorUrl", "Error"}
			var tableData [][]interface{}

			for _, profile := range profileList {
				tableData = append(tableData, []interface{}{
					profile.Name,
					profile.Version,
					profile.ExecutorUrl,
					profile.Error,
				})
			}

			s.Logger.Success(fmt.Sprintf("Undeployment of profile %s is started on env %s\n", *profileName, *envName))
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

	s.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (s *Profile) Help() string {
	if s.Create {
		return commandHelper("create", "profile", []string{
			"--file=yaml file to read profile definition",
		})
	}

	if s.List {
		return commandHelper("list", "profile", []string{
			"--name=name of the profile",
			"--service=name of service in the profile",
		})
	}

	if s.Describe {
		return commandHelper("describe", "service", []string{
			"--name=name of the profile to describe",
		})
	}

	if s.Delete {
		return commandHelper("delete", "profile", []string{
			"--name=name of profile to delete",
		})
	}

	if s.Deploy {
		return commandHelper("deploy", "profile", []string{
			"--name=name of profile to deploy",
			"--env=name of environment to deploy profile in",
			"--platform=platform of environment to deploy profile in",
			"--force=forcefully deploy your profile service into the env",
		})
	}

	if s.Undeploy {
		return commandHelper("deploy", "profile", []string{
			"--name=name of profile to deploy",
			"--env=name of environment to deploy profile in",
			"--force=forcefully deploy your profile service into the env",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (s *Profile) Synopsis() string {
	if s.Create {
		return "create a profile"
	}

	if s.List {
		return "list all profiles"
	}

	if s.Describe {
		return "describe a profile"
	}

	if s.Delete {
		return "delete a profile"
	}

	if s.Deploy {
		return "deploy a profile"
	}

	if s.Deploy {
		return "undeploy a profile"
	}

	return defaultHelper()
}
