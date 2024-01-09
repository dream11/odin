package create

import (
	"strings"

	"github.com/dream11/odin/internal/service"
	environmentProto "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var envName string
var provisioningType string
var accounts string

var environmentClient service.Environment

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Create environment",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	environmentCmd.Flags().StringVar(&envName, "name", "", "name of the environment to be created")
	environmentCmd.Flags().StringVar(&accounts, "accounts", "", "list of comma separated cloud provider accounts")
	environmentCmd.Flags().StringVar(&provisioningType, "provisioning-type", "dev", "provisioning type of the environment")
	err := environmentCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	createCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()

	err := environmentClient.CreateEnvironment(&ctx, &environmentProto.CreateEnvironmentRequest{
		EnvName:          envName,
		Accounts:         splitProviderAccount(accounts),
		ProvisioningType: provisioningType,
	})

	if err != nil {
		log.Fatal("Failed to create environment ", err)
	}
}

func splitProviderAccount(providerAccounts string) []string {
	if providerAccounts == "" {
		return nil
	}
	return strings.Split(providerAccounts, ",")
}
