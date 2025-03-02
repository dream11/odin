package status

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "This command is accessed by using one of the subcommands: [env]",
	Long: `This command is accessed by using one of the subcommands below:

Subcommands:
    env          Fetch deployment status of the environment`,
}

func init() {
	cmd.RootCmd.AddCommand(statusCmd)
}
