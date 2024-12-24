package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func grpcClient(ctx *context.Context, traceIDOptional ...string) (*grpc.ClientConn, *context.Context, error) {
	appConfig := config.GetConfig()

	traceID := ""
	if len(traceIDOptional) > 0 {
		traceID = traceIDOptional[0]
	}

	if appConfig.BackendAddress == "" {
		log.Fatal("Cannot create grpc client: Backend address is empty in config! Run `odin configure` to set backend address")
	}
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
	requestCtx := metadata.AppendToOutgoingContext(*ctx, "Authorization", fmt.Sprintf("Bearer %s", appConfig.AccessToken), "TraceId", traceID)
	return conn, &requestCtx, nil
}
