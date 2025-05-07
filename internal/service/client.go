package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/dream11/odin/api/configuration"
	"strings"
	"time"

	"github.com/dream11/odin/pkg/config"
	"github.com/dream11/odin/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

func getTLSOpts(appConfig *configuration.Configuration) grpc.DialOption {
	tlsConf := tls.Config{
		ServerName: strings.Split(appConfig.BackendAddress, ":")[0],
	}
	if appConfig.Plaintext {
		// Disable TLS
		return grpc.WithTransportCredentials(insecure.NewCredentials())
	}
	if appConfig.Insecure {
		// Perform TLS handshake but skip certificate verification
		tlsConf.InsecureSkipVerify = true
	}

	return grpc.WithTransportCredentials(credentials.NewTLS(&tlsConf))

}

func grpcClient(ctx *context.Context) (*grpc.ClientConn, *context.Context, error) {
	appConfig := config.GetConfig()



	if appConfig.BackendAddress == "" {
		log.Fatal("Cannot create grpc client: Backend address is empty in config! Run `odin configure` to set backend address")
	}

	opts := []grpc.DialOption{
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                10 * time.Second,
				Timeout:             20 * time.Second,
				PermitWithoutStream: true,
			}),
		getTLSOpts(appConfig),

	}
	conn, err := grpc.NewClient(appConfig.BackendAddress, opts...)
	if err != nil {
		return nil, nil, err
	}

	// Enrich context with authorization metadata only
	requestCtx := metadata.AppendToOutgoingContext(*ctx, "Authorization", fmt.Sprintf("Bearer %s", appConfig.AccessToken))
	return conn, &requestCtx, nil
}
