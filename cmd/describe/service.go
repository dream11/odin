package describe

import (
	"encoding/json"
	"fmt"
	"strconv"

	serviceBackend "github.com/dream11/odin/internal/service"
	service "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	err := json.Unmarshal([]byte(labelsJSON), &labels)
	if err != nil {
		log.Fatal("Error parsing JSON, the the key and values should be strings: ", err)
	}

	ctx := cmd.Context()
	response, err := serviceClient.DescribeService(&ctx, &service.DescribeServiceRequest{
		ServiceName: serviceName,
		Version: serviceVersion,
		Labels: labels,
		Params: map[string]string{
			"component":     component,
			"verbose":       strconv.FormatBool(verbose)},
	})

	if err != nil {
		log.Fatal("Failed to describe service ", err)
	}

	writeAsJSON(response)
}

func writeAsJSON(response *service.DescribeServiceResponse) {
	serviceData := map[string]interface{}{
		"name": response.Service.Name,
		"version": response.Service.Version,
		"defintion": response.Service.ServiceDefinition,
		"provision": response.Service.ProvisioningConfigFiles,
		"labels": response.Service.Labels,
		"versions": response.Service.Versions,
	}
	output, err := json.MarshalIndent(serviceData, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	fmt.Print(string(output))
}
