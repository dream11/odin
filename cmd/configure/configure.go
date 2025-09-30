package configure

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/dream11/odin/internal/auth"

	"github.com/dream11/odin/app"
	"github.com/dream11/odin/cmd"
	"github.com/dream11/odin/internal/service"
	appConfig "github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/dir"
	"github.com/dream11/odin/pkg/util"
	pb "github.com/dream11/odin/proto/gen/go/dream11/od/auth/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var odinBackendAddress string
var insecure bool
var plainText bool
var orgID int64

var configureClient = service.Configure{}
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure odin",
	Long:  "Configure odin using odin access key and odin secret access key",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	configureCmd.Flags().StringVar(&odinBackendAddress, "backend-address", "", "odin backend address with port")
	configureCmd.Flags().BoolVarP(&insecure, "insecure", "I", true, "odin insecure")
	configureCmd.Flags().BoolVarP(&plainText, "plaintext", "P", false, "skip tls verification")
	configureCmd.Flags().Int64Var(&orgID, "org-id", 0, "organisation id")
	cmd.RootCmd.AddCommand(configureCmd)
}

func execute(cmd *cobra.Command) {
	createConfigFileIfNotExist()

	config := appConfig.GetConfig()

	config.BackendAddress = getConfigKey("backend-address", odinBackendAddress, "ODIN_BACKEND_ADDRESS", config.BackendAddress)
	config.Insecure = insecure
	config.Plaintext = plainText
	config.OrgId = getConfigKey("org-id", orgID, "ODIN_ORG_ID", config.OrgId)

	ctx := cmd.Context()
	authProviderResponse, err := configureClient.GetAuthProvider(&ctx, &pb.GetAuthProviderRequest{
		OrgId: &config.OrgId,
	})
	if err != nil {
		util.LogGrpcError(err, "Failed to get auth provider ")
	}

	provider, err := auth.GetProvider(authProviderResponse.Type)
	if err != nil {
		log.Fatalf("Error getting auth provider: %v", err)
	}

	authData, err := provider.Authenticate(authProviderResponse.Data)
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}

	tokenResponse, err := configureClient.GetUserToken(&ctx, &pb.GetUserTokenRequest{
		OrgId: &config.OrgId,
		Data:  authData,
	})
	if err != nil {
		util.LogGrpcError(err, "Failed to get token ")
	}

	config.AccessToken = tokenResponse.Token
	appConfig.WriteConfig(config)
	fmt.Println("\n\033[32mConfigured!\033[0m")
}

func createConfigFileIfNotExist() {
	dirPath := path.Join(os.Getenv("HOME"), "."+app.App.Name)
	if err := dir.CreateDirIfNotExist(dirPath); err != nil {
		log.Fatalf("Error creating the .%s folder: %v", app.App.Name, err)
	}
	configPath := path.Join(dirPath, "config")
	if err := dir.CreateFileIfNotExist(configPath); err != nil {
		log.Fatal("Error creating the config file")
	}
}

func getConfigKey[T comparable](flagKey string, flagValue T, envVariableName string, configValue T) T {
	var zero T
	if flagValue != zero {
		return flagValue
	}

	if envValStr := os.Getenv(envVariableName); envValStr != "" {
		var result T
		var a any = &result
		switch p := a.(type) {
		case *string:
			*p = envValStr
		case *int64:
			val, err := strconv.ParseInt(envValStr, 10, 64)
			if err != nil {
				log.Fatalf("Invalid value for environment variable %s: %v", envVariableName, err)
			}
			*p = val
		default:
			log.Fatalf("Unsupported type for getConfigKey: %T", zero)
		}
		return result
	}

	if configValue != zero {
		return configValue
	}

	log.Fatalf("Required configuration not found. Please pass --%s flag or set environment variable %s", flagKey, envVariableName)
	return zero // Unreachable
}
