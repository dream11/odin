package service

import (
	"context"

	auth "github.com/dream11/odin/proto/gen/go/dream11/od/auth/v1"
)

// Configure used to perform odin congigure
type Configure struct{}

// GetUserToken Get User Token
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
