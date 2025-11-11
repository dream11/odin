package list

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/util"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var account string
var displayAll bool

var environmentClient = service.Environment{}
var environmentCmd = &cobra.Command{
	Use:   "env",
	Short: "List environments",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `List all types of environments created by current user or all environments`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	environmentCmd.Flags().StringVar(&account, "account", "", "cloud provider account name")
	environmentCmd.Flags().BoolVarP(&displayAll, "all", "A", false, "list all environments")
	listCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	response, err := environmentClient.ListEnvironments(&ctx, &environment.ListEnvironmentRequest{
		Params: map[string]string{
			"account":    account,
			"displayAll": strconv.FormatBool(displayAll)},
	})

	if err != nil {
		util.LogGrpcError(err, "Failed to list environments: ")
		os.Exit(1)
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(response, outputFormat)
}

func writeOutput(response *environment.ListEnvironmentResponse, format string) {

	switch format {
	case constant.TEXT:
		writeAsText(response)
	case constant.JSON:
		writeAsJSON(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func writeAsText(response *environment.ListEnvironmentResponse) {
	tableHeaders := []string{"Name", "State", "Account", "Provisioning Type"}
	var tableData [][]interface{}
	for _, env := range response.Environments {
		tableData = append(tableData, []interface{}{
			env.Name,
			env.State,
			env.Account,
		})
	}

	table.Write(tableHeaders, tableData)
}

func writeAsJSON(response *environment.ListEnvironmentResponse) {
	var environments []map[string]interface{}
	for _, env := range response.Environments {
		environments = append(environments, map[string]interface{}{
			"name":    env.Name,
			"status":  env.State,
			"account": env.Account,
		})
	}
	output, _ := json.MarshalIndent(environments, "", "  ")
	fmt.Print(string(output))
}
