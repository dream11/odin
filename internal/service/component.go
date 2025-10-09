package service

import (
	"context"
	"time"

	"github.com/avast/retry-go"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
)

// Component performs operation on component like operate
type Component struct{}

// OperateComponent operate Component
func (e *Component) OperateComponent(ctx *context.Context, request *serviceProto.OperateServiceRequest) error {
	log.Infof(constant.ComponentExecutionMessageTemplate, "Operating", request.GetComponentName(), request.GetEnvName())

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

// CompareOperationChanges compares the operation changes
func (e *Component) CompareOperationChanges(ctx *context.Context, request *serviceProto.OperateComponentDiffRequest) (*serviceProto.OperateComponentDiffResponse, error) {

	for retries := 0; retries < constant.MaxRetries; retries++ {
		ctxWithTimeout, cancel := context.WithTimeout(*ctx, constant.Delay)
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
		time.Sleep(constant.Delay)
		if retries == 0 {
			log.Warnf(constant.InitiatingRetryMessage)
		}
		log.Infof(constant.RetryingMessage, retries+1, constant.MaxRetries)
	}

	log.Fatalf(constant.MaxRetriesReached)
	return nil, nil
}
