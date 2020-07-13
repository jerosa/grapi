package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	readlistgrpc "github.com/jerosa/grapi/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var cli readlistgrpc.ReadListServiceClient
var ctx context.Context

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "readList",
	Short: "Simple gRPC read list",
}

// Execute adds all child commands to the root command and set flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	var (
		host = "localhost"
		port = "3333"
	)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}

	cli = readlistgrpc.NewReadListServiceClient(conn)
}
