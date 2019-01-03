package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sandbox/grpc/pb/customer"
	"sandbox/grpc/pb/todo"

	"google.golang.org/grpc"
)

// OptionFunc ..
type OptionFunc func(*grpc.Server)

// WithCustomerServer ..
func WithCustomerServer(customerSrv customer.CustomerServer) OptionFunc {
	return func(srv *grpc.Server) {
		customer.RegisterCustomerServer(srv, customerSrv)
	}
}

// WithTodoServer ..
func WithTodoServer(todoSrv todo.TodoServer) OptionFunc {
	return func(srv *grpc.Server) {
		todo.RegisterTodoServer(srv, todoSrv)
	}
}

// RunServer ..
func RunServer(ctx context.Context, port string, optFuncs ...OptionFunc) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	// customer.RegisterCustomerServer(server, cus)
	for _, optFunc := range optFuncs {
		optFunc(server)
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC servier...")
	return server.Serve(listen)
}
