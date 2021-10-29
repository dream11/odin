package commands

import (
	"fmt"

	"github.com/dream11/d11-cli/d11cli"
)

// Command verbs
type command struct {
	Create      bool    // Create a resource record
	Delete      bool    // Delete a resource record
	Describe    bool    // Describe a resource
	List        bool    // List the resources
	Deploy      bool    // Deploy resource
	Destroy     bool    // Destroy the deployed resource
}

func commandHelper(verb, resource string, options []string) string {
	var opts string
	if len(options) > 0 {
		opts = "[Options]\n\nOptions:\n"
	}

	for _, opt := range(options) {
		opts = opts + fmt.Sprintf("\t%s\n", opt)
	}
	return fmt.Sprintf("Usage: %s %s %s %s", d11cli.App.Name, verb, resource, opts)
}

func defaultHelper() string {
	return fmt.Sprintf("Usage: %s --help", d11cli.App.Name)
}