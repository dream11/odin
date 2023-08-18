package list

import (
	"context"
	"github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/pkg/table"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var name string
var provisioningType string
var account string
var displayAll bool
var appConfig *configuration.Configuration

var environmentClient = backend.Env{}
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "List environments",
	Long:  `List all types of environments created by current user or all environments`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Context())
	},
}

func execute(ctx context.Context) {
	response, err := environmentClient.ListEnvironments(ctx, &environment.ListEnvironmentRequest{
		Params: map[string]string{
			"name":             name,
			"account":          account,
			"provisioningType": provisioningType,
			"displayAll":       strconv.FormatBool(displayAll)},
	})

	if err != nil {
		log.Error("Failed to list environments", err)
		os.Exit(1)
	}

	tableHeaders := []string{"Name", "Created By", "Provisioning Type", "State", "Account"}
	var tableData [][]interface{}

	for _, env := range response.Environments {
		tableData = append(tableData, []interface{}{
			env.Name,
			env.CreatedBy,
			env.Status,
			env.ProviderAccountName,
		})
	}

	table.Write(tableHeaders, tableData)
}

func init() {
	envCmd.Flags().StringVar(&name, "name", "", "name of the env")
	envCmd.Flags().StringVar(&provisioningType, "provisioning-type", "", "provisioning type of the environment")
	envCmd.Flags().StringVar(&account, "account", "", "cloud provider account name")
	envCmd.Flags().BoolVarP(&displayAll, "all", "A", false, "list all environments")
	listCmd.AddCommand(envCmd)
}
