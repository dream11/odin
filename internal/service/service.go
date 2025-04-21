package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/retryable"
	"github.com/dream11/odin/pkg/util"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service performs operation on service like deploy. undeploy
type Service struct{}

var logsClient = Logs{}

var serviceTerminalConditions = map[string][]string{
	"DEPLOY":   {"SUCCESSFUL", "FAILED"},
	"UNDEPLOY": {"SUCCESSFUL", "FAILED"},
	"OPERATE":  {"SUCCESSFUL", "FAILED"},
	"VALIDATE": {"FAILED"},
}

// RetryableStatusCodes are the status codes that are retryable
var RetryableStatusCodes = []codes.Code{codes.DeadlineExceeded, codes.Canceled, codes.Unavailable}

// StreamReceiverInterface defines the Recv method that the stream must implement
type StreamReceiverInterface[R any] interface {
	Recv() (R, error)
}

type getStatus[R any] func(response R) (serviceAction, serviceStatus string)

type getMessage[R any] func(response R) string

// DeployService deploys service
func (e *Service) DeployService(ctx *context.Context, request *serviceProto.DeployServiceRequest) error {
	log.Info("Deploying Service...")

	// Create a context with cancelFunction for the entire operation
	streamCtx, cancelFunction := context.WithCancel(context.Background())
	defer cancelFunction()

	// Start log streaming in background
	go streamLogs(streamCtx, ctx, request.GetServiceDefinition().GetName())

	// Attempt deployment with retries
	return retry.Do(
		func() error {
			conn, requestCtx, err := grpcClient(ctx)
			if err != nil {
				return err
			}
			defer func() {
				err := conn.Close()
				if err != nil {
					log.Errorf("Error closing connection: %v\n", err)
				}
			}()

			client := serviceProto.NewServiceServiceClient(conn)
			stream, err := client.DeployService(*requestCtx, request)
			if err != nil {
				return err
			}
			getMessage := func(response *serviceProto.DeployServiceResponse) string {
				return util.GenerateResponseMessage(response.GetServiceResponse())
			}
			getStatus := func(response *serviceProto.DeployServiceResponse) (string, string) {
				return response.GetServiceResponse().GetServiceStatus().GetServiceStatus(),
					response.GetServiceResponse().GetServiceStatus().GetServiceAction()
			}

			return handleResponse(stream, cancelFunction, getMessage, getStatus)
		},
		retry.Delay(constant.Delay),
		retry.RetryIf(isRetryableError),
	)
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

		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
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
	log.Info("Deploying Service...")
	// Create a context with cancelFunction for the entire operation
	streamCtx, cancelFunction := context.WithCancel(context.Background())
	defer cancelFunction()

	// Start log streaming in background
	go streamLogs(streamCtx, ctx, request.GetServiceIdentifier().GetServiceName())

	// Attempt deployment with retries
	return retry.Do(
		func() error {
			conn, requestCtx, err := grpcClient(ctx)
			if err != nil {
				return err
			}
			defer func() {
				err := conn.Close()
				if err != nil {
					log.Errorf("Error closing connection: %v\n", err)
				}
			}()

			client := serviceProto.NewServiceServiceClient(conn)
			stream, err := client.DeployReleasedService(*requestCtx, request)
			if err != nil {
				return err
			}

			getMessage := func(response *serviceProto.DeployReleasedServiceResponse) string {
				return util.GenerateResponseMessage(response.GetServiceResponse())
			}
			getStatus := func(response *serviceProto.DeployReleasedServiceResponse) (string, string) {
				return response.GetServiceResponse().GetServiceStatus().GetServiceStatus(),
					response.GetServiceResponse().GetServiceStatus().GetServiceAction()
			}

			return handleResponse(stream, cancelFunction, getMessage, getStatus)
		},
		retry.Delay(constant.Delay),
		retry.RetryIf(isRetryableError),
	)
}

// UndeployService undeploy service
func (e *Service) UndeployService(ctx *context.Context, request *serviceProto.UndeployServiceRequest) error {
	log.Info("Undeploying Service...")
	traceID := util.GenerateTraceID()
	contextWithTrace := context.WithValue(*ctx, constant.TraceIDKey, traceID)

	// Create a context with cancelFunction for the entire operation
	streamCtx, cancelFunction := context.WithCancel(context.Background())
	defer cancelFunction()

	// Start log streaming in background
	go streamLogs(streamCtx, &contextWithTrace, request.GetServiceName())

	conn, requestCtx, err := grpcClient(&contextWithTrace)
	if err != nil {
		return err
	}

	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.UndeployService(*requestCtx, request)

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
			return err
		}
		if response != nil {
			if isActionCompleted(response.GetServiceResponse().GetServiceStatus().GetServiceAction(), response.GetServiceResponse().GetServiceStatus().GetServiceStatus()) {
				cancelFunction()
				message = response.GetServiceResponse().GetMessage()
				message += fmt.Sprintf("\n Service %s %s", response.ServiceResponse.ServiceStatus.ServiceAction, response.ServiceResponse.ServiceStatus)
				for _, compMessage := range response.ServiceResponse.ComponentsStatus {
					message += fmt.Sprintf("\n Component %s %s %s", compMessage.ComponentName, compMessage.ComponentAction, compMessage.ComponentStatus)
				}
				for _, compMessage := range response.GetServiceResponse().GetComponentsStatus() {
					if compMessage.GetComponentStatus() == "FAILED" {
						message += fmt.Sprintf("Component %s %s %s %s", compMessage.GetComponentName(), compMessage.GetComponentAction(), compMessage.GetComponentStatus(), compMessage.GetError())
					}
				}
			}
		}
	}
	log.Info(message)
	return err
}

// OperateService :service operations
func (e *Service) OperateService(ctx *context.Context, request *serviceProto.OperateServiceRequest) error {
	log.Info("Starting service operation...")

	// Create a context with cancelFunction for the entire operation
	streamCtx, cancelFunction := context.WithCancel(context.Background())
	defer cancelFunction()

	// Start log streaming in background
	go streamLogs(streamCtx, ctx, request.GetServiceName())

	// Attempt operation with retries
	return retry.Do(
		func() error {
			conn, requestCtx, err := grpcClient(ctx)
			if err != nil {
				return err
			}
			defer func() {
				err := conn.Close()
				if err != nil {
					log.Errorf("Error closing connection: %v\n", err)
				}
			}()

			client := serviceProto.NewServiceServiceClient(conn)
			stream, err := client.OperateService(*requestCtx, request)
			if err != nil {
				return err
			}
			getMessage := func(response *serviceProto.OperateServiceResponse) string {
				return util.GenerateResponseMessage(response.GetServiceResponse())
			}
			getStatus := func(response *serviceProto.OperateServiceResponse) (string, string) {
				return response.GetServiceResponse().GetServiceStatus().GetServiceStatus(),
					response.GetServiceResponse().GetServiceStatus().GetServiceAction()
			}

			return handleResponse(stream, cancelFunction, getMessage, getStatus)
		},
		retry.Delay(constant.Delay),
		retry.RetryIf(isRetryableError),
	)
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

// ReleaseService :service operations
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
			Labels:         service.Labels,
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

// streamLogs streams logs for a service
func streamLogs(streamCtx context.Context, ctx *context.Context, serviceName string) {
	var err error
	lastLogTime := int64(0)
	traceID := (*ctx).Value(constant.TraceIDKey).(string)
	follow := true
	// Start the spinner in a background goroutine
	go func() {
		spinnerInstance := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
		err = spinnerInstance.Color(constant.SpinnerColor, constant.SpinnerStyle)
		if err != nil {
			spinnerInstance.Stop()
		}
		spinnerInstance.Prefix = fmt.Sprintf("Fetching live logs for service: %s ", serviceName)
		spinnerInstance.Suffix = "\n"
		spinnerInstance.Start()
		time.Sleep(30 * time.Second)
		spinnerInstance.Stop()
	}()

	for {
		select {
		case <-streamCtx.Done():
			return
		default:
			// Get logs with retry on error
			lastLogTime, err = logsClient.GetLogs(ctx, &logs.GetLogsRequest{
				TraceId:     &traceID,
				Follow:      &follow,
				ServiceName: &serviceName,
				StartTime:   &lastLogTime,
			})

			if err != nil {
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}
}

// handleResponse streams the service deploy response and call cancel on action termination
func handleResponse[S StreamReceiverInterface[R], R any](stream S, cancelFunc context.CancelFunc, getMessage getMessage[R], getStatus getStatus[R]) error {
	var serviceAction, serviceStatus string
	for {
		response, err := stream.Recv()
		if err != nil {
			if isActionCompleted(serviceAction, serviceStatus) {
				cancelFunc()
				return nil
			}

			st, _ := status.FromError(err)
			if err == io.EOF || slices.Contains(RetryableStatusCodes, st.Code()) ||
				(strings.Contains(err.Error(), "RST_STREAM") && st.Code() == codes.Internal) {
				return retryable.NewRetryableError(err, true)
			}

			cancelFunc()
			return err
		}
		serviceStatus, serviceAction = getStatus(response)
		if isActionCompleted(serviceAction, serviceStatus) {
			cancelFunc()
			log.Info(getMessage(response))
			return nil
		}
	}
}

// isActionCompleted checks if the action is completed based on the service action and status
func isActionCompleted(serviceAction, status string) bool {
	if serviceAction == "" || status == "" {
		return false
	}
	return slices.Contains(serviceTerminalConditions[serviceAction], status)
}

// isRetryableError checks if the error is retryable
func isRetryableError(err error) bool {
	var re retryable.Error
	if errors.As(err, &re) && re.Retryable() {
		log.Info("Connection lost, retrying...")
		return true
	}
	return false
}
