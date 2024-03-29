package commands

import (
	"flag"

	"github.com/dream11/odin/api/label"
	"github.com/dream11/odin/internal/backend"
	"github.com/dream11/odin/internal/constant"
	"github.com/dream11/odin/pkg/table"
)

var labelClient backend.Label

type Label command

// Run : implements the actual functionality of the command
func (l *Label) Run(args []string) int {
	// Define flag set
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)
	// create flags
	name := flagSet.String("name", "", "name of label")
	versionCardinalityGreaterThanOne := flagSet.Bool("cardinality", false, "whether multiple versions of a service can have this label")

	err := flagSet.Parse(args)
	if err != nil {
		l.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if l.Create {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		if len(emptyParameters) == 0 {
			l.Logger.Info("Creating label: " + *name)
			labelData := label.Label{
				Name:                             *name,
				VersionCardinalityGreaterThanOne: *versionCardinalityGreaterThanOne,
			}

			labelClient.CreateLabel(labelData)
			l.Logger.Success("Successfully created label: " + *name)
			return 0
		}

		l.Logger.Error(constant.LABEL_NAME_OPTION)

		return 1
	}

	if l.List {
		l.Logger.Info("Listing all label(s)")
		labelList, err := labelClient.ListLables()

		if err != nil {
			l.Logger.Error(err.Error())
			return 1
		}

		tableHeaders := []string{"Name", "cardinality"}
		var tableData [][]interface{}

		for _, label := range labelList {
			tableData = append(tableData, []interface{}{
				label.Name,
				label.VersionCardinalityGreaterThanOne,
			})
		}
		table.Write(tableHeaders, tableData)

		return 0
	}

	if l.Delete {
		emptyParameters := emptyParameters(map[string]string{"--name": *name})
		l.Logger.Info("Deleting label: " + *name)
		if len(emptyParameters) == 0 {
			labelClient.DeleteLabel(*name)
			l.Logger.Success("Successfully deleted label: " + *name)
			return 0
		}

		l.Logger.Error(constant.LABEL_NAME_OPTION)
		return 1
	}
	l.Logger.Error("Not a valid command")
	return 127
}

// Help : returns an explanatory string
func (l *Label) Help() string {
	if l.Create {
		return commandHelper("create", "label", "", []Options{
			{Flag: "--name", Description: "name of the label"},
			{Flag: "--cardinality", Description: "whether the label can be linked to multiple versions of a service (optional)"},
		})
	}

	if l.List {
		return commandHelper("list", "label", "", []Options{})
	}

	if l.Delete {
		return commandHelper("delete", "label", "", []Options{
			{Flag: "--name", Description: "name of label to delete"},
		})
	}

	return defaultHelper()
}

// Synopsis : returns a brief helper text for the command's verbs
func (l *Label) Synopsis() string {
	if l.Create {
		return "create a label"
	}

	if l.List {
		return "list all labels"
	}

	if l.Delete {
		return "delete a label"
	}

	return defaultHelper()
}
