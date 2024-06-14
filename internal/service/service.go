package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
)

// Service performs operation on service like deploy. undeploy
type Service struct{}

// DeployService deploys service
func (e *Service) DeployService(ctx *context.Context, request *serviceProto.DeployServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.DeployService(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Deploying Service...")
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
			message = response.ServiceResponse.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
			for _, compMessage := range response.ServiceResponse.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}

	log.Info(message)
	return err
}

// DeployServiceSet deploys service-set
func (e *Service) DeployServiceSet(ctx *context.Context, request *serviceProto.DeployServiceSetRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.DeployServiceSet(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Deploying Service Set..")
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

			for _, serviceRespose := range response.GetServices() {
				message = fmt.Sprintf("\n Service %s %s %s", serviceRespose.ServiceIdentifier, serviceRespose.ServiceResponse.ServiceStatus, serviceRespose.ServiceResponse.Message)
				for _, compMessage := range serviceRespose.ServiceResponse.ComponentsStatus {
					message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
				}
			}
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}

	log.Info(message)
	return err
}

// DeployReleasedService deploys service
func (e *Service) DeployReleasedService(ctx *context.Context, request *serviceProto.DeployReleasedServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.DeployReleasedService(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Deploying Service...")
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
			message = response.GetServiceResponse().Message
			message += fmt.Sprintf("\n Service %s %s", response.GetServiceResponse().ServiceStatus.ServiceAction, response.GetServiceResponse().ServiceStatus.ServiceStatus)
			for _, compMessage := range response.GetServiceResponse().ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}

	log.Info(message)
	return err
}

// UndeployService undeploys service
func (e *Service) UndeployService(ctx *context.Context, request *serviceProto.UndeployServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}

	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.UndeployService(*requestCtx, request)

	if err != nil {
		return err
	}

	log.Info("Undeploying Service...")
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
			message = response.ServiceResponse.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
			for _, compMessage := range response.ServiceResponse.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}
	log.Info(message)
	return err
}

// OperateService :service operatioms
func (e *Service) OperateService(ctx *context.Context, request *serviceProto.OperateServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.OperateService(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Starting service operation...")
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
			message = response.ServiceResponse.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
			for _, compMessage := range response.ServiceResponse.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}
	log.Info(message)
	return err
}

// ListService deploys service
func (e *Service) ListService(ctx *context.Context, request *serviceProto.ListServiceRequest) (*serviceProto.ListServiceResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return &serviceProto.ListServiceResponse{}, err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	response, err := client.ListService(*requestCtx, request)
	return response, err
}

// ConvertToDeployServiceSetRequest converts service set to deploy service set request

func (e *Service) ConvertToDeployServiceSetRequest(serviceSet *serviceDto.ServiceSet, env string) serviceProto.DeployServiceSetRequest {
	var services []*serviceProto.ServiceIdentifier
	for _, service := range serviceSet.Services {
		services = append(services, &serviceProto.ServiceIdentifier{
			ServiceName:    service.Name,
			ServiceVersion: service.Version,
		})
	}

	return serviceProto.DeployServiceSetRequest{
		EnvName:  env,
		Name:     serviceSet.Name,
		Services: services,
	}
}
