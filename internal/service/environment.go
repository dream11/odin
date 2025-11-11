package service

import (
	"context"
	"errors"
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

	log.Info("\nCreating environment...\n")
	spinnerInstance := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	err = spinnerInstance.Color(constant.SpinnerColor, constant.SpinnerStyle)
	if err != nil {
		return err
	}
	var message string
	for {
		response, err := stream.Recv()
		spinnerInstance.Stop()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			return err
		}
		if response != nil {
			message = response.Message
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", response.Message)
			spinnerInstance.Start()
		}
	}
	log.Info(message)
	return err
}

// DeleteEnvironment deletes environment
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

	log.Info("\nDeleting environment...\n")
	spinnerInstance := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	err = spinnerInstance.Color(constant.SpinnerColor, constant.SpinnerStyle)
	if err != nil {
		return err
	}
	var message string
	for {
		response, err := stream.Recv()
		spinnerInstance.Stop()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			return err
		}
		if response != nil {
			message = response.Message
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", response.Message)
			spinnerInstance.Start()
		}
	}
	log.Info(message)
	return err
}

// DescribeEnvironment shows environment details including services and resources in it
func (e *Environment) DescribeEnvironment(ctx *context.Context, request *environment.DescribeEnvironmentRequest) (*environment.DescribeEnvironmentResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}

	client := environment.NewEnvironmentServiceClient(conn)
	response, err := client.DescribeEnvironment(*requestCtx, request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// EnvironmentStatus shows environment status including services and components in it
func (e *Environment) EnvironmentStatus(ctx *context.Context, request *environment.StatusEnvironmentRequest) (*environment.StatusEnvironmentResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}

	client := environment.NewEnvironmentServiceClient(conn)
	log.Info("Getting environment status...")
	spinnerInstance := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	_ = spinnerInstance.Color(constant.SpinnerColor, constant.SpinnerStyle)
	spinnerInstance.Start()
	stream, err := client.StatusEnvironment(*requestCtx, request)
	if err != nil {
		return nil, err
	}
	var response *environment.StatusEnvironmentResponse
	var prevResponse *environment.StatusEnvironmentResponse

	for {
		prevResponse = response
		response, err = stream.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			return nil, err
		}
	}
	spinnerInstance.Stop()
	return prevResponse, nil
}
