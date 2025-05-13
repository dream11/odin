package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd cobra root command
var RootCmd = &cobra.Command{
	Use:   "odin",
	Short: "Interface for service definitions & deployments into self-managed environments",
	Long:  `Deploy services in environments`,
}

func init() {
	RootCmd.PersistentFlags().StringP("profile", "p", "default", "odin profile")
	RootCmd.PersistentFlags().StringP("output", "o", "text", "odin output format")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "odin verbose logging")
	err := viper.BindPFlag("profile", RootCmd.PersistentFlags().Lookup("profile"))
	if err != nil {
		log.Fatal("Error while binding profile flag")
	}
	viper.SetDefault("profile", "default")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
