package describe

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/dream11/odin/pkg/util"

	serviceBackend "github.com/dream11/odin/internal/service"
	service "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/iancoleman/orderedmap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serviceName string
var serviceVersion string
var component string
var verbose bool

var serviceClient = serviceBackend.Service{}
var serviceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Describe service",
	Args:  cobra.ExactArgs(1),
	Long:  `Describe definition and provisionig files of a service`,
	Run: func(cmd *cobra.Command, args []string) {
		serviceName = args[0]
		execute(cmd)
	},
}

func init() {
	serviceCmd.Flags().StringVar(&serviceVersion, "version", "", "version of the service")
	serviceCmd.Flags().StringVar(&component, "component", "", "Display the config of a specific component only")
	serviceCmd.Flags().BoolVarP(&verbose, "verbose", "V", false, "display provisioning files data")
	err := componentCmd.MarkFlagRequired("version")
	if err != nil {
		log.Fatal("Error marking 'version' flag as required:", err)
	}
	describeCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {

	params := map[string]string{
		"verbose": strconv.FormatBool(verbose),
	}

	if component != "" {
		params["component"] = component
	}

	ctx := cmd.Context()
	response, err := serviceClient.DescribeService(&ctx, &service.DescribeServiceRequest{
		ServiceName: serviceName,
		Version:     serviceVersion,
		Params:      params,
	})

	if err != nil {
		util.LogGrpcError(err, "Failed to describe service: ")
		os.Exit(1)
	}

	writeAsJSON(response)
}

func writeAsJSON(response *service.DescribeServiceResponse) {
	serviceData := orderedmap.New()
	serviceData.Set("name", response.Service.Name)

	if response.Service.Version != nil && *response.Service.Version != "" {
		serviceData.Set("version", *response.Service.Version)
	}

	if response.Service.ServiceDefinition != nil && len(response.Service.ServiceDefinition.GetFields()) > 0 {
		serviceData.Set("definition", response.Service.ServiceDefinition)
	}
	if len(response.Service.ProvisioningConfigFiles) > 0 {
		serviceData.Set("provision", response.Service.ProvisioningConfigFiles)
	}

	output, err := json.MarshalIndent(serviceData, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	fmt.Println(string(output))
}
