package commandline

import (
	"os"

	"github.com/mitchellh/cli"
)

type ui struct {
	cli.Ui
}

// interact with cli
// take inputs, secret inputs, throw outputs/error etc
// for more, refer https://github.com/mitchellh/cli/blob/master/ui.go
var Interface *cli.PrefixedUi = &cli.PrefixedUi{
	AskPrefix:       "Input:",
	AskSecretPrefix: "Input(secret):",
	OutputPrefix:    "",
	InfoPrefix:      "[ INFO ] ",
	ErrorPrefix:     "[ ERROR ] ",
	WarnPrefix:      "[ WARNING ] ",
	Ui: &ui{
		&cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
			Reader:      os.Stdin,
		},
	},
}
