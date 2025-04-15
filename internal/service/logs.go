package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"strings"

	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
	"google.golang.org/grpc"
)

// Logs performs operation on logs like get logs
type Logs struct{}

// GetLogs retrieves and displays logs for a service
func (l *Logs) GetLogs(ctx *context.Context, request *logs.GetLogsRequest) (int64, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return request.GetStartTime(), err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error closing connection: %v\n", err)
		}
	}(conn)

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

		for _, logMessage := range response.Logs {
			if !strings.Contains(logMessage.GetMessage(), "DEBUG") {
				fmt.Println(logMessage.GetMessage())
			}
			lastLogTime = logMessage.GetTimestamp()
		}
	}

	return lastLogTime, nil
}
