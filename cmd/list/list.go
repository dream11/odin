package list

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources",
	Long:  `List resources`,
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}
