package list

import (
	"encoding/json"
	"fmt"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	serviceproto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
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
	Long: `List services all if no options provided, type odin list service --help for details`,
	Run: func(cmd *cobra.Command, args []string) {
		listService(cmd)
	},
}

var version, team, label string
var tableHeaders = []string{"Name", "Latest Version", "Label", "Description", "Team"}
var tableData [][]interface{}

func init() {
	serviceCmd.Flags().StringVar(&name, "name", "", "name of the service")
	serviceCmd.Flags().StringVar(&version, "version", "", "version of services to be listed")
	serviceCmd.Flags().StringVar(&team, "team", "", "name of team")
	serviceCmd.Flags().StringVar(&label, "label", "", "name of label")
	listCmd.AddCommand(serviceCmd)
}

func listService(cmd *cobra.Command) {
	ctx := cmd.Context()
	response, err := serviceClient.ListService(&ctx, &serviceproto.ListServiceRequest{
		Name:    name,
		Version: version,
		Team:    team,
		Label:   label,
	})

	if err != nil {
		log.Fatal("Failed to list services ", err)
	}
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	WriteListService(response, outputFormat)
}

func WriteListService(response *serviceproto.ListServiceResponse, format string) {

	switch format {
	case constant.TEXT:
		writeListServiceAsText(response)
	case constant.JSON:
		writeListServiceAsJSON(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func writeListServiceAsText(response *serviceproto.ListServiceResponse) {
	for _, serviceEntity := range response.Services {
		tableData = append(tableData, []interface{}{
			serviceEntity.Name,
			serviceEntity.Version,
			serviceEntity.Labels,
			serviceEntity.Description,
		})
	}
	table.Write(tableHeaders, tableData)
}

func writeListServiceAsJSON(response *serviceproto.ListServiceResponse) {
	var services []map[string]interface{}
	for _, serviceEntity := range response.Services {
		tableData = append(tableData, []interface{}{
			serviceEntity.Name,
			serviceEntity.Version,
			serviceEntity.Labels,
			serviceEntity.Description,
		})
	}
	output, _ := json.MarshalIndent(services, "", "  ")
	fmt.Print(string(output))
}
