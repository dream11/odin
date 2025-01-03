package util

import (
	"fmt"
	"net"
	"os"
	"strings"

	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/google/uuid"
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
