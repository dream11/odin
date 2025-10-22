package operate

import (
	"context"
	"encoding/json"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

var serviceClient = service.Service{}
var operateServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "operate service",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `odin operate service [Options]`,
	Run: func(cmd *cobra.Command, args []string) {
		executeOperateService(cmd)
	},
}

func init() {
	operateServiceCmd.Flags().StringVar(&name, "name", "", "name of the service")
	operateServiceCmd.Flags().StringVar(&env, "env", "", "name of the environment in which the service is deployed")
	operateServiceCmd.Flags().StringVar(&operation, "operation", "", "name of the operation to performed on the service")
	operateServiceCmd.Flags().StringVar(&options, "options", "{}", "options of the operation in JSON format")
	operateServiceCmd.Flags().StringVar(&file, "file", "", "path of the file which contains the options for the operation in JSON format")
	if err := operateServiceCmd.MarkFlagRequired("name"); err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	if err := operateServiceCmd.MarkFlagRequired("operation"); err != nil {
		log.Fatal("Error marking 'operation' flag as required:", err)
	}
	operateCmd.AddCommand(operateServiceCmd)
}

func executeOperateService(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)

	ctx := cmd.Context()
	traceID := util.GenerateTraceID()
	contextWithTrace := context.WithValue(ctx, constant.TraceIDKey, traceID)
	verboseEnabled, err := cmd.Flags().GetBool(constant.VerboseFlag)
	if err != nil {
		log.Fatal(err)
	}

	contextWithTrace = context.WithValue(contextWithTrace, constant.VerboseEnabledKey, verboseEnabled)

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

	//call operate service client
	err = serviceClient.OperateService(&contextWithTrace, &serviceProto.OperateServiceRequest{
		EnvName:              env,
		ServiceName:          name,
		IsComponentOperation: false,
		Operation:            operation,
		Config:               config,
	})

	if err != nil {
		util.LogGrpcError(err, "\nFailed to operate on service: ")
	}

}
