package delete

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources",
	Long:  `Delete resources`,
}

func init() {
	cmd.RootCmd.AddCommand(deleteCmd)
}
