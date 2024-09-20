package describe

import (
	"encoding/json"
	"fmt"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/table"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string


var environmentClient = service.Environment{}
var environmentCmd = &cobra.Command{
	Use:   "env",
	Short: "Describe environments",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: `Describe  environment details`,
	Run: func(cmd *cobra.Command, args []string) {
		executeEnv(cmd)
	},
}

func init() {
	environmentCmd.Flags().StringVar(&name, "name", "", "name of the environment")
	environmentCmd.Flags().StringVar(&serviceName, "service", "", "service deployed in this environment")
	environmentCmd.Flags().StringVar(&component, "component", "", "component of the service")
	describeCmd.AddCommand(environmentCmd)
}

func executeEnv(cmd *cobra.Command) {
	ctx := cmd.Context()
	params := map[string]string{}

	if serviceName != "" {
		params["service"] = serviceName
	}
	if component != "" {
		params["component"] = component
	}
	response, err := environmentClient.DescribeEnvironment(&ctx, &environment.DescribeEnvironmentRequest{
		Params:  params,
		EnvName: name,
	})

	if err != nil {
		log.Fatal("Failed to describe environment ", err)
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(response, outputFormat)
}

func writeOutput(response *environment.DescribeEnvironmentResponse, format string) {

	switch format {
	case constant.TEXT:
		writeAsTextEnvResponse(response)
	case constant.JSON:
		writeAsJSONEnvResponse(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func writeAsTextEnvResponse(response *environment.DescribeEnvironmentResponse) {

	tableHeaders := []string{"Name",
		"team",
		"state",
		"autoDeletionTime",
		"cloudProviderAccounts",
		"createdBy",
		"updatedBy",
		"createdAt",
		"updatedAt",
		"services"}
	var tableData [][]interface{}
	env := response.Environment
	var accountInfoList []string
	for _, accountInfo := range env.AccountInformation {
		accountInfoList = append(accountInfoList, accountInfo.ProviderAccountName)
	}
	accountInfoListJSON, err := json.Marshal(accountInfoList)
	if err != nil {
		log.Fatal("Failed to marshal account info list: ", err)
	}

	var servicesSummary []map[string]interface{}
	for _, svc := range env.Services {
		serviceMap := map[string]interface{}{
			"name":    svc.Name,
			"version": svc.Version,
		}
		if len(svc.Components) > 0 {
			serviceMap["components"] = svc.Components
		}
		servicesSummary = append(servicesSummary, serviceMap)
	}
	servicesSummaryJSON, err := json.Marshal(servicesSummary)
	if err != nil {
		log.Fatal("Failed to marshal services summary: ", err)
	}

	tableData = append(tableData, []interface{}{
		*env.Name,
		*env.Status,
		env.AutoDeletionTime.AsTime().String(),
		string(accountInfoListJSON),
		*env.CreatedBy,
		*env.UpdatedBy,
		env.CreatedAt.AsTime().String(),
		env.UpdatedAt.AsTime().String(),
		string(servicesSummaryJSON),
	})

	table.Write(tableHeaders, tableData)
}

func writeAsJSONEnvResponse(response *environment.DescribeEnvironmentResponse) {
	var environments []map[string]interface{}
	env := response.Environment
	var accountInfoList []string
	for _, accountInfo := range env.AccountInformation {
		accountInfoList = append(accountInfoList, accountInfo.ProviderAccountName)
	}

	var servicesSummary []map[string]interface{}
	for _, svc := range env.Services {
		serviceMap := map[string]interface{}{
			"name":    svc.Name,
			"version": svc.Version,
		}
		if len(svc.Components) > 0 {
			serviceMap["components"] = svc.Components
		}
		servicesSummary = append(servicesSummary, serviceMap)
	}

	environments = append(environments, map[string]interface{}{
		"name":                  env.Name,
		"state":                 env.Status,
		"autoDeletionTime":      env.AutoDeletionTime.AsTime().String(),
		"cloudProviderAccounts": accountInfoList,
		"createdBy":             env.CreatedBy,
		"updatedBy":             env.UpdatedBy,
		"createdAt":             env.CreatedAt.AsTime().String(),
		"updatedAt":             env.UpdatedAt.AsTime().String(),
		"services":              servicesSummary,
	})

	output, _ := json.MarshalIndent(environments, "", "  ")
	fmt.Print(string(output))
}
