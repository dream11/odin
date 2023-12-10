package service

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/dream11/odin/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func grpcClient(ctx *context.Context) (*grpc.ClientConn, *context.Context, error) {
	appConfig := config.GetConfig()
	var opts []grpc.DialOption
	if appConfig.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		cred := credentials.NewTLS(&tls.Config{})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	conn, err := grpc.Dial(appConfig.BackendAddress, opts...)

	if err != nil {
		return nil, nil, err
	}
	// Enrich context with authorisation metadata
	requestCtx := metadata.AppendToOutgoingContext(*ctx, "Authorization", fmt.Sprintf("Bearer %s", appConfig.AccessToken))
	return conn, &requestCtx, nil
}
