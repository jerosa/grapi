package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jerosa/grapi/internal/server"
	"github.com/jerosa/grapi/internal/server/grpc"
	"github.com/jerosa/grapi/internal/server/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		protocol = "tcp"
		host     = "localhost"
		port     = "3333"
	)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		srvCfg := server.Config{Protocol: protocol, Host: host, Port: port}
		srv := grpc.NewServer(srvCfg)

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
