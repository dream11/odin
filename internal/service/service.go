package service

import (
	"context"

	service "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
)

// Service performs operation on service like list, describe
type Service struct{}

// DescribeSerice Describe service
func (e *Service) DescribeService(ctx *context.Context, request *service.DescribeServiceRequest) (*service.DescribeServiceResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := service.NewServiceServiceClient(conn)
	response, err := client.DescribeService(*requestCtx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
