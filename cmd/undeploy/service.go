package undeploy

import (
	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/util"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string
var envName string

var serviceClient = service.Service{}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Undeploy service",
	Long:  `Undeploy service`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	serviceCmd.Flags().StringVar(&name, "name", "", "name of the service")
	serviceCmd.Flags().StringVar(&envName, "env", "", "name of the env")
	err := serviceCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatal("Error marking 'name' flag as required:", err)
	}
	undeployCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {
	envName = config.EnsureEnvPresent(envName)

	ctx := cmd.Context()
	traceId := util.GenerateTraceId()

	err := serviceClient.UndeployService(&ctx, &serviceProto.UndeployServiceRequest{
		EnvName:     envName,
		ServiceName: name,
	}, traceId)

	if err != nil {
		log.Fatal("Failed to undeploy service ", err)
	}
}
