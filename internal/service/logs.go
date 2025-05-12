package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
)

var restrictedLogLevels = []string{"DEBUG"}

// Logs performs operation on logs like get logs
type Logs struct{}

// GetLogs retrieves and displays logs for a service
func (l *Logs) GetLogs(ctx *context.Context, request *logs.GetLogsRequest) (int64, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return request.GetStartTime(), err
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
		return request.GetStartTime(), err
	}

	lastLogTime := request.GetStartTime()

	for {
		response, err := stream.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			return lastLogTime, err
		}

		if response == nil {
			continue
		}

		verboseEnabled := false
		if (*ctx).Value(constant.VerboseEnabledKey) != nil {
			verboseEnabled = (*ctx).Value(constant.VerboseEnabledKey).(bool)
		}

		for _, logMessage := range response.Logs {
			if !util.Contains(logMessage.GetLevel(), restrictedLogLevels) ||
				(verboseEnabled == true && logMessage.GetLevel() == "DEBUG") {
				fmt.Println(logMessage.GetMessage())
			}
			if logMessage.GetTimestamp() > lastLogTime {
				lastLogTime = logMessage.GetTimestamp()
			}
		}
	}

	return lastLogTime, nil
}
