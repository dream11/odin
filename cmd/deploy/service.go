package deploy

import (
	"encoding/json"
	"github.com/dream11/odin/internal/service"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var env string
var provisioning string
var file string

var serviceClient = service.Service{}

var serviceCmd = &cobra.Command{
	Use:   "Deploy service",
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
	serviceCmd.Flags().StringVar(&file, "file", "", "path to the service definition file")
	serviceCmd.Flags().StringVar(&provisioning, "provisioning", "", "path to the provisioning file")
	err := serviceCmd.MarkFlagRequired("env")
	if err != nil {
		log.Println("Error marking 'env' flag as required:", err)
		os.Exit(1)
	}

	deployCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()

	definitionData, err := readFromFile(file)
	if err != nil {
		log.Fatal("Error while reading definition file ", err)
	}
	provisioningData, err := readFromFile(provisioning)

	if err != nil {
		log.Fatal("Error while reading provisioning file ", err)
	}

	err = serviceClient.DeployService(&ctx, &serviceProto.DeployServiceRequest{
		EnvName:            env,
		ServiceDefinition:  definitionData,
		ProvisioningConfig: provisioningData,
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}

	//outputFormat, err := cmd.Flags().GetString("output")
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func readFromFile(filePath string) (string, error) {
	// Read the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading file:", err)
		return "", err
	}

	// Validate if the contents are valid JSON
	var jsonData interface{}
	err = json.Unmarshal(content, &jsonData)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return "", err

	}

	// Convert the JSON content to a string
	jsonString, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		log.Println("Error converting JSON to string:", err)
		return "", err

	}

	// Print the resulting string
	return string(jsonString), nil
}
