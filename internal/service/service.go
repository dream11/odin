package service

import (
	"context"
	"fmt"
	service "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
)

// Service performs operation on service like deploy, undeploy, describe
type Service struct{}

// DeployService deploys service
func (e *Service) DeployService(ctx *context.Context, request *service.DeployServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := service.NewServiceServiceClient(conn)
	stream, err := client.DeployService(*requestCtx, request)
	if err != nil {
		return err
	}

	// Receive and process streaming responses
	for {
		response, err := stream.Recv()
		if err != nil {
			// Check for the end of the stream
			if err == context.Canceled {
				break
			}

			log.Fatalf("Error receiving stream: %v", err)
		}

		// Process the received stream
		fmt.Println("Received DeployServiceResponse:", response)
	}

	return err
}
