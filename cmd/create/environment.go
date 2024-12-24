package create

import (
	"fmt"
	"strings"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/util"
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
	Use:   "env",
	Short: "Create environment",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func validateAccounts(accounts string) error {
	if accounts == "" {
		return fmt.Errorf("accounts parameter cannot be an empty string")
	}
	accountList := strings.Split(accounts, ",")
	for _, account := range accountList {
		if account == "" {
			return fmt.Errorf("accounts parameter should not end with a comma")
		}
	}
	return nil
}

func init() {
	environmentCmd.Flags().StringVar(&envName, "name", "", "name of the environment to be created")
	environmentCmd.Flags().StringVar(&accounts, "accounts", "", "list of comma separated cloud provider accounts")
	environmentCmd.Flags().StringVar(&provisioningType, "provisioning-type", "", "provisioning type of the environment")
	err := environmentCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	if err := environmentCmd.MarkFlagRequired("provisioning-type"); err != nil {
		log.Fatal("Error marking 'provisioning-type' flag as required:", err)
	}
	err = environmentCmd.MarkFlagRequired("accounts")
	if err != nil {
		log.Fatal("Error marking 'accounts' flag as required:", err)
	}
	createCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	traceId := util.GenerateTraceId()
	// Validate accounts parameter
	if err := validateAccounts(accounts); err != nil {
		log.Fatal("Invalid accounts parameter: ", err)
	}
	err := environmentClient.CreateEnvironment(&ctx, &environmentProto.CreateEnvironmentRequest{
		EnvName:          envName,
		Accounts:         util.SplitProviderAccount(accounts),
		ProvisioningType: provisioningType,
	}, traceId)

	if err != nil {
		log.Info("TraceId: ", traceId)
		log.Fatal("Failed to create environment ", err)
	}
}
