package service

import (
	"context"

	auth "github.com/dream11/odin/proto/gen/go/dream11/od/auth/v1"
)

// Environment performs operation on environment like create, list, describe, delete
type Configure struct{}

// ListEnvironments List environments
func (c *Configure) GetUserToken(ctx *context.Context, request *auth.GetUserTokenRequest) (*auth.GetUserTokenResponse, error) {
	conn, requestCtx, err := grpcClient(ctx)
	if err != nil {
		return nil, err
	}
	client := auth.NewAuthServiceClient(conn)
	response, err := client.GetUserToken(*requestCtx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
