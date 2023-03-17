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

	err := flagSet.Parse(args)
	if err != nil {
		c.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if c.Operate {
		if *envName == "" {
			*envName = utils.FetchKey(ENV_NAME_KEY)
		}
		emptyParameters := emptyParameters(map[string]string{"--name": *name, "--service": *serviceName, "--env": *envName, "--operation": *operation, "--options": *options})
		if len(emptyParameters) == 0 {
			var optionsJson interface{}
			err = json.Unmarshal([]byte(*options), &optionsJson)
			if err != nil {
				c.Logger.Error("Unable to parse options JSON " + err.Error())
				return 1
			}

			envTypeResp, err := envTypeClient.GetEnvType(*envName)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}
			if envTypeResp.Strict {
				consentMessage := "\nYou are executing the above command on production environment. Are you sure? Enter Y/n: "
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

			data := component.OperateComponentRequest{
				Data: component.Data{
					EnvName:     *envName,
					ServiceName: *serviceName,
					Operations: []component.Operation{
						{
							Name:   *operation,
							Values: optionsJson,
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
