package util

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"os"
	"strings"
	"time"

	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
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
	message := fmt.Sprintf("\n Service %s %s", response.ServiceStatus.ServiceAction, response.ServiceStatus)
	for _, compMessage := range response.ComponentsStatus {
		message += fmt.Sprintf("\n Component %s %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus, compMessage.Error)
	}
	return message
}

// FormatToHumanReadableDuration takes a date-time string representing the last deployment time, and returns a human-readable string representing the duration since the last deployment
func FormatToHumanReadableDuration(inputDateTime string) string {
	// Parse the date-time string
	layout := "02-01-2006 15:04:05:0000"
	parsedTime, err := time.Parse(layout, inputDateTime)
	if err != nil {
		return fmt.Sprintf("Failed to parse input time: %v", err)
	}

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

// contains checks if a string is present in an array of strings
func contains(str string, arr []string) bool {
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
		if contains(compMessage.ComponentName, components) {
			message += fmt.Sprintf("\n Component %s %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus, compMessage.Error)
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
