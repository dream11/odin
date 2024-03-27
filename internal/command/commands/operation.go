package commands

import (
	"encoding/json"
	"flag"
	"fmt"

	operationapi "github.com/dream11/odin/api/operation"
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
	entity := flagSet.String("entity", "", "name of the entity [env|service|component]")

	err := flagSet.Parse(args)
	if err != nil {
		o.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	isComponentOperation := false
	isEnvOperation := false
	isServiceOperation := false

	if o.List {
		if o.parseFlags(entity, componentType, &isEnvOperation, &isComponentOperation, &isServiceOperation) == 1 {
			return 1
		}

		var operationList []operationapi.Operation
		var err error

		infoMsg := "Listing all " + *entity + " operations"
		outputMsg := "\nCommand to describe " + *entity + " operations"
		descibeCommandMsg := "odin describe operation --name <operationName> --entity " + *entity

		if isComponentOperation {
			operationList, err = operationClient.ListComponentTypeOperations(*componentType)
			infoMsg = "Listing all component operations on component " + *componentType
			descibeCommandMsg = "odin describe operation --name <operationName> --entity component --component-type <componentTypeName>"
		} else if isServiceOperation {
			operationList, err = operationClient.ListServiceOperations()
		} else if isEnvOperation {
			operationList, err = operationClient.ListEnvOperations()
		}

		if err != nil {
			o.Logger.Error(err.Error())
			return 1
		}

		o.Logger.Info(infoMsg)

		tableHeaders := []string{"Name", "Descrption"}
		var tableData [][]interface{}

		for _, operation := range operationList {
			tableData = append(tableData, []interface{}{
				operation.Name,
				operation.Description,
			})
		}
		table.Write(tableHeaders, tableData)

		o.Logger.Output(outputMsg)
		o.Logger.ItalicEmphasize(descibeCommandMsg)

		return 0
	}

	if o.Describe {
		isNamePresent := len(*name) > 0
		if !isNamePresent {
			o.Logger.Error("--name cannot be blank")
			return 1
		}

		if o.parseFlags(entity, componentType, &isEnvOperation, &isComponentOperation, &isServiceOperation) == 1 {
			return 1
		}

		infoMsg := "Describing the " + *entity + " operation: " + *name
		errorMsg := "operation: " + *name + " is not a valid " + *entity + " operation"

		var operationList []operationapi.Operation
		var err error

		if isComponentOperation {
			operationList, err = operationClient.ListComponentTypeOperations(*componentType)
			infoMsg = "Describing operation: " + *name + " on component " + *componentType
			errorMsg = "operation: " + *name + " does not exist for the component: " + *componentType
		} else if isServiceOperation {
			operationList, err = operationClient.ListServiceOperations()
		} else if isEnvOperation {
			operationList, err = operationClient.ListEnvOperations()
		}

		if err != nil {
			o.Logger.Error(err.Error())
			return 1
		}

		var operationKeys interface{}

		for i := range operationList {
			if operationList[i].Name == *name {
				operationKeys = operationList[i].InputSchema
				break
			}
		}

		if operationKeys == nil {
			o.Logger.Error(errorMsg)
			return 1
		}

		operationKeysJson, err := json.MarshalIndent(operationKeys, "", "  ")
		if err != nil {
			o.Logger.Error(err.Error())
			return 1
		}

		o.Logger.Info(infoMsg)
		o.Logger.Output(fmt.Sprintf("\n%s", operationKeysJson))
		return 0
	}
	o.Logger.Error("Not a valid command")
	return 127
}

func (o *Operation) parseFlags(entity, componentType *string, isEnvOperation, isComponentOperation, isServiceOperation *bool) int {
	isComponentTypePresent := len(*componentType) > 0
	isEntityPresent := len(*entity) > 0

	if isEntityPresent {
		if *entity == "component" {
			if !isComponentTypePresent {
				o.Logger.Error("--component-type cannot be blank when --entity is component")
				return 1
			}
			*isComponentOperation = true
		} else if isComponentTypePresent {
			o.Logger.Error("--component-type should be used only when --entity is component")
			return 1
		} else if *entity == "env" {
			*isEnvOperation = true
		} else if *entity == "service" {
			*isServiceOperation = true
		} else {
			o.Logger.Error("Unknown value for --entity. Use one of env|service|component")
			return 1
		}
	} else {
		if isComponentTypePresent {
			*isComponentOperation = true
			*entity = "component"
		} else {
			*isServiceOperation = true
			*entity = "service"
		}
	}

	return 0
}

// Help : returns an explanatory string
func (o *Operation) Help() string {
	if o.List {
		return commandHelper("list", "operation", "", []Options{
			{Flag: "--component-type", Description: "component-type on which operations will be performed"},
			{Flag: "--entity", Description: "name of the entity [env|service|component]"},
		})
	}
	if o.Describe {
		return commandHelper("describe", "operation", "", []Options{
			{Flag: "--name", Description: "name of the operation"},
			{Flag: "--component-type", Description: "component-type on which operations will be performed"},
			{Flag: "--entity", Description: "name of the entity [env|service|component]"},
		})
	}
	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (o *Operation) Synopsis() string {
	if o.List {
		return "list all operations on environment, service or a component-type"
	}
	if o.Describe {
		return "describe a operation on environment, service or a component-type"
	}
	return defaultHelper()
}
