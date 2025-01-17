package describe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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
	environmentCmd.Flags().StringVar(&component, "component", "", "component of the service deployed in this environment")
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
		printEnvInfo(response)
	case constant.JSON:
		writeAsJSONEnvResponse(response)
	default:
		log.Fatal("Unknown output format: ", format)
	}
}

func printEnvInfo(response *environment.DescribeEnvironmentResponse) {
	env := response.Environment

	// Extracting necessary fields
	name := env.GetName()
	team := "infrastructure engineering" // Assuming static value as it's not in the response
	envType := "load"                    // Assuming static value as it's not in the response
	state := env.GetStatus()
	autoDeletionTime := env.AutoDeletionTime.AsTime().String()
	cloudProviderAccounts := []string{}
	providerAccountCluster := map[string][]string{}
	for _, accountInfo := range env.AccountInformation {
		cloudProviderAccounts = append(cloudProviderAccounts, accountInfo.ProviderAccountName)
		providerAccountCluster[accountInfo.ProviderAccountName] = getClusterNames(accountInfo)
	}
	createdBy := env.GetCreatedBy()
	updatedBy := env.GetUpdatedBy()
	createdAt := env.CreatedAt.AsTime().String()
	updatedAt := env.UpdatedAt.AsTime().String()
	services := []string{}
	for _, svc := range env.Services {
		if serviceName == "" {
			services = append(services, fmt.Sprintf("    - name: %s\n      version: %s", *svc.Name, *svc.Version))
		} else {
			if serviceName == *svc.Name {
				var customServicesOp = []string{}
				customServicesOp = append(customServicesOp, fmt.Sprintf("    - name: %s\n      version: %s\n", *svc.Name, *svc.Version))
				customServicesOp = append(customServicesOp, "      components: \n")
				componentBytes, err := json.MarshalIndent(svc.Components, "", "  ")
				if err != nil {
					log.Fatal("Failed to marshal services summary: ", err)
				}
				var formatedComponentData = string(componentBytes)
				formatedComponentData, _ = ConvertJSONToYAML(formatedComponentData)
				lines := strings.Split(formatedComponentData, "\n")

				for i, line := range lines {
					lines[i] = "\t" + line
				}
				formatedComponentData = strings.Join(lines, "\n")
				// Add two tabs before each line in the string
				customServicesOp = append(customServicesOp, formatedComponentData)

				services = append(services, strings.Join(customServicesOp, ""))
			}

		}
	}

	// Formatting and printing the information
	fmt.Printf("Describing Env: %s\n\n", name)
	fmt.Printf("name: %s\n", name)
	fmt.Printf("team: %s\n", team)
	fmt.Printf("envType: %s\n", envType)
	fmt.Printf("state: %s\n", state)
	fmt.Printf("autoDeletionTime: \"%s\"\n", autoDeletionTime)
	fmt.Printf("cloudProviderAccounts:\n")
	for _, account := range cloudProviderAccounts {
		fmt.Printf("    - %s\n", account)
	}
	fmt.Printf("cluster:\n")
	for account, clusters := range providerAccountCluster {
		fmt.Printf("    - %s\n", account)
		for _, cluster := range clusters {
			fmt.Printf("        - %s\n", cluster)
		}

	}
	fmt.Printf("createdBy: %s\n", createdBy)
	fmt.Printf("updatedBy: %s\n", updatedBy)
	fmt.Printf("createdAt: \"%s\"\n", createdAt)
	fmt.Printf("updatedAt: \"%s\"\n", updatedAt)
	fmt.Printf("services:\n%s\n", strings.Join(services, "\n"))
}

// ConvertJSONToYAML takes a JSON string as input and returns a formatted YAML string
func ConvertJSONToYAML(jsonStr string) (string, error) {
	// Unmarshal the JSON into a generic structure
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Marshal the structure into YAML
	yamlData, err := yaml.Marshal(jsonData)
	if err != nil {
		return "", fmt.Errorf("failed to convert to YAML: %v", err)
	}

	// Return the YAML string
	return string(yamlData), nil
}

func findValueByKey(val interface{}, key string) string {
	switch v := val.(type) {
	case map[string]interface{}: // If it's a map, check for the key
		if value, ok := v[key]; ok {
			return fmt.Sprintf("%v", value)
		}
		// Recurse through nested maps or slices
		for _, subVal := range v {
			return findValueByKey(subVal, key)
		}
	case []interface{}: // If it's a slice, recurse for each element
		for _, item := range v {
			return findValueByKey(item, key)
		}
	}
	return "" // Key not found
}

func getClusterNames(information *v1.AccountInformation) []string {
	clusterNames := []string{}
	for _, service := range information.ServiceAccountsSnapshot.Account.Services {
		if service.Category == "KUBERNETES" {
			for key, val := range service.GetData().AsMap() {
				if key == "clusters" {
					clusterNames = append(clusterNames, findValueByKey(val, "name"))
				}
			}
		}
	}
	return clusterNames
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
			"status":  svc.Status,
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
