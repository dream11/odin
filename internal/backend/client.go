package backend

import (
	"crypto/tls"
	"github.com/dream11/odin/api/configuration"
	"github.com/dream11/odin/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var appConfig *configuration.Configuration

func init() {
	appConfig = config.GetConfig()
}

func grpcClient() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if appConfig.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		creds := credentials.NewTLS(&tls.Config{})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(appConfig.BackendAddr, opts...)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
