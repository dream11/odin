package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func grpcClient(ctx *context.Context) (*grpc.ClientConn, *context.Context, error) {
	appConfig := config.GetConfig()
	var opts []grpc.DialOption
	if appConfig.Insecure {
		if util.IsIPAddress(strings.Split(appConfig.BackendAddress, ":")[0]) {
			// Disable TLS
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		} else {
			// Perform TLS handshake but skip certificate verification
			var tlsConf tls.Config
			tlsConf.InsecureSkipVerify = true
			opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConf)))
		}

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
