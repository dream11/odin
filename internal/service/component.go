package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	component "github.com/dream11/odin/proto/gen/go/dream11/od/component/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var maxRetries = 3
	var retries = 0
outerLoop:
	for {
		// Create a context with timeout for each Recv call
		recvCtx, cancel := context.WithTimeout(*requestCtx, 30*time.Second)

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
			log.Error("Operation timed out. Retry again")
			log.Errorf("TraceID: %s, error: %v", (*requestCtx).Value(constant.TraceIDKey), recvCtx.Err())
			cancel()
			return recvCtx.Err()
		case err := <-errorChan:
			spinnerInstance.Stop()
			cancel()
			if !isRetryable(err) {
				break outerLoop
			}
			if retries < maxRetries {
				log.Errorf("Error: %v", err)
				retries++
				time.Sleep(5 * time.Second)

				// Close the current stream
				if err := stream.CloseSend(); err != nil {
					log.Errorf("Failed to close stream: %v", err)
					return err
				}

				// Create a new stream connection
				stream, err = client.OperateService(*requestCtx, request)
				if err != nil {
					log.Errorf("Failed to create new stream: %v", err)
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

func isRetryable(err error) bool {
	if errors.Is(err, context.Canceled)  {
		return true
	}
	if err == io.EOF {
		return false
	}


	st, ok := status.FromError(err)
	if !ok {
		return false
	}

	if st.Code() == codes.Unavailable {
		return true
	}

	if st.Code() == codes.Internal && strings.Contains(st.Message(), "RST_STREAM") {
		return true
	}

	return false
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
