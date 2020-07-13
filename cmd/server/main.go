package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jerosa/grapi/internal/adding"
	"github.com/jerosa/grapi/internal/creating"
	"github.com/jerosa/grapi/internal/listing"
	"github.com/jerosa/grapi/internal/server"
	"github.com/jerosa/grapi/internal/server/grpc"
	"github.com/jerosa/grapi/internal/server/http"
	"github.com/jerosa/grapi/internal/storage/inmemory"

	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		protocol = "tcp"
		host     = "localhost"
		port     = "3333"

		repo            = inmemory.NewInMemoryReadListRepository()
		creatingService = creating.NewService(repo)
		addingService   = adding.NewService(repo)
		listingService  = listing.NewService(repo)
	)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		srvCfg := server.Config{Protocol: protocol, Host: host, Port: port}
		srv := grpc.NewServer(srvCfg, creatingService, addingService, listingService)

		log.Printf("gRPC server running at %s://%s:%s ...\n", protocol, host, port)
		return srv.Serve()
	})
	g.Go(func() error {
		httpAddr := fmt.Sprintf(":%s", port)
		httpSrv := http.NewServer(httpAddr)

		log.Printf("HTTP server running at %s ...\n", httpAddr)
		return httpSrv.Serve(ctx)
	})

	log.Fatal(g.Wait())
}
