package util

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

// SplitProviderAccount splits string into list of cloud provider accounts
func SplitProviderAccount(providerAccounts string) []string {
	if providerAccounts == "" {
		return nil
	}
	return strings.Split(providerAccounts, ",")
}

// IsIPAddress checks if given address is an IP address
func IsIPAddress(address string) bool {
	addr := net.ParseIP(address)
	return addr != nil
}

// GenerateResponseMessage generate response message from ServiceResponse
func GenerateResponseMessage(response *v1.ServiceResponse) string {
	message := fmt.Sprintf("\n Service: %s version: %s action: %s status: %s", response.Name, response.Version, response.ServiceStatus.ServiceAction, response.ServiceStatus.ServiceStatus)
	for _, compMessage := range response.ComponentsStatus {
		message += fmt.Sprintf("\n Component: %s action: %s status: %s ", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
	}
	return message
}

// GenerateServiceSetResponseMessage generate response message from ServiceSetResponse
func GenerateServiceSetResponseMessage(response *v1.DeployServiceSetServiceResponse) string {

	message := fmt.Sprintf("\n Service %s %s %s %s", response.ServiceIdentifier.ServiceName, response.ServiceIdentifier.ServiceVersion, response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
	var tableData [][]string
	row := []string{
		response.ServiceIdentifier.ServiceName,
		response.ServiceIdentifier.ServiceVersion,
		response.ServiceResponse.ServiceStatus.ServiceAction,
		response.ServiceResponse.ServiceStatus.ServiceStatus,
	}
	tableData = append(tableData, row)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Service Name", "Version", "Action", "Status", "Error"})
	table.AppendBulk(tableData)
	table.Render()
	return message

}

// FormatToHumanReadableDuration takes a date-time string representing the last deployment time, and returns a human-readable string representing the duration since the last deployment
func FormatToHumanReadableDuration(inputDateTime string) string {
	// Check if the input is a Unix timestamp prefixed by "seconds:"
	if strings.HasPrefix(inputDateTime, "seconds: ") {
		timestampStr := strings.TrimPrefix(inputDateTime, "seconds: ")
		timestampStr = strings.TrimSpace(timestampStr)

		// Parse the timestamp as an integer
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			return fmt.Sprintf("Failed to parse timestamp: %v", err)
		}

		// Convert the Unix timestamp to a time.Time object
		parsedTime := time.Unix(timestamp, 0)
		return calculateDuration(parsedTime)
	}

	// Handle the default case where input is in "DD-MM-YYYY HH:MM:SS:MS" format
	layout := "02-01-2006 15:04:05:0000"
	location, err := time.LoadLocation("Asia/Kolkata") // Adjust time zone as needed
	if err != nil {
		return fmt.Sprintf("Failed to load location: %v", err)
	}

	parsedTime, err := time.ParseInLocation(layout, inputDateTime, location)
	if err != nil {
		return fmt.Sprintf("Failed to parse input time: %v", err)
	}

	return calculateDuration(parsedTime)
}

func calculateDuration(parsedTime time.Time) string {
	// Calculate the duration
	duration := time.Since(parsedTime)
	// Handle negative durations
	if duration < 0 {
		duration = -duration
	}

	// Format the duration into a human-readable string
	if duration.Hours() >= 24*180 {
		months := int(duration.Hours() / (24 * 30))
		return fmt.Sprintf("%d months ago", months)
	} else if duration.Hours() >= 24 {
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	} else {
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		return fmt.Sprintf("%d hours %d minutes ago", hours, minutes)
	}
}

// Contains checks if a string is present in a slice of strings.
func Contains(str string, arr []string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

// GenerateResponseMessageComponentSpecific generate response message from ServiceResponse
func GenerateResponseMessageComponentSpecific(response *v1.ServiceResponse, components []string) string {
	message := fmt.Sprintf("\n Service %s %s", response.ServiceStatus.ServiceAction, response.ServiceStatus)
	for _, compMessage := range response.ComponentsStatus {
		if Contains(compMessage.ComponentName, components) {
			message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
		}
	}
	return message
}

// GenerateTraceID generates a trace id
func GenerateTraceID() string {
	return uuid.New().String()
}

// GetEnvOrDefault returns the value of an environment variable or a fallback value
func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
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

func GetHeaderText(name string, action string, status string, element string) string {
	var header strings.Builder
	var actionText string
	switch action {
	case "VALIDATE":
		actionText = "validation"
	case "DEPLOY":
		actionText = "deployment"
	case "UNDEPLOY":
		actionText = "un-deployment"
	}
	switch status {
	case "IN_PROGRESS":
		header.WriteString(fmt.Sprintf("%s: %s  %s in progress", element, name, actionText))
	case "SUCCESSFUL":
		header.WriteString(fmt.Sprintf("%s: %s  %s successful", element, name, actionText))
	case "FAILED":
		header.WriteString(fmt.Sprintf("%s: %s  %s failed", element, name, actionText))
	}
	return header.String()
}

func GetAvailableViewPortHeight(totalHeight int, headerHeight int, numberOfComponents int) int {
	bottomPadding := 12
	return totalHeight - ((headerHeight + 2) + numberOfComponents*(headerHeight+1)) - bottomPadding
}
