package commands

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dream11/odin/api/component"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/utils"
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
					c.Logger.Info("\nAborting...")
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

			componentClient.OperateComponent(*name, data)
			return 0
		}

		c.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	c.Logger.Error("Not a valid command")
	return 127
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
