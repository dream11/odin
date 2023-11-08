package service

import (
	"context"
	"fmt"
	"io"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
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
	spinner := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	spinner.Color(constant.SpinnerColor, constant.SpinnerStyle)

	var message string
	for {
		response, err := stream.Recv()
		spinner.Stop()
		if err != nil {
			if err == context.Canceled || err == io.EOF {
				break
			}
			return err
		}
		if response != nil {
			message = response.Message
			spinner.Prefix = fmt.Sprintf(" %s  ", response.Message)
			spinner.Start()
		}
	}
	log.Info(message)
	return err
}

// DeleteEnvironment : Delete environment
func (e *Environment) DeleteEnvironment(ctx *context.Context, request *environment.DeleteEnvironmentRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}

	client := environment.NewEnvironmentServiceClient(conn)
	stream, err := client.DeleteEnvironment(*requestCtx, request)

	if err != nil {
		return err
	}

	log.Info("Deleting environment...")
	spinner := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	spinner.Color(constant.SpinnerColor, constant.SpinnerStyle)

	var message string
	for {
		response, err := stream.Recv()
		spinner.Stop()
		if err != nil {
			if err == context.Canceled || err == io.EOF {
				break
			}
			return err
		}
		if response != nil {
			message = response.Message
			spinner.Prefix = fmt.Sprintf(" %s  ", response.Message)
			spinner.Start()
		}
	}
	log.Info(message)
	return err
}
