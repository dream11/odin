package deploy

import (
	"encoding/json"
	"os"

	"github.com/dream11/odin/internal/service"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var env string
var definitionFile string
var provisioningFile string

var serviceClient = service.Service{}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Deploy service",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: "Deploy service using files or service name",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	serviceCmd.Flags().StringVar(&env, "env", "", "environment for deploying the service")
	serviceCmd.Flags().StringVar(&definitionFile, "file", "", "path to the service definition file")
	serviceCmd.Flags().StringVar(&provisioningFile, "provisioning", "", "path to the provisioning file")
	err := serviceCmd.MarkFlagRequired("env")
	if err != nil {
		log.Fatal("Error marking 'env' flag as required:", err)
	}

	deployCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()

	definitionData, err := os.ReadFile(definitionFile)
	if err != nil {
		log.Fatal("Error while reading definition file ", err)
	}
	var definitionProto serviceDto.ServiceDefinition
	if err := json.Unmarshal(definitionData, &definitionProto); err != nil {
		log.Fatalf("Error unmarshalling definition file: %v", err)
	}

	provisioningData, err := os.ReadFile(provisioningFile)
	if err != nil {
		log.Fatal("Error while reading provisioning file ", err)
	}
	var compProvConfigs []*serviceDto.ComponentProvisioningConfig
	if err := json.Unmarshal(provisioningData, &compProvConfigs); err != nil {
		log.Fatalf("Error unmarshalling provisioning file: %v", err)
	}
	provisioningProto := &serviceDto.ProvisioningConfig{
		ComponentProvisioningConfig: compProvConfigs,
	}

	err = serviceClient.DeployService(&ctx, &serviceProto.DeployServiceRequest{
		EnvName:            env,
		ServiceDefinition:  &definitionProto,
		ProvisioningConfig: provisioningProto,
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}
}
