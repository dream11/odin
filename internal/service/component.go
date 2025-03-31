package service

import (
	"context"
	"fmt"
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
const MaxRetries = 10
const Timeout =20 * time.Second
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
	var maxRetries = MaxRetries
	var retries = 0
	outerLoop:
	for {
		// Create a context with timeout for each Recv call
		recvCtx, cancel := context.WithTimeout(*requestCtx, Timeout)

		responseChan := make(chan *serviceProto.OperateServiceResponse)
		errorChan := make(chan error)

		go func() {
			response, err := stream.Recv()
			if err != nil {
				errorChan <- err
			} else {
				responseChan <- response
			}
		}()

		select {
		case <-recvCtx.Done():
			spinnerInstance.Stop()
			cancel()
			log.Error("Operation timed out. Retrying again")
			if retries < maxRetries {
				var err error
				retries++
				log.Infof("Retrying ... (%d/%d)", retries, maxRetries)
				stream, err = retryOperateStream(client, requestCtx, request, stream)
				if err != nil {
					return err
				}
				continue outerLoop
			}
			log.Errorf("TraceID: %s, error: %v", (*requestCtx).Value(constant.TraceIDKey), recvCtx.Err())
			return recvCtx.Err()
		case err := <-errorChan:
			spinnerInstance.Stop()
			cancel()
			if !util.IsRetryable(err) {
				break outerLoop
			}
			if retries < maxRetries {
				var err error
				retries++
				log.Infof("Retrying ... (%d/%d)", retries, maxRetries)
				stream, err = retryOperateStream(client, requestCtx, request, stream)
				if err != nil {
					return err
				}
				continue outerLoop
			}
			log.Errorf("TraceID: %s", (*requestCtx).Value(constant.TraceIDKey))
			return err
		case response := <-responseChan:
			spinnerInstance.Stop()
			cancel()
			if response != nil {
				message = util.GenerateResponseMessageComponentSpecific(response.GetServiceResponse(), []string{request.GetComponentName()})
				logFailedComponentMessagesOnceForComponents(response.GetServiceResponse(), []string{request.GetComponentName()})
				spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
				spinnerInstance.Start()
			}
		}
	}
	log.Info(message)
	return nil
}

func retryOperateStream(client serviceProto.ServiceServiceClient, requestCtx *context.Context, request *serviceProto.OperateServiceRequest, stream serviceProto.ServiceService_OperateServiceClient) (serviceProto.ServiceService_OperateServiceClient, error) {

	// Close the current stream
	if err := stream.CloseSend(); err != nil {
		log.Errorf("Failed to close stream: %v", err)
		return nil, err
	}

	// Create a new stream connection
	newStream, err := client.OperateService(*requestCtx, request)
	if err != nil {
		log.Errorf("Failed to create new stream: %v", err)
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

	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	response, err := client.OperateComponentDiff(*requestCtx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
