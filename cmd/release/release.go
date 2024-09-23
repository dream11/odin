package release

import (
	"github.com/dream11/odin/cmd"
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "release resources",
	Long:  `release resources`,
}

func init() {
	cmd.RootCmd.AddCommand(releaseCmd)
}
