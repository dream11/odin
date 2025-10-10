package status

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/util"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	"github.com/spf13/cobra"
)

// setstatusCmd represents the env command
var setstatusCmd = &cobra.Command{
	Use:   "env <envName>",
	Short: "Fetch status of the environment",
	Long:  `Fetch status of the environment`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envName = args[0]
		getStatus(cmd)
	},
}
var environmentClient = service.Environment{}
var serviceName string
var envName string

func init() {
	statusCmd.AddCommand(setstatusCmd)
	setstatusCmd.Flags().String("service", "", "Name of the service (optional)")
}

func getStatus(cmd *cobra.Command) {
	ctx := cmd.Context()
	serviceName, _ = cmd.Flags().GetString("service")
	if envName == "" {
		log.Fatal("Error: --name is required")
	}
	response, err := environmentClient.EnvironmentStatus(&ctx, &environment.StatusEnvironmentRequest{
		EnvName:     envName,
		ServiceName: serviceName,
	})
	if err != nil {
		util.LogGrpcError(err, "Failed to get environment status: ")
	}
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(response, outputFormat)
}

func writeOutput(response *environment.StatusEnvironmentResponse, format string) {

	switch format {
	case constant.TEXT:
		writeAsTextEnvResponse(response)
	case constant.JSON:
		writeAsJSONEnvResponse(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func writeAsTextEnvResponse(response *environment.StatusEnvironmentResponse) {
	fmt.Printf("Fetching status for environment: %s\n", response.GetEnvName())
	fmt.Printf("Environment Status: %s\n", response.GetEnvStatus())
	var tableHeaders = []string{"NAME",
		"VERSION",
		"STATUS",
		"LAST DEPLOYED",
	}
	var tableData [][]interface{}

	if serviceName == "" {
		fmt.Println("\nServices:")
		for _, svc := range response.GetServicesStatus() {
			tableData = append(tableData, []interface{}{
				svc.GetServiceName(),
				svc.GetServiceVersion(),
				svc.GetServiceStatus(),
				util.FormatToHumanReadableDuration(svc.GetLastDeployed()),
			})
		}
	} else {
		tableHeaders = []string{"NAME",
			"STATUS"}
		fmt.Printf("Fetching status for service: %s in environment: %s\n", serviceName, envName)
		for _, svc := range response.GetServicesStatus() {
			if svc.GetServiceName() == serviceName {
				fmt.Printf("Service version: %s\n", svc.GetServiceVersion())
				fmt.Printf("Service Status: %s\n", svc.GetServiceStatus())
				fmt.Printf("Last deployed: %s\n", util.FormatToHumanReadableDuration(svc.GetLastDeployed()))
				fmt.Println("Component details:")
				for _, component := range svc.GetComponentStatus() {
					tableData = append(tableData, []interface{}{
						component.GetComponentName(),
						component.GetComponentStatus(),
					})
				}
			}
		}
	}
	table.Write(tableHeaders, tableData)
}

func writeAsJSONEnvResponse(response *environment.StatusEnvironmentResponse) {
	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	fmt.Print(string(output))
}
