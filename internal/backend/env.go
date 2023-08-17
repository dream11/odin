package backend

import (
	"context"
	environment "github.com/dream11/odin/proto/gen/go/dream11/od/environment/v1"
)

type Env struct{}

func (e *Env) ListEnvironments(ctx context.Context, request *environment.ListEnvironmentRequest) (*environment.ListEnvironmentResponse, error) {
	conn, err := grpcClient()
	if err != nil {
		return nil, err
	}
	client := environment.NewEnvironmentServiceClient(conn)
	response, err := client.ListEnvironment(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
