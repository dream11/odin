package update

import (
	"encoding/json"
	"fmt"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	fileUtil "github.com/dream11/odin/pkg/util"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

var envName string
var data string
var file string
var updateEnvClient = service.Environment{}

var updateEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "update env",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `odin update env [Options]`,
	Run: func(cmd *cobra.Command, args []string) {
		executeupdate(cmd)
	},
}

func init() {
	updateEnvCmd.Flags().StringVar(&envName, "name", "", "name of the component")
	updateEnvCmd.Flags().StringVar(&data, "data", "{}", "options of the operation in JSON format")
	updateEnvCmd.Flags().StringVar(&file, "file", "", "path of the file which contains the options for the operation in JSON format")
	if err := updateEnvCmd.MarkFlagRequired("name"); err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}

	updateCmd.AddCommand(updateEnvCmd)
}

func executeupdate(cmd *cobra.Command) {

	ctx := cmd.Context()
	var optionsData map[string]interface{}

	isOptionsPresent := data != "{}"
	isFilePresent := len(file) > 0

	if isOptionsPresent && isFilePresent {
		log.Fatal("You can provide either --data or --file but not both")
	}
	if !isOptionsPresent && !isFilePresent {
		log.Fatal("Either --data or --file is mandatory")
	}

	if isFilePresent {
		parsedConfig, err := fileUtil.ParseFile(file)
		if err != nil {
			log.Fatal("Error while parsing file " + file + " : " + err.Error())
		}
		optionsData = parsedConfig.(map[string]interface{})
	} else {
		err := json.Unmarshal([]byte(data), &optionsData)
		if err != nil {
			log.Fatal("Unable to parse JSON data " + err.Error())
		}
	}

	dataStruct, err := structpb.NewStruct(optionsData)
	if err != nil {
		log.Fatal("error converting JSON to structpb.Struct: ", err)
	}
	//call update env client
	response, err := updateEnvClient.UpdateEnvironment(&ctx, &environment.UpdateEnvironmentRequest{
		EnvName: envName,
		Data:    dataStruct,
	})

	if err != nil {
		log.Fatal("Failed to update environment ", err.Error())
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
	fmt.Print(string("Environment updated successfully"))
}

func writeAsJSON(response *environment.UpdateEnvironmentResponse) {
	var environments []map[string]interface{}
	output, _ := json.MarshalIndent(environments, "", "  ")
	fmt.Print(string(output))
}
