package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sandbox/grpc/pb/customer"
	"sandbox/grpc/pb/todo"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RunServer ..
func RunServer(ctx context.Context, grpcPort, httpPort string, opts ...runtime.ServeMuxOption) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	// customer
	if err := customer.RegisterCustomerHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, dialOpts); err != nil {
		return err
	}

	// todo
	if err := todo.RegisterTodoHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, dialOpts); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Printf("%v\n", c)
		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
