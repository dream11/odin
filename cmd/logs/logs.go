package logs

import (
	"context"
	"fmt"
	"github.com/dream11/odin/cmd"
	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/constant"
	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
	"github.com/spf13/cobra"
)

var traceId string
var logsClient = service.Logs{}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch logs",
	Long:  "Fetch logs",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	logsCmd.Flags().StringVar(&traceId, "traceid", "", "Trace Id to fetch logs")
	cmd.RootCmd.AddCommand(logsCmd)
}

func execute(cmd *cobra.Command) {
	ctx := cmd.Context()

	contextWithTrace := context.WithValue(ctx, constant.TraceIDKey, traceId)

	streamCtx, cancelFunction := context.WithCancel(context.Background())
	defer cancelFunction()

	var searchAfterParams []int64
	var err error
	fmt.Printf("Fetching logs for traceId: %s\n", traceId)
	for {
		select {
		case <-streamCtx.Done():
			return
		default:
			// Get logs with retry on error
			follow := true
			searchAfterParams, err = logsClient.GetLogs(&contextWithTrace, &logs.GetLogsRequest{
				TraceId:           traceId,
				Follow:            &follow,
				SearchAfterParams: searchAfterParams,
			})
			if err != nil {
				continue
			}
		}
	}

}
