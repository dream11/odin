package util

import (
	"fmt"
	"net"
	"strings"
)

// GetDefaultMACAddress returns the MAC address of the first non-loopback network interface
// with a valid MAC address. Returns empty string if no suitable interface is found.
func GetDefaultMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Error("failed to get network interfaces: %w", err)
	}

	for _, iface := range interfaces {
		// Skip loopback, down interfaces, and interfaces without hardware address
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		// Check if interface has a valid MAC address
		if len(iface.HardwareAddr) > 0 {
			macAddr := iface.HardwareAddr.String()
			if macAddr != "" && macAddr != "00:00:00:00:00:00" {
				return macAddr, nil
			}
		}
	}

	return "", fmt.Error("no suitable network interface with valid MAC address found")
}

// FormatMACAddress ensures the MAC address is in the standard format (lowercase with colons)
func FormatMACAddress(mac string) string {
	// Remove any whitespace
	mac = strings.TrimSpace(mac)

	// Convert to lowercase
	mac = strings.ToLower(mac)

	// If already in correct format, return as-is
	if strings.Count(mac, ":") == 5 {
		return mac
	}

	// If using dashes, convert to colons
	if strings.Count(mac, "-") == 5 {
		return strings.ReplaceAll(mac, "-", ":")
	}

	// If no separators, add colons every 2 characters
	if len(mac) == 12 {
		var result strings.Builder
		for i, char := range mac {
			if i > 0 && i%2 == 0 {
				result.WriteString(":")
			}
			result.WriteRune(char)
		}
		return result.String()
	}

	return mac
}

// ValidateMACAddress checks if the provided string is a valid MAC address
func ValidateMACAddress(mac string) bool {
	_, err := net.ParseMAC(mac)
	return err == nil
}
