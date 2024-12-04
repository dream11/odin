package service

import (
	"context"
	"errors"
	"fmt"
	"io"

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
			message = util.GenerateResponseMessageComponentSpecific(response.GetServiceResponse(), []string{request.GetComponentName()})
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}
	log.Info(message)
	return err
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
		return nil, err
	}

	return response, nil
}

func (c *Component) CompareOperationChanges(ctx *context.Context, request *serviceProto.CompareOperationChangesRequest) (*serviceProto.CompareOperationChangesResponse, error) {

	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := serviceProto.NewServiceServiceClient(conn)
	response, err := client.CompareOperationChanges(*requestCtx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
