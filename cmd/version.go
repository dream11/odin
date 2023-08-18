package cmd

import (
	"fmt"

	"github.com/dream11/odin/app"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of odin",
	Long:  `All software has versions. This is odin's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.App.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
