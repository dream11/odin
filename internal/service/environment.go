package service

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/briandowns/spinner"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
	log "github.com/sirupsen/logrus"
)

// Environment performs operation on environment like create, list, describe, delete
type Environment struct{}

// SpinnerColor Defines color of spinner
var SpinnerColor = "fgHiBlue"

// SpinnerStyle Defines style of spinner
var SpinnerStyle = "bold"

// SpinnerType Defines type of spinner
var SpinnerType = 14

// SpinnerDelay Defines spinner delay
var SpinnerDelay = 100 * time.Millisecond

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
	spinner := spinner.New(spinner.CharSets[SpinnerType], SpinnerDelay)
	spinner.Color(SpinnerColor, SpinnerStyle)

	for {
		response, err := stream.Recv()
		spinner.Stop()
		if err != nil {
			if err == context.Canceled || err == io.EOF {
				break
			}
			return err
		}
		spinner.Prefix = fmt.Sprintf(" %s  ", response.Message)
		spinner.Start()
	}
	log.Info("Environment created successfully")

	return err
}
