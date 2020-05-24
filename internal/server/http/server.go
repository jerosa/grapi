package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/jerosa/grapi/proto"
	"google.golang.org/grpc"
)

type Server struct {
	httpAddr string
}

func NewServer(httpAddr string) *Server {
	return &Server{httpAddr: httpAddr}
}

func (s *Server) Serve(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterReadListServiceHandlerFromEndpoint(ctx, mux, s.httpAddr, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(s.httpAddr, mux)
}
