package configure

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path"

	"github.com/dream11/odin/app"
	"github.com/dream11/odin/cmd"
	"github.com/dream11/odin/internal/service"
	appConfig "github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/dir"
	auth "github.com/dream11/odin/proto/gen/go/dream11/od/auth/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var odinAccessKey string
var odinSecretAccessKey string
var odinBackendAddress string
var odinInsecure bool

// default backend address
const defaultBackendAddress = "odin-backend.d11dev.com:443"

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
	configureCmd.Flags().StringVar(&odinAccessKey, "access-key", "", "odin access key")
	configureCmd.Flags().StringVar(&odinSecretAccessKey, "secret-access-key", "", "odin secret access key")
	configureCmd.Flags().StringVar(&odinBackendAddress, "backend-address", "", "odin backend address with port")
	configureCmd.Flags().BoolVarP(&odinInsecure, "insecure", "I", false, "odin insecure")
	cmd.RootCmd.AddCommand(configureCmd)
}

func execute(cmd *cobra.Command) {
	createConfigFileIfNotExist()

	config := appConfig.GetConfig()

	config.BackendAddress = getConfigKey("backend-address", odinBackendAddress, "ODIN_BACKEND_ADDRESS", config.BackendAddress, defaultBackendAddress)
	config.Insecure = odinInsecure
	config.Keys.AccessKey = getConfigKey("access-key", odinAccessKey, "ODIN_ACCESS_KEY", config.Keys.AccessKey, "")
	config.Keys.SecretAccessKey = getConfigKey("secret-access-key", odinSecretAccessKey, "ODIN_SECRET_ACCESS_KEY", config.Keys.SecretAccessKey, "")

	ctx := cmd.Context()
	response, err := configureClient.GetUserToken(&ctx, &auth.GetUserTokenRequest{
		ClientId:         string(config.Keys.AccessKey),
		ClientSecretHash: hashKey(config.Keys.SecretAccessKey),
	})
	if err != nil {
		log.Fatal("Failed to get token ", err)
	}

	config.AccessToken = response.Token
	appConfig.WriteConfig(config)
	fmt.Println("Configured!")
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

func getConfigKey(flagKey string, flagValue string, envVariableName string, configValue string, defaultValue string) string {
	if flagValue != "" {
		return flagValue
	} else if os.Getenv(envVariableName) != "" {
		return os.Getenv(envVariableName)
	} else if configValue != "" {
		return configValue
	} else if defaultValue != "" {
		return defaultValue
	}
	log.Fatalf("Please pass %s flag nor set environment variable %s to configure", flagKey, envVariableName)
	return ""
}

func hashKey(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashedResult := hash.Sum(nil)
	return hex.EncodeToString(hashedResult)
}
