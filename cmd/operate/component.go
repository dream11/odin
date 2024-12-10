package operate

import (
	"encoding/json"
	"github.com/dream11/odin/pkg/util"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/config"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

var name string
var serviceName string
var env string
var operation string
var options string
var file string
var componentClient = service.Component{}
var operateComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "operate component",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `odin operate component [Options]`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	operateComponentCmd.Flags().StringVar(&name, "name", "", "name of the component")
	operateComponentCmd.Flags().StringVar(&serviceName, "service", "", " name of the service in which the component is deployed")
	operateComponentCmd.Flags().StringVar(&env, "env", "", "name of the environment in which the service is deployed")
	operateComponentCmd.Flags().StringVar(&operation, "operation", "", "name of the operation to performed on the component")
	operateComponentCmd.Flags().StringVar(&options, "options", "{}", "options of the operation in JSON format")
	operateComponentCmd.Flags().StringVar(&file, "file", "", "path of the file which contains the options for the operation in JSON format")
	if err := operateComponentCmd.MarkFlagRequired("name"); err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	if err := operateComponentCmd.MarkFlagRequired("service"); err != nil {
		log.Fatal("Error marking 'service' flag as required:", err)
	}
	if err := operateComponentCmd.MarkFlagRequired("operation"); err != nil {
		log.Fatal("Error marking 'operation' flag as required:", err)
	}
	operateCmd.AddCommand(operateComponentCmd)
}

func execute(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)

	ctx := cmd.Context()
	traceID := util.GenerateTraceID()
	//validate the variables
	var optionsData map[string]interface{}

	isOptionsPresent := options != "{}"
	isFilePresent := len(file) > 0

	if isOptionsPresent && isFilePresent {
		log.Fatal("You can provide either --options or --file but not both")
	}

	if isFilePresent {
		parsedConfig, err := util.ParseFile(file)
		if err != nil {
			log.Fatal("Error while parsing file " + file + " : " + err.Error())
		}
		optionsData = parsedConfig.(map[string]interface{})
	} else {
		err := json.Unmarshal([]byte(options), &optionsData)
		if err != nil {
			log.Fatal("Unable to parse JSON data " + err.Error())
		}
	}

	config, err := structpb.NewStruct(optionsData)
	if err != nil {
		log.Fatal("error converting JSON to structpb.Struct: ", err)
	}
	//call operate component client
	err = componentClient.OperateComponent(&ctx, &serviceProto.OperateServiceRequest{
		EnvName:              env,
		ServiceName:          serviceName,
		ComponentName:        name,
		IsComponentOperation: true,
		Operation:            operation,
		Config:               config,
	}, traceID)

	if err != nil {
		log.Fatal("Failed to operate on component", err)
	}

}
