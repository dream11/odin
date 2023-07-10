package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/dream11/odin/api/component"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/utils"
	"github.com/fatih/color"
)

var componentClient backend.Component

type Component command

func (c *Component) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	name := flagSet.String("name", "", "name of the component")
	serviceName := flagSet.String("service", "", "name of the service in which the component is deployed")
	envName := flagSet.String("env", "", "name of the environment in which the service is deployed")
	operation := flagSet.String("operation", "", "name of the operation to performed on the component")
	options := flagSet.String("options", "", "options of the operation in JSON format")
	filePath := flagSet.String("file", "", "file to provide options for component operations")

	err := flagSet.Parse(args)
	if err != nil {
		c.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if c.Operate {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}
		emptyParameters := emptyParameters(map[string]string{"--name": *name, "--service": *serviceName, "--env": *envName, "--operation": *operation})
		if len(emptyParameters) == 0 {
			isOptionsPresent := len(*options) > 0
			isFilePresent := len(*filePath) > 0

			if isOptionsPresent && isFilePresent {
				c.Logger.Error("You can provide either --options or --file but not both")
				return 1
			}

			if !isOptionsPresent && !isFilePresent {
				c.Logger.Error("You should provide either --options or --file")
				return 1
			}

			var optionsData map[string]interface{}

			if isFilePresent {
				parsedConfig, err := parseFile(*filePath)
				if err != nil {
					c.Logger.Error("Error while parsing file " + *filePath + " : " + err.Error())
					return 1
				}
				optionsData = parsedConfig.(map[string]interface{})
			} else if isOptionsPresent {
				err = json.Unmarshal([]byte(*options), &optionsData)
				if err != nil {
					c.Logger.Error("Unable to parse JSON data " + err.Error())
					return 1
				}
			}

			data := component.OperateComponentRequest{
				Data: component.Data{
					EnvName:     *envName,
					ServiceName: *serviceName,
					Operations: []component.Operation{
						{
							Name:   *operation,
							Values: optionsData,
						},
					},
				},
			}

			diffValues, err := componentClient.CompareOperationChanges(*name, data)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}
			oldComponentValues := diffValues.OldValues
			newComponentValues := diffValues.NewValues

			if len(oldComponentValues) > 0 {
				c.Logger.Info("\nBelow changes will happen after this operation:")
				tableHeaders := []string{"Component Name", "Key", "Old Value", "New Value"}
				var tableData [][]interface{}

				componentName := *name
				flatendOldComponentValues := flattenMap(oldComponentValues, "")
				flatendNewComponentValues := flattenMap(newComponentValues, "")

				keys := make([]string, 0, len(flatendNewComponentValues))
				for k := range flatendOldComponentValues {
					keys = append(keys, k)
				}
				sort.Strings(keys)

				for _, key := range keys {
					oldValue := flatendOldComponentValues[key]
					newValue := flatendNewComponentValues[key]
					if fmt.Sprintf("%v", oldValue) != fmt.Sprintf("%v", newValue) {
						var oldValueString string
						var newValueString string

						switch oldValue := oldValue.(type) {
						case []interface{}:
							strSlice := make([]string, len(oldValue))
							for i, v := range oldValue {
								strSlice[i] = fmt.Sprintf("%v", v)
							}
							oldValueString = color.RedString("[" + strings.Join(strSlice, ", ") + "]")
						default:
							oldValueString = color.RedString(fmt.Sprintf("%v", oldValue))
						}

						switch newValue := newValue.(type) {
						case []interface{}:
							strSlice := make([]string, len(newValue))
							for i, v := range newValue {
								strSlice[i] = fmt.Sprintf("%v", v)
							}
							newValueString = color.GreenString("[" + strings.Join(strSlice, ", ") + "]")
						default:
							newValueString = color.GreenString(fmt.Sprintf("%v", flatendNewComponentValues[key]))
						}

						tableData = append(tableData, []interface{}{
							componentName,
							key,
							oldValueString,
							newValueString,
						})
					}
				}
				table.Write(tableHeaders, tableData)
			}

			envTypeResp, err := envTypeClient.GetEnvType(*envName)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}
			if envTypeResp.Strict {
				consentMessage := fmt.Sprintf("\nYou are executing the above command on a restricted environment. Are you sure? Enter \033[1m%s\033[0m to continue:", *envName)
				val, err := c.Input.Ask(consentMessage)

				if err != nil {
					c.Logger.Error(err.Error())
					return 1
				}

				if val != *envName {
					c.Logger.Info("Aborting the operation")
					return 1
				}
			} else {
				message := "\nDo you want to proceed with the above command? [Y/n]:"
				allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
				val, err := c.Input.AskWithConstraints(message, allowedInputs)

				if err != nil {
					c.Logger.Error(err.Error())
					return 1
				}

				if val != "Y" {
					c.Logger.Info("Aborting the operation")
					return 1
				}
			}

			dataForScalingConsent := map[string]interface{}{
				"env_name":       *envName,
				"component_name": *name,
				"action":         *operation,
				"config":         optionsData,
			}
			componentListResponse, err := serviceClient.ScalingServiceConsent(*serviceName, dataForScalingConsent)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}
			for _, component := range componentListResponse.Response {
				consentMessage := fmt.Sprintf("\nYou have enabled reactive scaling for %s, this means %s will no longer be scaled using Scaler post this operation. Do you wish to continue? [Y/n]:", component, component)
				allowedInputs := map[string]struct{}{"Y": {}, "n": {}}
				val, err := c.Input.AskWithConstraints(consentMessage, allowedInputs)

				if err != nil {
					c.Logger.Error(err.Error())
					return 1
				}

				if val != "Y" {
					c.Logger.Info("Aborting the operation")
					return 1
				}
			}

			componentClient.OperateComponent(*name, data)
			return 0
		}

		c.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	c.Logger.Error("Not a valid command")
	return 127
}

func flattenMap(m map[string]interface{}, prefix string) map[string]interface{} {
	flattened := make(map[string]interface{})
	for k, v := range m {
		key := prefix + k
		if prefix != "" {
			key = prefix + "." + k
		}
		if vm, ok := v.(map[string]interface{}); ok {
			flattenedMap := flattenMap(vm, key)
			if len(flattenedMap) == 0 {
				flattened[key] = make(map[string]interface{})
			}
			for fk, fv := range flattenedMap {
				flattened[fk] = fv
			}
		} else {
			flattened[key] = v
		}
	}
	return flattened
}

// Help : returns an explanatory string
func (c *Component) Help() string {
	if c.Operate {
		return commandHelper("operate", "component", "", []Options{
			{Flag: "--name", Description: "name of the component"},
			{Flag: "--service", Description: "name of the service in which the component is deployed"},
			{Flag: "--env", Description: "name of the environment in which the service is deployed"},
			{Flag: "--operation", Description: "name of the operation to performed on the component"},
			{Flag: "--options", Description: "options of the operation in JSON format"},
			{Flag: "--file", Description: "path of the file which contains the options for the operation in JSON format"},
		})
	}
	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *Component) Synopsis() string {
	if c.Operate {
		return "Operate on a component"
	}
	return defaultHelper()
}
