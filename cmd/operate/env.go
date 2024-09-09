package operate

import (
	"encoding/json"
	"fmt"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

var envName string
var data string
var operateEnvClient = service.Environment{}

var operateEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "operate env",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `odin operate env [Options]`,
	Run: func(cmd *cobra.Command, args []string) {
		executeOperate(cmd)
	},
}

func init() {
	operateEnvCmd.Flags().StringVar(&envName, "name", "", "name of the component")
	operateEnvCmd.Flags().StringVar(&data, "data", "{}", "options of the operation in JSON format")
	if err := operateEnvCmd.MarkFlagRequired("name"); err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	if err := operateEnvCmd.MarkFlagRequired("data"); err != nil {
		log.Fatal("Error marking 'data' flag as required:", err)
	}

	operateCmd.AddCommand(operateEnvCmd)
}

func executeOperate(cmd *cobra.Command) {

	ctx := cmd.Context()
	var optionsData map[string]interface{}

	err := json.Unmarshal([]byte(data), &optionsData)
	if err != nil {
		log.Fatal("Unable to parse JSON data " + err.Error())
	}

	dataStruct, err := structpb.NewStruct(optionsData)
	if err != nil {
		log.Fatal("error converting JSON to structpb.Struct: ", err)
	}
	//call operate env client
	response, err := operateEnvClient.OperateEnvironment(&ctx, &environment.UpdateEnvironmentRequest{
		EnvName: envName,
		Data:    dataStruct,
	})

	if err != nil {
		log.Fatal("Failed to update environment.", err)
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(response, outputFormat)

}

func writeOutput(response *environment.UpdateEnvironmentResponse, format string) {

	switch format {
	case constant.TEXT:
		writeAsText(response)
	case constant.JSON:
		writeAsJSON(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func writeAsText(response *environment.UpdateEnvironmentResponse) {
	fmt.Print(string("autoDeletionTime updated successfully"))
}

func writeAsJSON(response *environment.UpdateEnvironmentResponse) {
	var environments []map[string]interface{}
	output, _ := json.MarshalIndent(environments, "", "  ")
	fmt.Print(string(output))
}
