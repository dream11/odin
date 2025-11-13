package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/dream11/odin/internal/ui"
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
)

// SplitProviderAccount splits string into list of cloud provider accounts
func SplitProviderAccount(providerAccounts string) []string {
	if providerAccounts == "" {
		return nil
	}
	return strings.Split(providerAccounts, ",")
}

// GenerateResponseMessage generate response message from ServiceResponse
func GenerateResponseMessage(response *v1.ServiceResponse) string {
	var builder strings.Builder

	// Write main service status
	builder.WriteString(fmt.Sprintf(
		"\nService: %-15s Version: %-15s Action: %-10s Status: %-10s",
		response.GetName(),
		response.GetVersion(),
		response.GetServiceStatus().GetServiceAction(),
		response.GetServiceStatus().GetServiceStatus(),
	))

	// Header for components
	builder.WriteString("\n  Component Name         Action     Status")

	// Write each component in aligned tabular format
	for _, comp := range response.ComponentsStatus {
		builder.WriteString(fmt.Sprintf(
			"\n  %-22s %-10s %-10s",
			comp.ComponentName,
			comp.ComponentAction,
			comp.ComponentStatus,
		))
	}

	// Log failures
	for _, comp := range response.ComponentsStatus {
		if comp.GetComponentStatus() == "FAILED" {
			log.Error(fmt.Sprintf(
				"Component %s %s %s - Error: %s\n\n",
				comp.GetComponentName(),
				comp.GetComponentAction(),
				comp.GetComponentStatus(),
				comp.GetError(),
			))
		}
	}

	return builder.String()
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

// GenerateTraceID generates a trace id
func GenerateTraceID() string {
	traceID := uuid.New().String()
	log.Infof("Generated trace ID: %s\n", traceID)
	return traceID
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

// AskForConfirmation asks for confirmation before proceeding with the operation
func AskForConfirmation(expectedValue, consentMessage string) {
	inputHandler := ui.Input{}
	val, err := inputHandler.Ask(consentMessage)
	if err != nil {
		log.Fatal(err.Error())
	}
	if val != expectedValue {
		log.Fatal(fmt.Errorf("invalid input, aborting the operation"))
	}
}

// IsRetryable checks if the error is retryable
func IsRetryable(err error) bool {
	if errors.Is(err, context.Canceled) {
		return true
	}
	if err == io.EOF {
		return false
	}
	st, ok := status.FromError(err)
	return ok && (st.Code() == codes.Unavailable || (st.Code() == codes.Internal && strings.Contains(st.Message(), "RST_STREAM")))
}

// LogGrpcError log grpc error
func LogGrpcError(err error, prefix string) {
	st, ok := status.FromError(err)
	if ok {
		log.Error(prefix + st.Message())
	} else {
		log.Error(prefix + err.Error())
	}
}
