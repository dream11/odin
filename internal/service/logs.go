package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
)

// Logs performs operation on logs like get logs
type Logs struct{}

// GetLogs Get logs
func (l *Logs) GetLogs(ctx *context.Context, request *logs.GetLogsRequest) (int64, error) {
	// Get logs
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return 0, err
	}
	client := logs.NewLogsServiceClient(conn)
	stream, err := client.GetLogs(*requestCtx, request)
	if err != nil {
		return 0, err
	}

	lastLogTime := int64(0)
	if request.GetStartTime() != 0 {
		lastLogTime = request.GetStartTime()
	}
	for {
		response, logStreamErr := stream.Recv()
		if logStreamErr != nil {
			if errors.Is(logStreamErr, context.Canceled) || logStreamErr == io.EOF {
				break
			} else {
				return lastLogTime, logStreamErr
			}
		}
		if response != nil {
			for _, logMessage := range response.Logs {
				fmt.Println(logMessage.GetMessage())
				lastLogTime = logMessage.GetTimestamp()
			}
		}
	}
	return lastLogTime, nil
}
