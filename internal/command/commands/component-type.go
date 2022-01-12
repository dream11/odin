package commands

import (
)


// Component : command declaration
type ComponentType command

// Run : implements the actual functionality of the command
func (c *ComponentType) Run(args []string) int {
	if c.List {
		c.Logger.Info("Listing all components types")

		return 0
	}

	c.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (c *ComponentType) Help() string {
	if c.List {
		return commandHelper("list", "component-type", []string{})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *ComponentType) Synopsis() string {
	if c.List {
		return "list all components types"
	}

	return defaultHelper()
}
