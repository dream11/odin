package commands

import (
	"strings"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
)

// initiate backend client for environment type
var envTypeClient backend.EnvType

// Env : command declaration
type EnvType command

// Run : implements the actual functionality of the command
func (e *EnvType) Run(args []string) int {

	if e.ListEnvType {
		e.Logger.Info("Listing all env type")
		envList, err := envClient.ListEnvType()
		e.Logger.Info(strings.Join(envList, ""))
		if err != nil {
			e.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name"}
		var tableData [][]interface{}

		// for _, env := range envList {
		// 	tableData = append(tableData, []interface{}{
		// 		e.Name,
		// 	})
		// }

		table.Write(tableHeaders, tableData)

		return 0
	}

	e.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (e *EnvType) Help() string {

	if e.ListEnvType {
		return commandHelper("list env-type", "environment", "", []Options{})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (e *EnvType) Synopsis() string {

	if e.ListEnvType {
		return "List env types"
	}
	return defaultHelper()
}
