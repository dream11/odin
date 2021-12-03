package commands

import (
	"fmt"

	odin "github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/ui"
)

/*
command : interface for resources
The verbs can be associated with any resource
*/
type command struct {
	Create   bool // Create a resource record
	Delete   bool // Delete a resource record
	Update   bool // Update a resource record
	Describe bool // Describe a resource
	Label    bool // Label a resource
	List     bool // List the resources
	Status   bool // current Status of resource
	Logs     bool // execution Logs of resource
	Deploy   bool // Deploy resource
	Destroy  bool // Destroy the deployed resource

	Logger ui.Logger // Use this to log messages
	Input  ui.Input  // Use this to take inputs
}

// help text generator
func commandHelper(verb, resource string, options []string) string {
	var opts string
	if len(options) > 0 {
		opts = "[Options]\n\nOptions:\n"
	}

	for _, opt := range options {
		opts = opts + fmt.Sprintf("\t%s\n", opt)
	}
	return fmt.Sprintf("Usage: %s %s %s %s", odin.App.Name, verb, resource, opts)
}

func defaultHelper() string {
	return fmt.Sprintf("Usage: %s --help", odin.App.Name)
}

// validate if any value provided is empty or not
func emptyParameterValidation(params []string) bool {
	for _, val := range params {
		if len(val) == 0 {
			return false
		}
	}
	return true
}
