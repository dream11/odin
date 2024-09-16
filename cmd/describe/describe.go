package describe

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe resources",
	Long:  `Describe resources`,
}

func init() {
	cmd.RootCmd.AddCommand(describeCmd)
}
