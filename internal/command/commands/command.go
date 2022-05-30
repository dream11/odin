package commands

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	odin "github.com/dream11/odin/app"
	"github.com/dream11/odin/internal/ui"
	"github.com/olekukonko/tablewriter"
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
// func commandHelper(verb, resource string, options []string) string {
// 	var opts string
// 	if len(options) > 0 {
// 		opts = "[Options]\n\nOptions:\n"
// 	}

// 	for _, opt := range options {
// 		opts = opts + fmt.Sprintf("\t%s\n", opt)
// 	}
// 	return fmt.Sprintf("Usage: %s %s %s %s", odin.App.Name, verb, resource, opts)
// }

func commandHelper(verb, resource string, description string, options []Options) string {
	buf := new(bytes.Buffer)

	// Write description to buffer
	if description != "" {
		buf.WriteString("\n\nDescription:\n")
		buf.WriteString(description + "\n")
	}

	// Write options to buffer
	if len(options) > 0 {
		buf.WriteString("[Options]\n\nOptions:\n")
		table := tablewriter.NewWriter(buf)
		table.SetRowLine(false)
		table.SetColumnSeparator("")
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
		table.SetAutoWrapText(false)
		for _, opt := range options {
			table.Append([]string{opt.Flag, opt.Description})
		}
		table.Render()
	}

	return fmt.Sprintf("Usage: %s %s %s %s", odin.App.Name, verb, resource, buf)
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
