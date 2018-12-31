package main

import (
	"context"
	"log"
	"net"
	"sandbox/grpc/customer"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	savedCustomers []*customer.Request
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	customer.RegisterCustomerServer(s, &server{})
	log.Printf("Listening on port %v\n", port)
	s.Serve(listen)
}

func (s *server) Create(ctx context.Context, in *customer.Request) (*customer.Response, error) {
	log.Println("Creating customers")
	s.savedCustomers = append(s.savedCustomers, in)
	return &customer.Response{Id: in.Id, Success: true}, nil
}

func (s *server) Get(filter *customer.Filter, stream customer.Customer_GetServer) error {
	log.Println("Get customers")
	for _, cust := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(cust.Name, filter.Keyword) {
				continue
			}
		}

		if err := stream.Send(cust); err != nil {
			return err
		}
	}
	return nil
}
