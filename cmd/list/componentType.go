package list

import (
	"encoding/json"
	"fmt"
	"github.com/dream11/odin/pkg/util"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	component "github.com/dream11/odin/proto/gen/go/dream11/od/component/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var componentTypeName string
var componentTypeVersion string

var componentTypeClient = service.Component{}
var componentTypeCmd = &cobra.Command{
	Use:   "component-type",
	Short: "List componentTypes",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `List of all componentTypes`,
	Run: func(cmd *cobra.Command, args []string) {
		componentExecute(cmd)
	},
}

func init() {
	componentTypeCmd.Flags().StringVar(&componentTypeName, "name", "", "name of the component type")
	componentTypeCmd.Flags().StringVar(&componentTypeVersion, "version", "", "version of the component type")
	listCmd.AddCommand(componentTypeCmd)
}

func componentExecute(cmd *cobra.Command) {
	ctx := cmd.Context()
	traceID := util.GenerateTraceID()
	params := make(map[string]string)

	// Add non-empty parameters to the map
	if componentTypeName != "" {
		params["name"] = componentTypeName
	}
	if componentTypeVersion != "" {
		params["version"] = componentTypeVersion
	}

	// Make the API call with the populated parameters
	response, err := componentTypeClient.ListComponentType(&ctx, &component.ListComponentTypeRequest{
		Params: params,
	}, traceID)

	if err != nil {
		log.Fatal("Failed to list component types ", err)
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	componentWriteOutput(response, outputFormat)
}

func componentWriteOutput(response *component.ListComponentTypeResponse, format string) {

	switch format {
	case constant.TEXT:
		componentWriteAsText(response)
	case constant.JSON:
		componentWriteAsJSON(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func componentWriteAsText(response *component.ListComponentTypeResponse) {
	tableHeaders := []string{"Name", "Version"}
	var tableData [][]interface{}
	for _, compType := range response.Components {
		tableData = append(tableData, []interface{}{
			compType.ComponentType,
			compType.ComponentVersion,
		})
	}

	table.Write(tableHeaders, tableData)
}

func componentWriteAsJSON(response *component.ListComponentTypeResponse) {
	var componentTypes []map[string]interface{}
	for _, compType := range response.Components {
		componentTypes = append(componentTypes, map[string]interface{}{
			"name":    compType.ComponentType,
			"version": compType.ComponentType,
		})
	}
	output, _ := json.MarshalIndent(componentTypes, "", "  ")
	fmt.Print(string(output))
}
