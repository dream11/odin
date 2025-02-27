package deploy

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"regexp"

	"github.com/dream11/odin/cmd/util"
	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/config"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	envProto "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var env string
var definitionFile string
var provisioningFile string
var serviceName string
var serviceVersion string
var serviceClient = service.Service{}
var labels string
var envClient = service.Environment{}
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
	serviceCmd.Flags().StringVar(&serviceName, "name", "", "released service name")
	serviceCmd.Flags().StringVar(&serviceVersion, "version", "", "released service version")
	serviceCmd.Flags().StringVar(&labels, "labels", "", "comma separated labels for the service version ex key1=value1,key2=value2")

	deployCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)
	ctx := cmd.Context()

	if isStrictEnvironment(ctx, env) {

		util.AskForConfirmation(env)
	}

	if (serviceName == "" && serviceVersion == "" && labels == "") && (definitionFile != "" && provisioningFile != "") {
		deployUsingFiles(ctx)
	} else if (serviceName != "" && serviceVersion != "" && labels == "") && (definitionFile == "" && provisioningFile == "") {
		deployUsingServiceNameAndVersion(ctx)
	} else if (serviceName != "" && labels != "" && serviceVersion == "") && (definitionFile == "" && provisioningFile == "") {
		if err := validateLabels(labels); err != nil {
			log.Fatal("Invalid labels format: ", err)
		}
		deployUsingServiceNameAndLabels(ctx)
	} else {
		log.Fatal("Invalid combination of flags. Use either (service name and version) or (service name and labels) or (definitionFile and provisioningFile).")
	}
}

func deployUsingFiles(ctx context.Context) {
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

func deployUsingServiceNameAndVersion(ctx context.Context) {
	log.Info("deploying service :", serviceName, ":", serviceVersion, " in env :", env)
	err := serviceClient.DeployReleasedService(&ctx, &serviceProto.DeployReleasedServiceRequest{
		EnvName: env,
		ServiceIdentifier: &serviceProto.ServiceIdentifier{
			ServiceName:    serviceName,
			ServiceVersion: serviceVersion,
		},
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}
}

func deployUsingServiceNameAndLabels(ctx context.Context) {
	log.Info("deploying service :", serviceName, " with labels: ", labels, " in env :", env)
	err := serviceClient.DeployReleasedService(&ctx, &serviceProto.DeployReleasedServiceRequest{
		EnvName: env,
		ServiceIdentifier: &serviceProto.ServiceIdentifier{
			ServiceName: serviceName,
			Labels:      labels,
		},
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}
}

func validateLabels(labels string) error {
	labelPattern := `^(\w+=\w+)(,\w+=\w+)*$`
	matched, err := regexp.MatchString(labelPattern, labels)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("labels must be in format key1=value1,key2=value2")
	}
	return nil
}
func isStrictEnvironment(ctx context.Context, env string) bool {
	envTypeResp, err := envClient.IsStrictEnvironment(&ctx, &envProto.IsStrictEnvironmentRequest{
		EnvName: env,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return envTypeResp.IsEnvStrict
}
