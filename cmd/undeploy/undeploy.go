package undeploy

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

var undeployCmd = &cobra.Command{
	Use:   "undeploy",
	Short: "Undeploy resources",
	Long:  `undeploy resources`,
}

func init() {
	cmd.RootCmd.AddCommand(undeployCmd)
}
