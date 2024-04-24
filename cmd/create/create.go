package create

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources",
	Long:  `Create resources`,
}

func init() {
	cmd.RootCmd.AddCommand(createCmd)
}
