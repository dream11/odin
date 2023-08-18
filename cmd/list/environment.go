package list

import (
	"context"
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
	Long:  `List all types of environments created by current user or all environments`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Context())
	},
}

func execute(ctx context.Context) {
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

func init() {
	environmentCmd.Flags().StringVar(&name, "name", "", "name of the env")
	environmentCmd.Flags().StringVar(&provisioningType, "provisioning-type", "", "provisioning type of the environment")
	environmentCmd.Flags().StringVar(&account, "account", "", "cloud provider account name")
	environmentCmd.Flags().BoolVarP(&displayAll, "all", "A", false, "list all environments")
	listCmd.AddCommand(environmentCmd)
}
