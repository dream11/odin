package deploy

import (
	"encoding/json"
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
		EnvName: env,
		Name:    deployServiceSetRequest.Name,
	}

	services, errs := serviceClient.GetConflictingServices(&ctx, conflictingServicesRequest)
	if errs != nil {
		log.Fatal("Failed to list services with conflicting versions. ", errs)
		return
	}
	var serviceNames []string
	for _, service := range services.Services {

		allowedInputsSlice := []string{"y", "n"}
		allowedInputs := make(map[string]struct{}, len(allowedInputsSlice))
		for _, input := range allowedInputsSlice {
			allowedInputs[input] = struct{}{}
		}
		message := "Service already deployed with different version.Do you want to deploy service " + service.Name + " ? (y/n)"
		inputHandler := ui.Input{}
		val, err := inputHandler.AskWithConstraints(message, allowedInputs)

		if err != nil {
			log.Fatal(err.Error())
		}

		if val != "y" {
			log.Info("Skipping service ", service.Name, " from deploy")
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
