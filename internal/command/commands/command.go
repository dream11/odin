package commands

import (
	"fmt"
	"sort"
	"strings"

	odin "github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/ui"
)

/*
command : interface for resources
The verbs can be associated with any resource
*/
type command struct {
	Create          bool // Create a resource record
	Delete          bool // Delete a resource record
	Update          bool // Update a resource record
	Describe        bool // Describe a resource
	Label           bool // Label a resource
	List            bool // List the resources
	Status          bool // current Status of resource
	Logs            bool // execution Logs of resource
	Deploy          bool // Deploy resource
	Undeploy        bool // Undeploy resource
	Destroy         bool // Destroy the deployed resource
	GetHistory      bool // Get changelog of resource
	DescribeHistory bool // Describe a changelog of resource
	Generate        bool // Generate resources
	Unlabel         bool // Unlabel a resource
	CreateDeploy    bool // Create and deploy a resource record

	Logger ui.Logger // Use this to log messages
	Input  ui.Input  // Use this to take inputs
}

type Options struct {
	Flag        string
	Description string
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

func commandHelper2(verb, resource string, options []Options) string {
	var opts string
	if len(options) > 0 {
		opts = "[Options]\n\nOptions:\n"
	}

	for _, opt := range options {
		opts = opts + fmt.Sprintf("\t%s:\t%s\n", opt.Flag, opt.Description)
	}
	return fmt.Sprintf("Usage: %s %s %s %s", odin.App.Name, verb, resource, opts)
}

func defaultHelper() string {
	return fmt.Sprintf("Usage: %s --help", odin.App.Name)
}

// get empty parameter list
func emptyParameters(params map[string]string) string {
	emptyParameters := []string{}
	for key, val := range params {
		if len(val) == 0 {
			emptyParameters = append(emptyParameters, key)
		}
	}
	sort.Strings(emptyParameters)
	return strings.Join(emptyParameters, ", ")
}
