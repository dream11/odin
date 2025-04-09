package service

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	component "github.com/dream11/odin/proto/gen/go/dream11/od/component/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
)

// Component performs operation on component like operate
type Component struct{}

// OperateComponent operate Component
func (e *Component) OperateComponent(ctx *context.Context, request *serviceProto.OperateServiceRequest) error {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	stream, err := client.OperateService(*requestCtx, request)
	if err != nil {
		return err
	}

	log.Info("Starting component operation...")
	spinnerInstance := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	err = spinnerInstance.Color(constant.SpinnerColor, constant.SpinnerStyle)
	if err != nil {
		return err
	}

	var message string
	var retries = 0

	responseChan := make(chan *serviceProto.OperateServiceResponse)
	errorChan := make(chan error)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				response, err := stream.Recv()
				if err != nil {
					errorChan <- err
				}
				responseChan <- response
			}
		}
	}(*requestCtx)

	for {
		recvCtx, cancel := context.WithTimeout(*requestCtx, constant.Timeout)

		select {
		case <-recvCtx.Done():
			spinnerInstance.Stop()
			cancel()
			if !util.CanPerformRetry(retries, constant.MaxRetries) {
				return nil
			}
			stream, err = reconnectOperateStream(client, requestCtx, request, stream)
			if err != nil {
				return nil
			}
			retries++
		case err := <-errorChan:
			spinnerInstance.Stop()
			cancel()
			if !util.IsRetryable(err) || !util.CanPerformRetry(retries, constant.MaxRetries) {
				if err != io.EOF {
					return err
				} else if err == io.EOF {
					log.Info(message)
				}
				return nil
			}
			stream, err = reconnectOperateStream(client, requestCtx, request, stream)
			if err != nil {
				return nil
			}
			retries++
		case response := <-responseChan:
			spinnerInstance.Stop()
			cancel()
			if response != nil {
				message = util.GenerateResponseMessageComponentSpecific(response.GetServiceResponse(), []string{request.GetComponentName()})
				logFailedComponentMessagesOnceForComponents(response.GetServiceResponse(), []string{request.GetComponentName()})
				spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
				spinnerInstance.Start()
				retries = 0
			}
		}
	}
}

func reconnectOperateStream(client serviceProto.ServiceServiceClient, requestCtx *context.Context, request *serviceProto.OperateServiceRequest, stream serviceProto.ServiceService_OperateServiceClient) (serviceProto.ServiceService_OperateServiceClient, error) {
	if err := stream.CloseSend(); err != nil {
		return nil, err
	}

	newStream, err := client.OperateService(*requestCtx, request)
	if err != nil {
		return nil, err
	}

	return newStream, nil

}

// ListComponentType List component types
func (e *Component) ListComponentType(ctx *context.Context, request *component.ListComponentTypeRequest) (*component.ListComponentTypeResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := component.NewComponentServiceClient(conn)
	response, err := client.ListComponentType(*requestCtx, request)
	if err != nil {
		log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
		return nil, err
	}

	return response, nil
}

// DescribeComponentType List component types
func (e *Component) DescribeComponentType(ctx *context.Context, request *component.DescribeComponentTypeRequest) (*component.DescribeComponentTypeResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := component.NewComponentServiceClient(conn)
	response, err := client.DescribeComponentType(*requestCtx, request)
	if err != nil {
		log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
		return nil, err
	}

	return response, nil
}

// CompareOperationChanges compares the operation changes
func (e *Component) CompareOperationChanges(ctx *context.Context, request *serviceProto.OperateComponentDiffRequest) (*serviceProto.OperateComponentDiffResponse, error) {

	for retries := 0; retries < constant.MaxRetries; retries++ {
		ctxWithTimeout, cancel := context.WithTimeout(*ctx, constant.Timeout)
		defer cancel()

		conn, requestCtx, err := grpcClient(&ctxWithTimeout)
		if err != nil {
			return nil, err
		}

		client := serviceProto.NewServiceServiceClient(conn)
		response, err := client.OperateComponentDiff(*requestCtx, request)
		if err == nil {
			return response, nil
		}

		if !util.IsRetryable(err) {
			return nil, err
		}
		time.Sleep(constant.Timeout)
		if retries == 0 {
			log.Warnf(constant.InitiatingRetryMessage)
		}
		log.Infof(constant.RetryingMessage, retries+1, constant.MaxRetries)
	}

	log.Fatalf(constant.MaxRetriesReached)
	return nil, nil
}
