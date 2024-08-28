package set

import (
	"github.com/dream11/odin/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var setEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "odin set default environment",
	Long:  `modify environment in config file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Error: need one more parameter for environment name")
		}
		setEnvironment(args[0])
	},
}

func init() {
	setCmd.AddCommand(setEnvCmd)
}

func setEnvironment(envName string) {
	config.UpdateEnvName(envName)
}
