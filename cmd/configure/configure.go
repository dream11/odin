package configure

import (
	"context"
	"fmt"
	"os"
	"path"

	apiConfig "github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/app"
	"github.com/dream11/odin/cmd"
	"github.com/dream11/odin/internal/auth"
	"github.com/dream11/odin/internal/service"
	appConfig "github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/dir"
	"github.com/dream11/odin/pkg/util"
	pb "github.com/dream11/odin/proto/gen/go/dream11/od/auth/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	configureCmd.Flags().String("backend-address", "", "odin backend address with port")
	configureCmd.Flags().BoolP("insecure", "I", true, "odin insecure")
	configureCmd.Flags().BoolP("plaintext", "P", false, "skip tls verification")
	configureCmd.Flags().Int64("org-id", 0, "organisation id")

	// Bind flags to viper for automatic precedence handling
	if err := viper.BindPFlag("backend_address", configureCmd.Flags().Lookup("backend-address")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("insecure", configureCmd.Flags().Lookup("insecure")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("plaintext", configureCmd.Flags().Lookup("plaintext")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("org_id", configureCmd.Flags().Lookup("org-id")); err != nil {
		panic(err)
	}

	cmd.RootCmd.AddCommand(configureCmd)
}

func execute(cmd *cobra.Command) {
	createConfigFileIfNotExist()

	// Read configurations from existing config file, env variables and flags in viper
	appConfig.GetConfig()

	if !viper.IsSet("backend_address") || viper.GetString("backend_address") == "" {
		log.Fatalf("Required configuration not found. Please pass --backend-address flag or set environment variable ODIN_BACKEND_ADDRESS")
	}
	if !viper.IsSet("org_id") {
		log.Fatalf("Required configuration not found. Please pass --org-id flag or set environment variable ODIN_ORG_ID")
	}

	// Collect user input and write base config to file against the active profile
	baseConfig := &apiConfig.Configuration{
		BackendAddress: viper.GetString("backend_address"),
		OrgId:          viper.GetInt64("org_id"),
		Insecure:       viper.GetBool("insecure"),
		Plaintext:      viper.GetBool("plaintext"),
	}
	appConfig.WriteConfig(baseConfig)

	ctx := cmd.Context()
	traceID := util.GenerateTraceID()
	contextWithTrace := context.WithValue(ctx, constant.TraceIDKey, traceID)
	authProviderResponse, err := configureClient.GetAuthProvider(&contextWithTrace, &pb.GetAuthProviderRequest{
		OrgId: &baseConfig.OrgId,
	})
	if err != nil {
		log.Fatalf("Failed to get auth provider: %v ", err)
	}

	provider, err := auth.GetProvider(authProviderResponse.Type)
	if err != nil {
		log.Fatalf("Error getting auth provider: %v", err)
	}

	authData, err := provider.Authenticate(authProviderResponse.Data)
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}

	tokenResponse, err := configureClient.GetUserToken(&contextWithTrace, &pb.GetUserTokenRequest{
		OrgId: &baseConfig.OrgId,
		Data:  authData,
	})
	if err != nil {
		util.LogGrpcError(err, "Failed to get token ")
	}

	// Persist token to config file against the active profile
	baseConfig.AccessToken = tokenResponse.Token
	appConfig.WriteConfig(baseConfig)

	fmt.Println("\033[32mConfigured!\033[0m")
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
