package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of odin",
	Long:  `All software has versions. This is odin's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2.0.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
