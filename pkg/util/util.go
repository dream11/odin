package util

import (
	"net"
	"sort"
	"strings"
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

// EmptyParameters get empty parameter list check if the map have any key with empty value if yes then return a comma seperated string of keys which have empty values
func EmptyParameters(params map[string]string) string {
	emptyParameters := []string{}
	for key, val := range params {
		if len(val) == 0 {
			emptyParameters = append(emptyParameters, key)
		}
	}
	sort.Strings(emptyParameters)
	return strings.Join(emptyParameters, ", ")
}
