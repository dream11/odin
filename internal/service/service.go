package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/retryable"
	"github.com/dream11/odin/pkg/util"
	logs "github.com/dream11/odin/proto/gen/go/dream11/od/logs/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
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
	log.Infof(constant.ServiceExecutionMessageTemplate, "Deploying", request.GetServiceDefinition().GetName(), request.GetEnvName())

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

// UndeployService undeploy service
func (e *Service) UndeployService(ctx *context.Context, request *serviceProto.UndeployServiceRequest) error {
	log.Infof(constant.ServiceExecutionMessageTemplate, "Undeploying", request.GetServiceName(), request.GetEnvName())
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

	getMessage := func(response *serviceProto.UndeployServiceResponse) string {
		return util.GenerateResponseMessage(response.GetServiceResponse())
	}
	getStatus := func(response *serviceProto.UndeployServiceResponse) (string, string) {
		return response.GetServiceResponse().GetServiceStatus().GetServiceStatus(),
			response.GetServiceResponse().GetServiceStatus().GetServiceAction()
	}

	return handleResponse(stream, cancelFunction, getMessage, getStatus)
}

// OperateService :service operations
func (e *Service) OperateService(ctx *context.Context, request *serviceProto.OperateServiceRequest) error {
	log.Infof(constant.ServiceExecutionMessageTemplate, "Operating", request.GetServiceName(), request.GetEnvName())

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

// streamLogs streams logs for a service
func streamLogs(streamCtx context.Context, ctx *context.Context, serviceName string) {
	var err error
	var searchAfterParams []int64
	traceID := (*ctx).Value(constant.TraceIDKey).(string)
	follow := true
	fmt.Printf("Fetching live logs for service: %s \n", serviceName)
	for {
		select {
		case <-streamCtx.Done():
			return
		default:
			// Get logs with retry on error
			searchAfterParams, err = logsClient.GetLogs(ctx, &logs.GetLogsRequest{
				TraceId:           traceID,
				Follow:            &follow,
				ServiceName:       &serviceName,
				SearchAfterParams: searchAfterParams,
			})
			if err != nil {
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
			// Wait for few seconds to ensure all logs are received
			log.Info(getMessage(response))
			log.Info(constant.CheckingAdditionalLogsMessage)
			time.Sleep(30 * time.Second)
			cancelFunc()
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
