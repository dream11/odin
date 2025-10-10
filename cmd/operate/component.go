package operate

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	"github.com/dream11/odin/pkg/util"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/fatih/color"
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

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case []interface{}:
		strSlice := make([]string, len(v))
		for i, elem := range v {
			strSlice[i] = formatValue(elem)
		}
		if len(strSlice) == 1 {
			return strSlice[0]
		}
		return "[" + strings.Join(strSlice, ", ") + "]"
	case map[string]interface{}, struct{}:
		jsonBytes, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		return string(jsonBytes)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func execute(cmd *cobra.Command) {
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
	//call operate component client
	if operation == "redeploy" {
		diffValues, err := componentClient.CompareOperationChanges(&contextWithTrace, &serviceProto.OperateComponentDiffRequest{
			EnvName:       env,
			ServiceName:   serviceName,
			ComponentName: name,
			OperationName: operation,
			Config:        config,
		})
		if err != nil {
			log.Fatal("Failed to compare operation changes\n", err)
		}
		oldComponentValues := diffValues.OldValues
		newComponentValues := diffValues.NewValues

		if oldComponentValues != nil && len(oldComponentValues.Fields) > 0 {
			log.Info("\nBelow changes will happen after this operation:\n")
			tableHeaders := []string{"Component Name", "Config", "Old Value", "New Value"}
			var tableData [][]interface{}

			componentName := name
			flatendOldComponentValues := flattenMap(oldComponentValues.AsMap(), "")
			flatendNewComponentValues := flattenMap(newComponentValues.AsMap(), "")

			keys := make([]string, 0, len(flatendNewComponentValues))
			for k := range flatendOldComponentValues {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, key := range keys {
				oldValue := flatendOldComponentValues[key]
				newValue := flatendNewComponentValues[key]
				if fmt.Sprintf("%v", oldValue) != fmt.Sprintf("%v", newValue) {
					var oldValueString string
					var newValueString string
					oldValueString = strings.Join(applyColorToLines(formatValue(oldValue), color.RedString), "\n")
					newValueString = strings.Join(applyColorToLines(formatValue(newValue), color.GreenString), "\n")
					tableData = append(tableData, []interface{}{
						componentName,
						key,
						oldValueString,
						newValueString,
					})
				}
			}
			table.Write(tableHeaders, tableData)
		}

		var message string
		if oldComponentValues == nil || len(oldComponentValues.Fields) == 0 {
			message = "\nNo changes from previous deployment. Do you want to continue? [y/n]:"
		} else {
			message = "\nDo you want to proceed with the above command? [y/n]:"
		}
		allowedInputsSlice := []string{"y", "n"}
		allowedInputs := make(map[string]struct{}, len(allowedInputsSlice))
		for _, input := range allowedInputsSlice {
			allowedInputs[input] = struct{}{}
		}

		inputHandler := ui.Input{}
		val, err := inputHandler.AskWithConstraints(message, allowedInputs)

		if err != nil {
			log.Fatal(err.Error())
		}

		if val != "y" {
			log.Info("Aborting the operation")
			return
		}

	}
	err = componentClient.OperateComponent(&contextWithTrace, &serviceProto.OperateServiceRequest{
		EnvName:              env,
		ServiceName:          serviceName,
		ComponentName:        name,
		IsComponentOperation: true,
		Operation:            operation,
		Config:               config,
	})

	if err != nil {
		util.LogGrpcError(err, "Failed to operate on component: ")
	}

}

func flattenMap(m map[string]interface{}, prefix string) map[string]interface{} {
	flattened := make(map[string]interface{})
	for k, v := range m {
		key := prefix + k
		if prefix != "" {
			key = prefix + "." + k
		}
		if vm, ok := v.(map[string]interface{}); ok {
			flattenedMap := flattenMap(vm, key)
			if len(flattenedMap) == 0 {
				flattened[key] = make(map[string]interface{})
			}
			for fk, fv := range flattenedMap {
				flattened[fk] = fv
			}
		} else {
			flattened[key] = v
		}
	}
	return flattened
}

func applyColorToLines(value string, colorFunc func(format string, a ...interface{}) string) []string {
	lines := strings.Split(value, "\n")
	for i, line := range lines {
		lines[i] = colorFunc(line)
	}
	return lines
}
