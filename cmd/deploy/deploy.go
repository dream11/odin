package deploy

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy resources",
	Long:  `Deploy resources`,
}

func init() {
	cmd.RootCmd.AddCommand(deployCmd)
}
