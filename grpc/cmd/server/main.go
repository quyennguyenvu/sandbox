package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sandbox/grpc/cmd/server/protocol/grpc"
	"sandbox/grpc/cmd/server/protocol/rest"
	"sandbox/grpc/handler"

	"github.com/joho/godotenv"
)

func main() {
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer ..
func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpPort := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("GRPC_PORT")

	ctx := context.Background()

	go func() {
		_ = rest.RunServer(ctx, grpcPort, httpPort)
	}()

	withServer := []grpc.OptionFunc{
		grpc.WithCustomerServer(handler.NewCustomerHandler()),
		grpc.WithTodoServer(handler.NewTodoHandler()),
	}

	return grpc.RunServer(ctx, grpcPort, withServer...)
}
