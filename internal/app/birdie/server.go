package birdie

import (
	"github.com/shamanis/birdie/internal/pkg/logging"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	DefaultPort = ":50051"
)

var (
	listenPort = os.Getenv("LISTEN_PORT")
	opts       []grpc.ServerOption
	logger     = logging.New()
)

func StartService() {
	if listenPort == "" {
		listenPort = DefaultPort
	}
	listener, err := net.Listen("tcp", listenPort)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(opts...)

	RegisterBirdieServer(grpcServer, &Service{})
	logger.Infof("Start gRPC server on %s", listenPort)
	grpcServer.Serve(listener)
}
