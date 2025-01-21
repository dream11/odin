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

const (
	Yes = "y"
	No  = "n"
)

func init() {

	serviceSetDeployCmd.Flags().StringVar(&env, "env", "", "environment for deploying the service-set")
	serviceSetDeployCmd.Flags().StringVar(&provisioningFile, "file", "", "path to the service set provisioning file")
	serviceSetDeployCmd.Flags().StringVar(&serviceSetName, "name", "", "released service set name")
	deployCmd.AddCommand(serviceSetDeployCmd)
}

func executeDeployServiceSet(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)

	ctx := cmd.Context()
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
	var serviceNames []string
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
			serviceNames = append(serviceNames, service.Name)
		}
		// Remove services from deployServiceSetRequest
		var updatedServices []*serviceProto.ServiceIdentifier
		for _, svc := range deployServiceSetRequest.Services {
			shouldRemove := false
			for _, name := range serviceNames {
				if svc.ServiceName == name {
					shouldRemove = true
					break
				}
			}
			if !shouldRemove {
				updatedServices = append(updatedServices, svc)
			}
		}
		deployServiceSetRequest.Services = updatedServices

	}

	err := serviceClient.DeployServiceSet(&ctx, &deployServiceSetRequest)
	if err != nil {
		log.Fatal("Failed to deploy service set. ", err)
	}
}
