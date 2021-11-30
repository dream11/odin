package commands

import (
	"github.com/dream11/odin/internal/backend"
)

// initiate backend client for component
var componentClient backend.Component

// Component : command declaration
type Component command

// Run : implements the actual functionality of the command
func (c *Component) Run(args []string) int {
	if c.List {
		c.Logger.Info("Listing all components")
		componentTypeList, err := componentClient.ListComponents()
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}

		for _, component := range componentTypeList {
			c.Logger.Output(component.Name)
		}

		return 0
	}

	c.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (c *Component) Help() string {
	if c.List {
		return commandHelper("list", "component", []string{})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *Component) Synopsis() string {
	if c.List {
		return "list all components"
	}

	return defaultHelper()
}
