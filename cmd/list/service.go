package list

import (
	"encoding/json"
	"fmt"
	"github.com/dream11/odin/pkg/util"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serviceClient = service.Service{}
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "List services",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `List all services if no options provided, type odin list service --help for details`,
	Run: func(cmd *cobra.Command, args []string) {
		listService(cmd)
	},
}

var serviceName, version, tags string

func init() {
	serviceCmd.Flags().StringVar(&serviceName, "name", "", "name of the service")
	serviceCmd.Flags().StringVar(&version, "version", "", "version of services to be listed")
	serviceCmd.Flags().StringVar(&tags, "tags", "", "comma separated tags eg. key1=value1,key2=value2")
	listCmd.AddCommand(serviceCmd)
}

func listService(cmd *cobra.Command) {
	ctx := cmd.Context()
	traceId := util.GenerateTraceId()
	response, err := serviceClient.ListService(&ctx, &serviceProto.ListServiceRequest{
		Name:    serviceName,
		Version: version,
		Tags:    tags,
	}, traceId)

	if err != nil {
		log.Fatal("Failed to list services ", err)
	}
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal("Failed to get --output global option : ", err)
	}
	writeListService(response, outputFormat)
}

func writeListService(response *serviceProto.ListServiceResponse, format string) {

	switch format {
	case constant.TEXT:
		writeListServiceAsText(response)
	case constant.JSON:
		writeListServiceAsJSON(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func writeListServiceAsText(response *serviceProto.ListServiceResponse) {
	var tableHeaders = []string{"Name", "Version", "Tags"}
	var tableData [][]interface{}
	for _, serviceEntity := range response.Services {
		tableData = append(tableData, []interface{}{
			serviceEntity.Name,
			serviceEntity.Version,
			serviceEntity.Tags,
		})
	}
	table.Write(tableHeaders, tableData)
}

func writeListServiceAsJSON(response *serviceProto.ListServiceResponse) {
	var services []map[string]interface{}
	for _, serviceEntity := range response.Services {
		services = append(services, map[string]interface{}{
			"name":    serviceEntity.Name,
			"version": serviceEntity.Version,
			"Tags":    serviceEntity.Tags,
		})
	}
	output, _ := json.MarshalIndent(services, "", "  ")
	fmt.Print(string(output))
}
