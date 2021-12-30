package commands

import (
	"encoding/json"
	"flag"
	"strings"

	"github.com/dream11/odin/api/infra"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/file"
	"github.com/dream11/odin/pkg/table"
	"gopkg.in/yaml.v3"
)

// initiate backend client for infra
var infraClient backend.Infra

// Infra : command declaration
type Infra command

// Run : implements the actual functionality of the command
func (i *Infra) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	name := flagSet.String("name", "", "name of environment")
	team := flagSet.String("team", "", "display environments created by a team")
	purpose := flagSet.String("purpose", "", "reason to create infra")
	envType := flagSet.String("envType", "", "envType to attach with infra")
	providerAccount := flagSet.String("account", "", "account name to provision the infra in")
	filePath := flagSet.String("file", "infra.yaml", "file to read infra config")
	detailed := flagSet.Bool("detailed", false, "get detailed view")

	err := flagSet.Parse(args)
	if err != nil {
		i.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if i.Create {
		if emptyParameterValidation([]string{*envType}) {
			i.Logger.Warn("Creating env in  " + *envType)

			infraConfig := infra.Infra{
				Team:    *team,
				Purpose: *purpose,
				Env:     *envType,
				Account: *providerAccount,
			}

			response, err := infraClient.CreateInfra(infraConfig)
			if err != nil {
				i.Logger.Error(err.Error())
				return 1
			}

			i.Logger.Success("Infra: " + response.Name + " created!")

			return 0
		}

		i.Logger.Error("envType cannot be blank")
		return 1
	}

	if i.Update {
		if emptyParameterValidation([]string{*name}) {
			i.Logger.Warn("Updating infra: " + *name)

			configData, err := file.Read(*filePath)
			if err != nil {
				i.Logger.Error("Unable to read from " + *filePath + "\n" + err.Error())
				return 1
			}

			var parsedConfig interface{}

			if strings.Contains(*filePath, ".yaml") || strings.Contains(*filePath, ".yml") {
				err = yaml.Unmarshal(configData, &parsedConfig)
				if err != nil {
					i.Logger.Error("Unable to parse YAML. " + err.Error())
					return 1
				}
			} else if strings.Contains(*filePath, ".json") {
				err = json.Unmarshal(configData, &parsedConfig)
				if err != nil {
					i.Logger.Error("Unable to parse JSON. " + err.Error())
					return 1
				}
			} else {
				i.Logger.Error("Unrecognized file format")
				return 1
			}

			infraClient.UpdateInfra(*name, parsedConfig)

			return 0
		}

		i.Logger.Error("infra name cannot be blank")
		return 1
	}

	if i.Describe {
		if emptyParameterValidation([]string{*name}) {
			i.Logger.Info("Describing " + *name)
			infraResp, err := infraClient.DescribeInfra(*name)
			if err != nil {
				i.Logger.Error(err.Error())
				return 1
			}

			for _, infra := range infraResp {
				i.Logger.Info(infra.Name + " details!")
				details, err := yaml.Marshal(infra)
				if err != nil {
					i.Logger.Error(err.Error())
					return 1
				}

				i.Logger.Output(string(details))
			}

			return 0
		}

		i.Logger.Error("infra name cannot be blank")
		return 1
	}

	if i.List {
		i.Logger.Info("Listing all infra(s)")
		infraList, err := infraClient.ListInfra()
		if err != nil {
			i.Logger.Error(err.Error())
			return 1
		}

		if *detailed {
			for _, inf := range infraList {
				i.Logger.Info("Infra definition for: " + inf.Name)

				infraYaml, err := yaml.Marshal(inf)
				if err != nil {
					i.Logger.Error("Unable to parse infra definition! " + err.Error())
					return 1
				}

				i.Logger.Output(string(infraYaml))
			}
		} else {
			tableHeaders := []string{"Name", "Purpose", "Team", "Env", "State", "Account", "Deletion Time"}
			var tableData [][]interface{}

			for _, inf := range infraList {
				tableData = append(tableData, []interface{}{
					inf.Name,
					inf.Purpose,
					inf.Team,
					inf.Env,
					inf.State,
					inf.Account,
					inf.DeletionTime,
				})
			}

			err = table.Write(tableHeaders, tableData)
			if err != nil {
				i.Logger.Error(err.Error())
				return 1
			}
		}

		return 0
	}

	if i.Delete {
		if emptyParameterValidation([]string{*name}) {
			i.Logger.Warn("Deleting infra:" + *name)
			infraClient.DeleteInfra(*name)

			return 0
		}

		i.Logger.Error("infra name cannot be blank")
		return 1
	}

	i.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (i *Infra) Help() string {
	if i.Create {
		return commandHelper("create", "infra", []string{
			"--team=team name to associate the infra with",
			"--purpose=reason to create infra",
			"--env=env to create infra in",
			"--account=account name to provision the infra in (optional)",
		})
	}

	if i.Update {
		return commandHelper("update", "infra", []string{
			"--name=name of infra to update",
			"--file=file path to pick update config",
		})
	}

	if i.List {
		return commandHelper("list", "infra", []string{
			"--detailed (get a detailed view)",
		})
	}

	if i.Describe {
		return commandHelper("describe", "infra", []string{
			"--name=name of infra to describe",
		})
	}

	if i.Delete {
		return commandHelper("delete", "infra", []string{
			"--name=name of environment to delete",
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (i *Infra) Synopsis() string {
	if i.Create {
		return "create an infra"
	}

	if i.Update {
		return "update an infra"
	}

	if i.List {
		return "list all active infra"
	}

	if i.Describe {
		return "describe an infra"
	}

	if i.Delete {
		return "delete an infra"
	}

	return defaultHelper()
}
