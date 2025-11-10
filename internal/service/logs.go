package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var restrictedLogLevels = []string{"DEBUG", "WARN"}

// Logs performs operation on logs like get logs
type Logs struct{}

// GetLogs retrieves and displays logs for a service
func (l *Logs) GetLogs(ctx *context.Context, request *logs.GetLogsRequest) ([]int64, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return request.GetSearchAfterParams(), err
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error closing connection: %v\n", err)
		}
	}()

	client := logs.NewLogsServiceClient(conn)
	stream, err := client.GetLogs(*requestCtx, request)
	if err != nil {
		return request.GetSearchAfterParams(), err
	}

	searchAfterParams := request.GetSearchAfterParams()

	for {
		response, err := stream.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
				time.Sleep(5 * time.Second)
				continue
			}
			return searchAfterParams, err
		}

		if response == nil {
			continue
		}

		verboseEnabled := false
		if (*ctx).Value(constant.VerboseEnabledKey) != nil {
			verboseEnabled = (*ctx).Value(constant.VerboseEnabledKey).(bool)
			if verboseEnabled {
				restrictedLogLevels = []string{}
			}
		}

		for _, logMessage := range response.Logs {
			if !util.Contains(logMessage.GetLevel(), restrictedLogLevels) {
				fmt.Println(logMessage.GetMessage())
				searchAfterParams = logMessage.GetSearchAfterParams()
			}
		}
	}

	return searchAfterParams, nil
}
