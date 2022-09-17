package birdie

import (
	"github.com/shamanis/birdie/internal/pkg/logging"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	DefaultNetwork = "tcp"
	DefaultAddr    = "0.0.0.0:50051"
)

var (
	listenAddr    = os.Getenv("LISTEN_ADDR")
	listenNetwork = os.Getenv("LISTEN_NETWORK")
	opts          []grpc.ServerOption
	logger        = logging.New()
)

func StartService() {
	if listenNetwork == "" {
		listenNetwork = DefaultNetwork
	}
	if listenAddr == "" {
		listenAddr = DefaultAddr
	}

	listener, err := net.Listen(listenNetwork, listenAddr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(opts...)

	RegisterBirdieServer(grpcServer, &Service{})
	logger.Infof("Start gRPC server on %s://%s", listenNetwork, listenAddr)
	grpcServer.Serve(listener)
}
