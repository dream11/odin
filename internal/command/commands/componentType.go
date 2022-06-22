package commands

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
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

		table.Write(tableHeaders, tableData)

		c.Logger.Output("\nCommand to describe component types")
		c.Logger.ItalicEmphasize("odin describe component-type --name <componentTypeName> --version <componentTypeVersion>")
		return 0
	}

	if c.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *componentTypeName})
		if len(emptyParameters) == 0 {
			c.Logger.Info("Describing component type: " + *componentTypeName + "@" + *componentTypeVersion)
			componentDetailsResponse, err := componentTypeClient.DescribeComponentType(*componentTypeName, *componentTypeVersion)
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}

			details, err := json.MarshalIndent(componentDetailsResponse.Details, "", "  ")
			if err != nil {
				c.Logger.Error(err.Error())
				return 1
			}

			var tableHeaders []string
			var tableData [][]interface{}
			if len(componentDetailsResponse.ExposedConfigs) > 0 {
				tableHeaders = []string{"Config", "Mandatory", "Data Type"}
				for _, exposed_config := range componentDetailsResponse.ExposedConfigs {
					tableData = append(tableData, []interface{}{
						exposed_config.Config,
						exposed_config.Mandatory,
						exposed_config.DataType,
					})
				}
			}

			c.Logger.Output(fmt.Sprintf("\n%s", details))
			c.Logger.ItalicEmphasize("\nList of exposed configs :\n")
			table.Write(tableHeaders, tableData)

			return 0
		}

		c.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))

		return 1
	}
	c.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (c *ComponentType) Help() string {
	if c.List {
		return commandHelper("list", "component-type", "", []Options{
			{Flag: "--name", Description: "name of component type"},
			{Flag: "--version", Description: "of component type"},
		})
	}
	if c.Describe {
		return commandHelper("list", "component-type", "", []Options{
			{Flag: "--name", Description: "name of component type (required)"},
			{Flag: "--version", Description: "of component type (deafult latest)"},
		})
	}
	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (c *ComponentType) Synopsis() string {
	if c.List {
		return "list all components types"
	}
	if c.Describe {
		return "describe component type"
	}
	return defaultHelper()
}
