package cli

import (
	"os"

	"github.com/mitchellh/cli"
)

type ui struct {
	cli.Ui
}

func Ui() *cli.PrefixedUi {
	// for more, refer https://github.com/mitchellh/cli/blob/master/ui.go
	cliUi := &cli.PrefixedUi{
		AskPrefix: "Input:",
		AskSecretPrefix: "Input(secret):",
		OutputPrefix: "",
		InfoPrefix: "[INFO]",
		ErrorPrefix: "[ERROR]",
		WarnPrefix: "[WARNING]",
		Ui: &ui{
			&cli.BasicUi{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Reader:      os.Stdin,
			},
		},
	}

	return cliUi
}