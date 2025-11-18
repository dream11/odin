package util

import (
	"fmt"
	"net"
)

// GetDefaultMACAddress returns the MAC address of the first non-loopback network interface
// with a valid MAC address. Returns empty string if no suitable interface is found.
func GetDefaultMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to get network interfaces: %w", err)
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

	return "", fmt.Errorf("no suitable network interface with valid MAC address found")
}
