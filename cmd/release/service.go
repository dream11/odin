package release

import (
	"encoding/json"
	"github.com/dream11/odin/pkg/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/dream11/odin/internal/service"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var path string
var serviceClient service.Service

// releaseServiceCmd represents the release service command
var releaseServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "release service",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	releaseServiceCmd.Flags().StringVar(&path, "path", "", "path for the service directory")
	if err := releaseServiceCmd.MarkFlagRequired("path"); err != nil {
		log.Fatal("Error marking 'path' flag as required:", err)
	}
	releaseCmd.AddCommand(releaseServiceCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	traceID := util.GenerateTraceID()

	var err error

	var serviceReleaseRequest serviceProto.ReleaseServiceRequest
	// Check if the folder exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("Folder does not exist:", path)
		return
	}

	// Check if the definition.json file exists
	definitionFilePath := filepath.Join(path, "definition.json")
	if _, err := os.Stat(definitionFilePath); os.IsNotExist(err) {
		log.Fatal("definition.json file does not exist")
	}

	// Check if the provision-*.json file exists
	// Open the directory
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening content the directory :", path)
	}
	defer dir.Close()

	// Read the directory contents
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal("Error while getting content of the directory :", path)
	}
	var provisioningConfigMap = make(map[string]*serviceDto.ProvisioningConfig)
	// Iterate through the files
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "provision") {
			// Split the filename to get the postfix
			parts := strings.Split(file.Name(), "-")
			if len(parts) > 1 {
				postfix := strings.Split(parts[1], ".")[0] // Get the part after "provision-" and before the file extension
				provisionFilePath := filepath.Join(path, file.Name())
				provisioningData, err := os.ReadFile(provisionFilePath)
				if err != nil {
					log.Fatal("Error while reading provisioning file :", provisionFilePath, err)
				}
				var compProvConfigs []*serviceDto.ComponentProvisioningConfig
				if err := json.Unmarshal(provisioningData, &compProvConfigs); err != nil {
					log.Fatalf("Error unmarshalling provisioning file: %v", err)
				}
				provisioningProto := serviceDto.ProvisioningConfig{
					ComponentProvisioningConfig: compProvConfigs,
				}
				provisioningConfigMap[postfix] = &provisioningProto
			}
		}
	}

	definitionData, err := os.ReadFile(definitionFilePath)
	if err != nil {
		log.Fatal("Error while reading definition file ", err)
	}
	var definitionProto serviceDto.ServiceDefinition
	if err := json.Unmarshal(definitionData, &definitionProto); err != nil {
		log.Fatalf("Error unmarshalling definition file: %v", err)
	}
	serviceReleaseRequest.ProvisioningConfigs = provisioningConfigMap
	serviceReleaseRequest.ServiceDefinition = &definitionProto
	err = serviceClient.ReleaseService(&ctx, &serviceReleaseRequest, traceID)
	if err != nil {
		log.Fatal("Failed to release service ", err)
	}
}
