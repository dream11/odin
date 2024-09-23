package describe

import (
	"encoding/json"
	"fmt"
	serviceBackend "github.com/dream11/odin/internal/service"
	comp "github.com/dream11/odin/proto/gen/go/dream11/od/component/v1"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var componentName string
var componentVersion string

var componentClient = serviceBackend.Component{}
var componentCmd = &cobra.Command{
	Use:   "component-type",
	Short: "Describe component-type",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `Describe component-type schema and defaults with all flavours and operations`,
	Run: func(cmd *cobra.Command, args []string) {
		executeDescribeComponentType(cmd)
	},
}

func init() {
	componentCmd.Flags().StringVar(&componentName, "name", "", "name of the service")
	componentCmd.Flags().StringVar(&componentVersion, "version", "", "version of the service")
	err := componentCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	err = componentCmd.MarkFlagRequired("version")
	if err != nil {
		log.Fatal("Error marking 'version' flag as required:", err)
	}
	describeCmd.AddCommand(componentCmd)
}

func executeDescribeComponentType(cmd *cobra.Command) {
	params := map[string]string{
		"version": componentVersion,
	}
	ctx := cmd.Context()
	response, err := componentClient.DescribeComponentType(&ctx, &comp.DescribeComponentTypeRequest{
		ComponentType: componentName,
		Params: params,
	})

	if err != nil {
		log.Fatal("Failed to describe service: ", err)
	}

	writeAsJSONDescribeComponentType(response)
}

func writeAsJSONDescribeComponentType(response *comp.DescribeComponentTypeResponse) {
	output, err := json.MarshalIndent(response.Component, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	fmt.Println(string(output))
}

