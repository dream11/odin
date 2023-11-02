package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path"

	"github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/service"
	viperConfig "github.com/dream11/odin/pkg/config"
	auth "github.com/dream11/odin/proto/gen/go/dream11/od/auth/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var odinAccessKey string
var odinSecretAccessKey string
var odinBackendAddress string
var odinInsecure bool
const DEFAULT_BACKEND_ADDR = "odin-backend.d11dev.com:443"

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
	RootCmd.AddCommand(configureCmd)
}

func execute(cmd *cobra.Command) {
	configPath := path.Join(app.WorkDir.Location, app.WorkDir.ConfigFile)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			log.Fatal("Error creating the config file:", err)
		}
		defer file.Close()
	} else if err != nil {
		log.Fatal("Error checking the config file:", err)
	}
	config := viperConfig.GetConfig()

	if odinBackendAddress != "" {
		config.BackendAddr = odinBackendAddress
	} else if os.Getenv("ODIN_BACKEND_ADDRESS") != "" {
		config.BackendAddr = os.Getenv("ODIN_BACKEND_ADDRESS")
	} else {
		config.BackendAddr = DEFAULT_BACKEND_ADDR
	}

	config.Insecure = odinInsecure

	if odinAccessKey != "" {
		config.Keys.AccessKey = odinAccessKey
	} else if os.Getenv("ODIN_ACCESS_KEY") != "" {
		config.Keys.AccessKey = os.Getenv("ODIN_ACCESS_KEY")
	} else {
		log.Fatal("Neither access-key flag is passed nor environment variable ODIN_ACCESS_KEY is set. Please either pass your access key using odin-access-key or set your access key in ODIN_ACCESS_KEY environment variable")
	}

	if odinSecretAccessKey != "" {
		config.Keys.SecretAccessKey = odinSecretAccessKey
	} else if os.Getenv("ODIN_SECRET_ACCESS_KEY") != "" {
		config.Keys.SecretAccessKey = os.Getenv("ODIN_SECRET_ACCESS_KEY")
	} else {
		log.Fatal("Neither secret-access-key flag is passed nor environment variable ODIN_SECRET_ACCESS_KEY is set. Please either pass your access key using odin-secret-access-key or set your access key in ODIN_SECRET_ACCESS_KEY environment variable")
	}

	sha256 := sha256.New()
	sha256.Write([]byte(config.Keys.SecretAccessKey))
	hashedResult := sha256.Sum(nil)

	ctx := cmd.Context()
	response, err := configureClient.GetUserToken(&ctx, &auth.GetUserTokenRequest{
		ClientId: string(config.Keys.AccessKey),
		ClientSecretHash: hex.EncodeToString(hashedResult),
	})
	if err != nil {
		log.Fatal("Failed to get token ", err)
	}
	config.AccessToken = response.GetToken()

	profile := viper.GetString("profile")
	viper.Set(profile, config)
	if err := viperConfig.SetConfig(); err != nil {
		log.Fatal("Unable to write configuration: ", err)
	}

	fmt.Println("Configured!")
}
