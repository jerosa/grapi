package grpc

import (
	"fmt"
	"net"
	"os"

	"github.com/jerosa/grapi/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type grpcServer struct {
	config server.Config
}

func NewServer(
	config server.Config,
) server.Server {
	return &grpcServer{config: config}
}

func (s *grpcServer) Serve() error {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	srv := grpc.NewServer()

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}
