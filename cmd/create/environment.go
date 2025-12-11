package create

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/util"
	environmentProto "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	providerAccountProto "github.com/dream11/odin/proto/gen/go/dream11/oam/provideraccount/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

var envName string
var provisioningType string
var accounts string
var accountsFile string

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

func readAccountsFromFile(filePath string) ([]*providerAccountProto.GetProviderAccountResponse, error) {
	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read accounts file: %w", err)
	}

	// Parse JSON array
	var accountsJSON []json.RawMessage
	if err := json.Unmarshal(data, &accountsJSON); err != nil {
		return nil, fmt.Errorf("failed to parse accounts file as JSON array: %w", err)
	}

	// Convert each JSON object to GetProviderAccountResponse
	var accountResponses []*providerAccountProto.GetProviderAccountResponse
	for i, accountJSON := range accountsJSON {
		// Wrap in account field to match GetProviderAccountResponse structure
		wrappedJSON := fmt.Sprintf(`{"account": %s}`, string(accountJSON))

		var accountResp providerAccountProto.GetProviderAccountResponse
		if err := protojson.Unmarshal([]byte(wrappedJSON), &accountResp); err != nil {
			return nil, fmt.Errorf("failed to parse account at index %d: %w", i, err)
		}
		accountResponses = append(accountResponses, &accountResp)
	}

	return accountResponses, nil
}

func init() {
	environmentCmd.Flags().StringVar(&envName, "name", "", "name of the environment to be created")
	environmentCmd.Flags().StringVar(&accounts, "accounts", "", "list of comma separated cloud provider accounts")
	environmentCmd.Flags().StringVar(&accountsFile, "accounts-file", "", "path to JSON file containing account details")
	environmentCmd.Flags().StringVar(&provisioningType, "provisioning-type", "", "provisioning type of the environment")
	err := environmentCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	if err := environmentCmd.MarkFlagRequired("provisioning-type"); err != nil {
		log.Fatal("Error marking 'provisioning-type' flag as required:", err)
	}
	createCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()

	// Validate mutual exclusivity of --accounts and --accounts-file
	hasAccounts := accounts != ""
	hasAccountsFile := accountsFile != ""

	if hasAccounts && hasAccountsFile {
		log.Fatal("Cannot use both --accounts and --accounts-file flags together")
	}

	// If neither flag provided, try default file
	defaultAccountsFile := filepath.Join(os.Getenv("HOME"), ".odin", "local_account.json")
	if !hasAccounts && !hasAccountsFile {
		if _, err := os.Stat(defaultAccountsFile); err == nil {
			log.Debugf("Using default accounts file: %s", defaultAccountsFile)
			accountsFile = defaultAccountsFile
			hasAccountsFile = true
		} else {
			log.Fatal("Either --accounts or --accounts-file must be provided, or create default file at ~/.odin/local_account.json")
		}
	}

	// Auto-detect MAC address to use as routing key
	macAddr, err := util.GetDefaultMACAddress()
	if err != nil {
		log.Fatalf("Failed to auto-detect MAC address: %v", err)
	}
	log.Debugf("Using MAC address as routing key: %s", macAddr)

	// Create the request
	req := &environmentProto.CreateEnvironmentRequest{
		EnvName:          envName,
		ProvisioningType: provisioningType,
		RoutingKey:       &macAddr,
	}

	// Populate accounts or account_details based on which flag was used
	if hasAccounts {
		// Existing behavior: validate and pass account names
		if err := validateAccounts(accounts); err != nil {
			log.Fatal("Invalid accounts parameter: ", err)
		}
		req.Accounts = util.SplitProviderAccount(accounts)
	} else if hasAccountsFile {
		// New behavior: read and pass full account data
		accountDetails, err := readAccountsFromFile(accountsFile)
		if err != nil {
			log.Fatalf("Failed to read accounts file: %v", err)
		}
		req.AccountDetails = accountDetails
	}

	err = environmentClient.CreateEnvironment(&ctx, req)

	if err != nil {
		util.LogGrpcError(err, "Failed to create environment: ")
	}
}
