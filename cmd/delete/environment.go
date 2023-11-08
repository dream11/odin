package delete

import (
	"encoding/json"
	"fmt"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var name string

var environmentClient = service.Environment{}

var environmentCmd = &cobra.Command{
	Use:   "environment",
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
	deleteCmd.AddCommand(environmentCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()
	response, err := environmentClient.DeleteEnvironment(&ctx, &environment.DeleteEnvironmentRequest{
		EnvName: name,
	})

	if err != nil {
		log.Fatal("Failed to delete environment ", err)
	}

	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(response, outputFormat)
}

func writeOutput(response *environment.DeleteEnvironmentResponse, format string) {

	switch format {
	case constant.TEXT:
		fmt.Print(response.Message)
	case constant.JSON:
		output, _ := json.Marshal(response)
		fmt.Print(string(output))
	default:
		log.Fatal("Unknown output format: ", format)
	}
}
