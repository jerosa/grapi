package grpc

import (
	"fmt"
	"net"
	"os"

	"github.com/jerosa/grapi/internal/adding"
	"github.com/jerosa/grapi/internal/creating"
	"github.com/jerosa/grapi/internal/listing"
	"github.com/jerosa/grapi/internal/server"
	pb "github.com/jerosa/grapi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type grpcServer struct {
	config          server.Config
	creatingService creating.Service
	addingService   adding.Service
	listingService  listing.Service
}

// NewServer creates a new server
func NewServer(
	config server.Config,
	cs creating.Service,
	aS adding.Service,
	lS listing.Service,
) server.Server {
	return &grpcServer{config: config, creatingService: cs, addingService: aS, listingService: lS}
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
	serviceServer := NewReadListServer(
		s.creatingService,
		s.addingService,
		s.listingService,
	)
	pb.RegisterReadListServiceServer(srv, serviceServer)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}
