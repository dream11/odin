package delete

import (
	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/util"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	"github.com/spf13/cobra"
)

var name string

var environmentClient = service.Environment{}

var environmentCmd = &cobra.Command{
	Use:   "env <name>",
	Short: "Delete environment",
	Long:  `Delete environment`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]
		execute(cmd)
	},
}

func init() {
	deleteCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	err := environmentClient.DeleteEnvironment(&ctx, &environment.DeleteEnvironmentRequest{
		EnvName: name,
	})

	if err != nil {
		util.LogGrpcError(err, "Failed to delete environment:")
	}
}
