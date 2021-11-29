package commands

import (
	"encoding/json"
	"flag"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
)

// initiate backend client for profile
var profileClient backend.Profile

// Profile : command declaration
type Profile command

// Run : implements the actual functionality of the command
func (p *Profile) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)

	// create flags
	filePath := flagSet.String("file", "profile.yaml", "file name of profile yaml")
	profileName := flagSet.String("name", "nil", "name of profile to be used")
	profileVersion := flagSet.String("version", "nil", "version of profile to be used")
	envName := flagSet.String("env", "nil", "name of environment to use")
	infraName := flagSet.String("infra", "nil", "name of infra to deploy the profile in")
	teamName := flagSet.String("team", "", "name of user's team")

	// positional parse flags from [3:]
	err := flagSet.Parse(os.Args[3:])
	if err != nil {
		p.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if p.Create {
		configData, err := file.Read(*filePath)
		if err != nil {
			p.Logger.Error("Unable to read from " + *filePath + "\n" + err.Error())
			return 1
		}

		var parsedConfig interface{}

		if strings.Contains(*filePath, ".yaml") || strings.Contains(*filePath, ".yml") {
			err = yaml.Unmarshal(configData, &parsedConfig)
			if err != nil {
				p.Logger.Error("Unable to parse YAML. " + err.Error())
				return 1
			}
		} else if strings.Contains(*filePath, ".json") {
			err = json.Unmarshal(configData, &parsedConfig)
			if err != nil {
				p.Logger.Error("Unable to parse JSON. " + err.Error())
				return 1
			}
		} else {
			p.Logger.Error("Unrecognized file format")
			return 1
		}

		// TODO: validate no conversion required
		//configJson, err := json.Marshal(parsedConfig)
		//if err != nil {
		//	p.Logger.Error("Unable to translate config to Json. " + err.Error())
		//	return 1
		//}

		// TODO: validate request
		profileClient.CreateProfile(parsedConfig)

		return 0
	}

	if p.Describe {
		p.Logger.Info("Describing profile: " + *profileName + "@" + *profileVersion)
		// TODO: validate request & receive parsed input to display
		profileClient.DescribeProfile(*profileName, *profileVersion)

		return 0
	}

	if p.List {
		p.Logger.Info("Listing all profiles")
		// TODO: validate request & receive parsed input to display
		profileClient.ListProfiles(*teamName, *profileVersion)

		return 0
	}

	if p.Deploy {
		p.Logger.Info("Deploying profile: " + *profileName + "@" + *profileVersion + " in " + *envName + "/" + *infraName)
		// TODO: call PG api deploys a profile version in given env
		// POST /deploy?profile=<profile>&version=<version>&env=<env>

		return 0

	}

	if p.Destroy {
		p.Logger.Info("Destroying profile: " + *profileName + "@" + *profileVersion + " in " + *envName + "/" + *infraName)
		// TODO: call PG api that destroys a profile version from given env
		// DELETE /deploy?profile=<profile>&version=<version>&env=<env>

		return 0

	}

	if p.Status {
		p.Logger.Info("Fetching status for profile: " + *profileName + " in " + *envName + "/" + *infraName)
		// TODO: call PG api that returns status of profile in env
		// GET /profileStatus?profile=<profile>&env=<env>

		return 0
	}

	if p.Logs {
		p.Logger.Info("Fetching logs for profile: " + *profileName + " in " + *envName + "/" + *infraName)
		// TODO: call PG api that returns execution logs of profile in env
		// GET /profileLogs?profile=<profile>&env=<env>

		return 0
	}

	if p.Delete {
		p.Logger.Warn("Deleting profile: " + *profileName + "@" + *profileVersion)
		// TODO: validate request
		profileClient.DeleteProfile(*profileName, *profileVersion)

		return 0
	}

	p.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (p *Profile) Help() string {
	if p.Create {
		return commandHelper("create", "profile", []string{
			"--file=yaml file to read profile properties from",
		})
	}

	if p.Describe {
		return commandHelper("describe", "profile", []string{
			"--name=name of profile to describe",
			"--version=version of profile to describe",
		})
	}

	if p.List {
		return commandHelper("list", "profile", []string{
			"--team=team name to list profiles for",
			"--version=version of profile to list",
		})
	}

	if p.Deploy {
		return commandHelper("deploy", "profile", []string{
			"--name=name of profile to deploy",
			"--version=version of profile to deploy",
			"--env=name of env to use",
			"--infra=name of infra to deploy the profile in",
		})
	}

	if p.Destroy {
		return commandHelper("destroy", "profile", []string{
			"--name=name of profile to destroy",
			"--version=version of profile to destroy",
			"--env=name of env to destroy the profile in",
		})
	}

	if p.Status {
		return commandHelper("status", "profile", []string{
			"--name=name of profile to get status",
			"--env=name of env to get status",
		})
	}

	if p.Logs {
		return commandHelper("logs", "profile", []string{
			"--name=name of profile to get logs",
			"--env=name of env to get logs",
		})
	}

	if p.Delete {
		return commandHelper("delete", "profile", []string{
			"--name=name of profile to delete",
			"--version=version of profile to delete",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (p *Profile) Synopsis() string {
	if p.Create {
		return "create a profile"
	}

	if p.Describe {
		return "describe a profile"
	}

	if p.List {
		return "list all active profiles"
	}

	if p.Deploy {
		return "deploy a profile"
	}

	if p.Destroy {
		return "destroy a profile"
	}

	if p.Status {
		return "current status of a profile"
	}

	if p.Logs {
		return "execution logs for profile"
	}

	if p.Delete {
		return "delete a profile version"
	}

	return defaultHelper()
}
