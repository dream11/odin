package operate

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

// operateCmd represents the operate command
var operateCmd = &cobra.Command{
	Use:   "operate",
	Short: "This command is accessed by using one of the subcommands: [env, service, component]",
	Long: `This command is accessed by using one of the subcommands below:

Subcommands:
    component    Operate on a component
    env          Operate an environment
    service      Perform operations on a service`,
}

func init() {
	cmd.RootCmd.AddCommand(operateCmd)
}
