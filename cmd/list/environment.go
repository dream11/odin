package list

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string
var provisioningType string
var account string
var displayAll bool

var environmentClient = backend.Environment{}
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "List environments",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `List all types of environments created by current user or all environments`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	response, err := environmentClient.ListEnvironments(&ctx, &environment.ListEnvironmentRequest{
		Params: map[string]string{
			"name":             name,
			"account":          account,
			"provisioningType": provisioningType,
			"displayAll":       strconv.FormatBool(displayAll)},
	})

	if err != nil {
		log.Fatal("Failed to list environments ", err)
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(response, outputFormat)
}

func writeOutput(response *environment.ListEnvironmentResponse, format string) {
	if format == "text" {
		writeAsText(response)
	} else if format == "json" {
		writeAsJson(response)
	} else {
		log.Fatal("Unknown output format: ", format)
	}
}

func writeAsText(response *environment.ListEnvironmentResponse) {
	tableHeaders := []string{"Name", "Created By", "State", "Account"}
	var tableData [][]interface{}
	for _, env := range response.Environments {
		tableData = append(tableData, []interface{}{
			*env.Name,
			*env.CreatedBy,
			*env.Status,
			*env.ProviderAccountName,
		})
	}

	table.Write(tableHeaders, tableData)
}

func writeAsJson(response *environment.ListEnvironmentResponse) error {
	var environments []map[string]interface{}
	for _, env := range response.Environments {
		environments = append(environments, map[string]interface{}{
			"name":      *env.Name,
			"createdBy": *env.CreatedBy,
			"status":    *env.Status,
			"account":   *env.ProviderAccountName,
		})
	}
	output, err := json.MarshalIndent(environments, "", "  ")
	if err != nil {
		return err
	}
	fmt.Print(string(output))
	return nil
}

func init() {
	environmentCmd.Flags().StringVar(&name, "name", "", "name of the env")
	environmentCmd.Flags().StringVar(&provisioningType, "provisioning-type", "", "provisioning type of the environment")
	environmentCmd.Flags().StringVar(&account, "account", "", "cloud provider account name")
	environmentCmd.Flags().BoolVarP(&displayAll, "all", "A", false, "list all environments")
	listCmd.AddCommand(environmentCmd)
}
