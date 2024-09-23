package update

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "This command is accessed by using one of the subcommands env.",
	Long: `This command is accessed by using one of the subcommands below.

Subcommands:
    env          update an environment`,
}

func init() {
	cmd.RootCmd.AddCommand(updateCmd)
}
