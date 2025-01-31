package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

// Service performs operation on service like deploy. undeploy
type Service struct{}

var responseMap = make(map[string]string)

// DeployService deploys service
func (e *Service) DeployService(ctx *context.Context, request *serviceProto.DeployServiceRequest) (serviceProto.ServiceService_DeployServiceClient, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	return client.DeployService(*requestCtx, request)
}

func logFailedComponentMessagesOnce(response *serviceProto.ServiceResponse) {
	for _, compMessage := range response.ComponentsStatus {
		componentActionKey := compMessage.GetComponentName() + compMessage.GetComponentAction() + compMessage.GetComponentStatus()
		//code to not print the same message for component action again
		if responseMap[componentActionKey] == "" {
			if compMessage.GetComponentStatus() == "FAILED" {
				log.Error(fmt.Sprintf("Component %s %s %s %s", compMessage.GetComponentName(), compMessage.GetComponentAction(), compMessage.GetComponentStatus(), compMessage.GetError()))
			}
			responseMap[componentActionKey] = componentActionKey
		}
	}
}

func logFailedComponentMessagesOnceForComponents(response *serviceProto.ServiceResponse, components []string) {
	for _, compMessage := range response.ComponentsStatus {
		componentActionKey := compMessage.GetComponentName() + compMessage.GetComponentAction() + compMessage.GetComponentStatus()
		//code to not print the same message for component action again
		if responseMap[componentActionKey] == "" {
			if util.Contains(compMessage.ComponentName, components) {
				if compMessage.GetComponentStatus() == "FAILED" {
					log.Error(fmt.Sprintf("Component %s %s %s %s", compMessage.GetComponentName(), compMessage.GetComponentAction(), compMessage.GetComponentStatus(), compMessage.GetError()))
				}
				responseMap[componentActionKey] = componentActionKey
			}
		}
	}
}

// DeployServiceSet deploys service-set
func (e *Service) DeployServiceSet(ctx *context.Context, request *serviceProto.DeployServiceSetRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
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

		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
			return err
		}

		if response != nil {
			spinnerInstance.Stop()
			var buf bytes.Buffer
			table := tablewriter.NewWriter(&buf)
			table.SetHeader([]string{"Service Name", "Version", "Action", "Status", "Error"})
			for _, serviceResponse := range response.GetServices() {
				var errorMessage string
				if serviceResponse.ServiceResponse.ServiceStatus.ServiceStatus == "FAILED" {
					traceID := (*requestCtx).Value(constant.TraceIDKey)
					errorMessage += fmt.Sprintf("[%s] TraceID: %s \n", serviceResponse.ServiceResponse.ServiceStatus.Error, traceID)
				}
				row := []string{
					serviceResponse.ServiceIdentifier.ServiceName,
					serviceResponse.ServiceIdentifier.ServiceVersion,
					serviceResponse.ServiceResponse.ServiceStatus.ServiceAction,
					serviceResponse.ServiceResponse.ServiceStatus.ServiceStatus,
					errorMessage,
				}
				table.Append(row)
			}

			table.Render()
			message = buf.String()
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}
	fmt.Println(message)
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
			log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
			return err
		}

		if response != nil {
			message = response.GetServiceResponse().Message
			message += fmt.Sprintf("\n Service %s %s", response.GetServiceResponse().ServiceStatus.ServiceAction, response.GetServiceResponse().ServiceStatus.ServiceStatus)
			for _, compMessage := range response.GetServiceResponse().ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			logFailedComponentMessagesOnce(response.GetServiceResponse())
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}

	log.Info(message)
	return err
}

// UndeployService undeploy service
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
			log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
			return err
		}
		if response != nil {
			message = response.ServiceResponse.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
			for _, compMessage := range response.ServiceResponse.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			logFailedComponentMessagesOnce(response.GetServiceResponse())
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
			log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
			return err
		}
		if response != nil {
			message = response.ServiceResponse.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
			for _, compMessage := range response.ServiceResponse.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
			}
			logFailedComponentMessagesOnce(response.GetServiceResponse())
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

// ReleaseService :service operatioms
func (e *Service) ReleaseService(ctx *context.Context, request *serviceProto.ReleaseServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.ReleaseService(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Starting release service operation...")
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
			log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
			return err
		}
		if response != nil {
			message = response.Message
			message += fmt.Sprintf("\n Service %s %s", response.ServiceStatus.ServiceAction, response.ServiceStatus)
			for _, compMessage := range response.ComponentsStatus {
				message += fmt.Sprintf("\n Component %s %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus, compMessage.Error)
				if compMessage.ComponentStatus == "FAILED" {
					return errors.New(compMessage.Error)
				}
			}

			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}
	log.Info("Service released successfully !")
	return err
}

// ConvertToDeployServiceSetRequest converts service set to deploy service set request
func (e *Service) ConvertToDeployServiceSetRequest(serviceSet *serviceDto.ServiceSet, env string) serviceProto.DeployServiceSetRequest {
	var services []*serviceProto.ServiceIdentifier
	for _, service := range serviceSet.Services {
		services = append(services, &serviceProto.ServiceIdentifier{
			ServiceName:    service.Name,
			ServiceVersion: service.Version,
			Tags:           service.Tags,
			ForceFlag:      true,
		})
	}

	return serviceProto.DeployServiceSetRequest{
		EnvName:  env,
		Name:     serviceSet.Name,
		Services: services,
	}
}

// DescribeService describe service
func (e *Service) DescribeService(ctx *context.Context, request *serviceProto.DescribeServiceRequest) (*serviceProto.DescribeServiceResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	response, err := client.DescribeService(*requestCtx, request)
	if err != nil {
		log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
		return nil, err
	}

	return response, nil
}

// GetConflictingServices deploys service
func (e *Service) GetConflictingServices(ctx *context.Context, request *serviceProto.GetConflictingServicesRequest) (*serviceProto.GetConflictingServicesResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return &serviceProto.GetConflictingServicesResponse{}, err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	response, err := client.GetConflictingServices(*requestCtx, request)
	if err != nil {
		log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
	}
	return response, err
}
