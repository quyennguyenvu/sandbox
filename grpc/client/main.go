package main

import (
	"context"
	"io"
	"log"
	"sandbox/grpc/customer"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := customer.NewCustomerClient(conn)

	cust := &customer.Request{
		Id:    2,
		Name:  "Quyen",
		Phone: "1023213",
		Addresses: []*customer.Request_Address{
			&customer.Request_Address{
				Street:            "ton duc thang",
				City:              "Hanoi",
				State:             "HN",
				Zip:               "100000",
				IsShippingAddress: false,
			},
			&customer.Request_Address{
				Street:            "ba huyen thanh quan",
				City:              "Saigon",
				State:             "SG",
				Zip:               "70000",
				IsShippingAddress: true,
			},
		},
	}

	createCustomer(client, cust)

	cust = &customer.Request{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*customer.Request_Address{
			&customer.Request_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, cust)
	filter := &customer.Filter{Keyword: ""}
	getCustomers(client, filter)

}

func createCustomer(client customer.CustomerClient, customer *customer.Request) {
	resp, err := client.Create(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new customer has been added with id: %d", resp.Id)
	}
}

func getCustomers(client customer.CustomerClient, filter *customer.Filter) {
	stream, err := client.Get(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		cust, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Get(_) = _, %v", client, err)
		}
		log.Printf("\nCustomer: %v", cust)
	}
}
