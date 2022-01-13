package commands

import (
	"flag"
	"github.com/dream11/odin/internal/backend"
	"gopkg.in/yaml.v3"
)

// initiate backend client for component
var componentClient backend.ComponentType

// Component : command declaration
type Component command

// Run : implements the actual functionality of the command
func (c *Component) Run(args []string) int {
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	componentName := flagSet.String("name", "", "name of component")
	componentVersion := flagSet.String("version", "", "version of component")

	err := flagSet.Parse(args)
	if err != nil {
		c.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if c.Describe {
		if emptyParameterValidation([]string{*componentName, *componentVersion}) {
			c.Logger.Info("Describing component: " + *componentName + "@" + *componentVersion)
			componentResp, err := componentClient.DescribeComponent(*componentName, *componentVersion)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}

			details, err := yaml.Marshal(componentResp)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}

			c.Logger.Output(string(details))

			return 0
		}

		c.Logger.Error("component name & version cannot be blank")
		return 1
	}
	c.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (c *Component) Help() string {
	if c.Describe {
		return commandHelper("describe", "component", []string{
			"--name=name of component (required)",
			"--version=of component (required)",})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *Component) Synopsis() string {
	if c.Describe {
		return "describe a service component"
	}

	return defaultHelper()
}
