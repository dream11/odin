package cmd

import (
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "odin",
	Short: "Interface for service definitions & deployments into self-managed environments",
	Long:  `Deploy services in environments`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("profile", "p", "default", "odin profile")
	viper.BindPFlag("profile", RootCmd.PersistentFlags().Lookup("profile"))
	viper.SetDefault("profile", "default")
}