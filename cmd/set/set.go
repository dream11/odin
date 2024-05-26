package set

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set config",
	Long:  `set default values in config file`,
}

func init() {
	cmd.RootCmd.AddCommand(setCmd)
}
