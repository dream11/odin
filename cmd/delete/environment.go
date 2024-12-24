package delete

import (
	"github.com/dream11/odin/internal/service"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string

var environmentClient = service.Environment{}

var environmentCmd = &cobra.Command{
	Use:   "env",
	Short: "Delete environment",
	Long:  `Delete environment`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	environmentCmd.Flags().StringVar(&name, "name", "", "name of the env")
	err := environmentCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	deleteCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	err := environmentClient.DeleteEnvironment(&ctx, &environment.DeleteEnvironmentRequest{
		EnvName: name,
	})

	if err != nil {
		log.Fatal("Failed to delete environment ", err)
	}
}
