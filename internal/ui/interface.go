package ui

import (
	"os"

	"github.com/mitchellh/cli"
)

type ui struct {
	cli.Ui
}

// interact with cli
// take inputs, secret inputs, throw outputs/error etc.
// for more, refer https://github.com/mitchellh/cli/blob/master/ui.go
var userInterface *cli.PrefixedUi = &cli.PrefixedUi{
	AskPrefix:       "",
	AskSecretPrefix: "(Secret) ",
	OutputPrefix:    "",
	InfoPrefix:      "",
	ErrorPrefix:     "[ ERROR ] ",
	WarnPrefix:      "[ WARNING ] ",
	Ui: &ui{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
			Reader:      os.Stdin,
		},
	},
}
