package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	service "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
)

// Service performs operation on service like deploy. undeploy
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

	log.Info("Deploying Service...")
	spinner := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	err = spinner.Color(constant.SpinnerColor, constant.SpinnerStyle)
	if err != nil {
		return err
	}

	var message string
	for {
		response, err := stream.Recv()
		spinner.Stop()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			return err
		}

		if response != nil {
			message = response.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceStatus.ServiceAction, response.ServiceStatus.ServiceStatus)
			for _, compMessage := range response.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			spinner.Prefix = fmt.Sprintf(" %s  ", message)
			spinner.Start()
		}
	}

	log.Info(message)
	return err
}

// UndeployService undeploys service
func (e *Service) UndeployService(ctx *context.Context, request *service.UndeployServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}

	client := service.NewServiceServiceClient(conn)
	stream, err := client.UndeployService(*requestCtx, request)

	if err != nil {
		return err
	}

	log.Info("Undeploying Service...")
	spinner := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	err = spinner.Color(constant.SpinnerColor, constant.SpinnerStyle)
	if err != nil {
		return err
	}
	var message string
	for {
		response, err := stream.Recv()
		spinner.Stop()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
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
