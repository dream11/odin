package deploy

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/config"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

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

const (
	Yes = "y"
	No  = "n"
)

func init() {

	serviceSetDeployCmd.Flags().StringVar(&env, "env", "", "environment for deploying the service-set")
	serviceSetDeployCmd.Flags().StringVar(&provisioningFile, "file", "", "path to the service set provisioning file")
	deployCmd.AddCommand(serviceSetDeployCmd)
}

func executeDeployServiceSet(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)

	ctx := cmd.Context()

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

	conflictingServicesRequest := &serviceProto.GetConflictingServicesRequest{
		EnvName:  env,
		Name:     deployServiceSetRequest.Name,
		Services: deployServiceSetRequest.Services,
	}

	services, errs := serviceClient.GetConflictingServices(&ctx, conflictingServicesRequest)
	if errs != nil {
		log.Fatal(fmt.Sprintf("Failed to list services with conflicting versions: %s", errs.Error()))
		return
	}
	for _, service := range services.Services {

		allowedInputsSlice := []string{Yes, No}
		allowedInputs := make(map[string]struct{}, len(allowedInputsSlice))
		for _, input := range allowedInputsSlice {
			allowedInputs[input] = struct{}{}
		}
		message := fmt.Sprintf("Service: %s already deployed with different version : %s \n Do you want to deploy service with new version : %s ? (y/n)", service.Name, service.ExistingVersion, service.NewVersion)
		inputHandler := ui.Input{}
		val, err := inputHandler.AskWithConstraints(message, allowedInputs)

		if err != nil {
			log.Fatal(fmt.Sprintf("An error occurred while processing input: %s", err.Error()))
		}

		if val != Yes {
			log.Info(fmt.Sprintf("Skipping service %s from deploy", service.Name))
			for _, svc := range deployServiceSetRequest.Services {
				if svc.ServiceName == service.Name {
					svc.ForceFlag = false
				}
			}
		}

	}

	err := serviceClient.DeployServiceSet(&ctx, &deployServiceSetRequest)
	if err != nil {
		log.Fatal("Failed to deploy service set. ", err)
	}
}
