package util

import (
	"fmt"
	"net"
	"strings"

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
