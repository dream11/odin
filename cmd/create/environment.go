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
var accounts string

var environmentClient service.Environment

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "env <name>",
	Short: "Create environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envName = args[0]
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
	environmentCmd.Flags().StringVar(&accounts, "accounts", "", "list of comma separated cloud provider accounts")
	err := environmentCmd.MarkFlagRequired("accounts")
	if err != nil {
		log.Fatal("Error marking 'accounts' flag as required:", err)
	}
	createCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	// Validate accounts parameter
	if err := validateAccounts(accounts); err != nil {
		log.Fatal("Invalid accounts parameter: ", err)
	}
	err := environmentClient.CreateEnvironment(&ctx, &environmentProto.CreateEnvironmentRequest{
		EnvName:  envName,
		Accounts: util.SplitProviderAccount(accounts),
	})

	if err != nil {
		util.LogGrpcError(err, "Failed to create environment: ")
	}
}
