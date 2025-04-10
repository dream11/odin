package service

import (
	"context"
	"fmt"
	"time"

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
	reconnect := func() (serviceProto.ServiceService_OperateServiceClient, error) {
		return reconnectOperateStream(client, requestCtx, request, stream)
	}

	generateResponse := func(response *serviceProto.OperateServiceResponse) string {
		message := util.GenerateResponseMessageComponentSpecific(response.GetServiceResponse(), []string{request.GetComponentName()})
		logFailedComponentMessagesOnceForComponents(response.GetServiceResponse(), []string{request.GetComponentName()})
		return message
	}

	return handleStreamResponse(stream, reconnect, generateResponse)
}

func reconnectOperateStream(client serviceProto.ServiceServiceClient, requestCtx *context.Context, request *serviceProto.OperateServiceRequest, stream serviceProto.ServiceService_OperateServiceClient) (serviceProto.ServiceService_OperateServiceClient, error) {
	if err := stream.CloseSend(); err != nil {
		return nil, err
	}

	newStream, err := client.OperateService(*requestCtx, request)
	if err == nil {
		return newStream, nil
	}

	return nil, fmt.Errorf(constant.FailedRetryMessage)

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

		conn, requestCtx, err := grpcClient(ctx)
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
