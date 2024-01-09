package service

import (
	"context"

	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
)

// Environment performs operation on environment like create, list, describe, delete
type Environment struct{}

// ListEnvironments List environments
func (e *Environment) ListEnvironments(ctx *context.Context, request *environment.ListEnvironmentRequest) (*environment.ListEnvironmentResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := environment.NewEnvironmentServiceClient(conn)
	response, err := client.ListEnvironment(*requestCtx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateEnvironment creates environment
func (e *Environment) CreateEnvironment(ctx *context.Context, request *environment.CreateEnvironmentRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := environment.NewEnvironmentServiceClient(conn)
	stream, err := client.CreateEnvironment(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Creating environment...")
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == context.Canceled {
				break
			}
			return err
		}

		log.Info(response)
	}

	return err
}
