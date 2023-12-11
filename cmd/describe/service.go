package describe

import (
	"encoding/json"
	"fmt"
	"strconv"

	serviceBackend "github.com/dream11/odin/internal/service"
	service "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/iancoleman/orderedmap"
)

var serviceName string
var serviceVersion string
var labelsJSON string
var labels map[string]string
var component string
var verbose bool

var serviceClient = serviceBackend.Service{}
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Describe service",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `Describe definition and provisionig files of a service`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	serviceCmd.Flags().StringVar(&serviceName, "name", "", "name of the service")
	serviceCmd.Flags().StringVar(&serviceVersion, "version", "", "version of the service")
	serviceCmd.Flags().StringVar(&labelsJSON, "labels", "", "labels of the service in the artifactory")
	serviceCmd.Flags().StringVar(&component, "component", "", "Display the config of a specific component only")
	serviceCmd.Flags().BoolVarP(&verbose, "verbose", "V", false, "display provisioning files data")
	describeCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {

	validateFlags()

	params := map[string]string{
		"verbose": strconv.FormatBool(verbose),
	}

	if component != "" {
		params["component"] = component
	}
	
	ctx := cmd.Context()
	response, err := serviceClient.DescribeService(&ctx, &service.DescribeServiceRequest{
		ServiceName: serviceName,
		Version: serviceVersion,
		Labels: labels,
		Params: params,
	})

	if err != nil {
		log.Fatal("Failed to describe service: ", err)
	}

	writeAsJSON(response)
}

func validateFlags() {

	if serviceName == "" {
		log.Fatal("Please pass the --name flag")
	}

	if serviceVersion == "" && labelsJSON == "" {
		log.Fatal("Please pass either --version flag or --labels flag")
	}

	if serviceVersion != "" && labelsJSON != "" {
		log.Fatal("Please pass either --version flag or --labels flag but not both")
	}

	if labelsJSON != "" {
		err := json.Unmarshal([]byte(labelsJSON), &labels)
		if err != nil {
			log.Fatal("Error parsing JSON, the the key and values should be strings: ", err)
		}
	}
}

func writeAsJSON(response *service.DescribeServiceResponse) {
	serviceData := orderedmap.New()
	serviceData.Set("name", response.Service.Name)

	if response.Service.Version != nil && *response.Service.Version != "" {
		serviceData.Set("version", *response.Service.Version)
	}
	if response.Service.Versions != nil && len(response.Service.Versions) > 0 {
		serviceData.Set("versions", response.Service.Versions)
	}
	if response.Service.Labels != nil {
		serviceData.Set("labels", response.Service.Labels)
	}
	if response.Service.ServiceDefinition != nil && len(response.Service.ServiceDefinition.GetFields()) > 0 {
		serviceData.Set("definition", response.Service.ServiceDefinition)
	}
	if response.Service.ProvisioningConfigFiles != nil && len(response.Service.ProvisioningConfigFiles) > 0 {
		serviceData.Set("provision", response.Service.ProvisioningConfigFiles)
	}

	output, err := json.MarshalIndent(serviceData, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	fmt.Println(string(output))
}
