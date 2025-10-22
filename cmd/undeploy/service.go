package undeploy

import (
	"context"
	"fmt"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string
var envName string

var serviceClient = service.Service{}

var serviceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Undeploy service",
	Long:  `Undeploy service`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]
		execute(cmd)
	},
}

func init() {
	serviceCmd.Flags().StringVar(&envName, "env", "", "name of the env")
	undeployCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {
	if envName == "prod" {
		log.Infof("Undeploying service %s in production environment enter PROD to confirm", name)
		consentMessage := fmt.Sprintf(constant.ConsentMessageTemplate, "PROD")
		util.AskForConfirmation("PROD", consentMessage)
	}
	envName = config.EnsureEnvPresent(envName)

	ctx := cmd.Context()
	verboseEnabled, err := cmd.Flags().GetBool(constant.VerboseFlag)
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.WithValue(ctx, constant.VerboseEnabledKey, verboseEnabled)

	err = serviceClient.UndeployService(&ctx, &serviceProto.UndeployServiceRequest{
		EnvName:     envName,
		ServiceName: name,
	})

	if err != nil {
		util.LogGrpcError(err, "Failed to undeploy service: ")
	}
}
