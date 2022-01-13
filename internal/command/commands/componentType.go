package commands

import (
	"flag"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
	"gopkg.in/yaml.v3"
)

// initiate backend client for component type
var componentTypeClient backend.ComponentType

// Component Type : command declaration
type ComponentType command

// Run : implements the actual functionality of the command
func (c *ComponentType) Run(args []string) int {
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	componentTypeName := flagSet.String("name", "", "name of component type")
	componentTypeVersion := flagSet.String("version", "", "version of component type")

	err := flagSet.Parse(args)
	if err != nil {
		c.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}
	if c.List {
		c.Logger.Info("Listing all component types")
		componentTypeList, err := componentTypeClient.ListComponentTypes(*componentTypeName, *componentTypeVersion)
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}
		var tableHeaders []string
		var tableData [][]interface{}
		if len(*componentTypeName) == 0 {
			tableHeaders = []string{"Component Name", "Latest Version", "Total Versions Available"}
			for _, componentType := range componentTypeList {
				tableData = append(tableData, []interface{}{
					componentType.Name,
					componentType.Version,
					componentType.TotalVersions,
				})
			}
		} else {
			tableHeaders = []string{"Component Name", "Version"}
			for _, componentType := range componentTypeList {
				tableData = append(tableData, []interface{}{
					componentType.Name,
					componentType.Version,
				})
			}
		}

		err = table.Write(tableHeaders, tableData)
		if err != nil {
			c.Logger.Error(err.Error())
			return 1
		}
		c.Logger.Output("Got component type list. You can now describe component type using\nodin describe component-type --name <componentName> --version <componentVersion>")
		return 0
	}

	if c.Describe {
		if emptyParameterValidation([]string{*componentTypeName}) {
			c.Logger.Info("Describing component type: " + *componentTypeName + "@" + *componentTypeVersion)
			componentTypeResp, err := componentTypeClient.DescribeComponentType(*componentTypeName, *componentTypeVersion)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}

			details, err := yaml.Marshal(componentTypeResp)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}

			c.Logger.Output(string(details))

			return 0
		}

		c.Logger.Error("component type name cannot be blank")
		return 1
	}
	c.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (c *ComponentType) Help() string {
	if c.List {
		return commandHelper("list", "component-type", []string{
			"--name=name of component type",
			"--version=of component type"})
	}
	if c.Describe {
		return commandHelper("list", "component-type", []string{
			"--name=name of component type (required)",
			"--version=of component type (deafult latest)"})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *ComponentType) Synopsis() string {
	if c.List {
		return "list all components types"
	}
	if c.List {
		return "describe component type"
	}
	return defaultHelper()
}
