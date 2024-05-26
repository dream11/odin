package set

import (
	"github.com/dream11/odin/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "set profile",
	Long:  `modify profile in config file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.SetProfile(args[0])
		log.Info("profile set to [", args[0], "] successfully")
	},
}

func init() {
	setCmd.AddCommand(profileCmd)
}
