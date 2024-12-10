package deploy

import (
	"encoding/json"
	"github.com/dream11/odin/pkg/util"
	"os"

	"github.com/dream11/odin/pkg/config"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serviceSetName string
var serviceSetDeployCmd = &cobra.Command{
	Use:   "service-set",
	Short: "Deploy service-set",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: "Deploy service-set using files or service set name",
	Run: func(cmd *cobra.Command, args []string) {
		executeDeployServiceSet(cmd)
	},
}

func init() {

	serviceSetDeployCmd.Flags().StringVar(&env, "env", "", "environment for deploying the service-set")
	serviceSetDeployCmd.Flags().StringVar(&provisioningFile, "file", "", "path to the service set provisioning file")
	serviceSetDeployCmd.Flags().StringVar(&serviceSetName, "name", "", "released service set name")
	deployCmd.AddCommand(serviceSetDeployCmd)
}

func executeDeployServiceSet(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)

	ctx := cmd.Context()
	traceID := util.GenerateTraceID()
	if serviceSetName == "" && provisioningFile == "" {
		log.Fatal("Please provide either --name or --file.")
	}
	if provisioningFile != "" && serviceSetName != "" {
		log.Fatal("--name should not be provided when --file is provided.")
	}
	var deployServiceSetRequestDTO serviceDto.ServiceSet
	var deployServiceSetRequest serviceProto.DeployServiceSetRequest

	deployServiceSetRequest.EnvName = env

	if provisioningFile != "" {
		provisioningData, err := os.ReadFile(provisioningFile)
		if err != nil {
			log.Fatal("Error while reading provisioning file ", err)
		}
		if err := json.Unmarshal(provisioningData, &deployServiceSetRequestDTO); err != nil {
			log.Fatal("error unmarshalling provisioning file: %w", err)
		}
		deployServiceSetRequest = serviceClient.ConvertToDeployServiceSetRequest(&deployServiceSetRequestDTO, env)
	}
	if serviceSetName != "" {
		deployServiceSetRequest.Name = serviceSetName
	}

	err := serviceClient.DeployServiceSet(&ctx, &deployServiceSetRequest, traceID)
	if err != nil {
		log.Info("TraceId: ", traceID)
		log.Fatal("Failed to deploy service set", err)
	}
}
