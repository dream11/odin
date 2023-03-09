package commands

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
)

var operationClient backend.Operation

type Operation command

func (o *Operation) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	name := flagSet.String("name", "", "name of the operation")
	componentType := flagSet.String("component-type", "", "component-type on which operations will be performed")

	err := flagSet.Parse(args)
	if err != nil {
		o.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if o.List {
		emptyParameters := emptyParameters(map[string]string{"--component-type": *componentType})
		if len(emptyParameters) == 0 {
			operationList, err := operationClient.ListOperations(*componentType)
			if err != nil {
				o.Logger.Error(err.Error())
				return 1
			}

			o.Logger.Info("Listing all operation(s)")
			tableHeaders := []string{"Name", "Descrption"}
			var tableData [][]interface{}

			for _, operation := range operationList {
				tableData = append(tableData, []interface{}{
					operation.Name,
					operation.Description,
				})
			}
			table.Write(tableHeaders, tableData)

			return 0
		}

		o.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	if o.Describe {
		emptyParameters := emptyParameters(map[string]string{"--name": *name, "--component-type": *componentType})
		if len(emptyParameters) == 0 {
			operationList, err := operationClient.ListOperations(*componentType)
			if err != nil {
				o.Logger.Error(err.Error())
				return 1
			}

			o.Logger.Info("Describing operation: " + *name + " on component " + *componentType)
			var operationKeys interface{}

			for i := range operationList {
				if operationList[i].Name == *name {
					operationKeys = operationList[i].InputSchema
					break
				}
			}

			if operationKeys == nil {
				o.Logger.Error(fmt.Sprintf("operation: %s does not exist for the component: %s", *name, *componentType))
				return 1
			}

			operationKeysJson, err := json.MarshalIndent(operationKeys, "", "  ")
			if err != nil {
				o.Logger.Error(err.Error())
				return 1
			}

			o.Logger.Output(fmt.Sprintf("\n%s", operationKeysJson))
			return 0
		}

		o.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}
	o.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (o *Operation) Help() string {
	if o.List {
		return commandHelper("list", "operation", "", []Options{
			{Flag: "--component-type", Description: "component-type on which operations will be performed"},
		})
	}
	if o.Describe {
		return commandHelper("describe", "operation", "", []Options{
			{Flag: "--name", Description: "name of the operation"},
			{Flag: "--component-type", Description: "component-type on which operations will be performed"},
		})
	}
	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (o *Operation) Synopsis() string {
	if o.List {
		return "list all operations on a component-type"
	}
	if o.Describe {
		return "describe operation on a component-type"
	}
	return defaultHelper()
}
