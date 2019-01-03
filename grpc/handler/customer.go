package handler

import (
	"context"
	"log"
	"sandbox/grpc/pb/customer"
	"strings"

	"github.com/golang/protobuf/ptypes"
)

// CustomerHandler ..
type CustomerHandler interface {
	List(filter *customer.ListFilter, stream customer.Customer_ListServer) error
	Create(ctx context.Context, in *customer.Request) (*customer.Response, error)
}

type customerHandlerImpl struct {
	savedCustomers []*customer.Request
}

// NewCustomerHandler ..
func NewCustomerHandler() CustomerHandler {
	return &customerHandlerImpl{}
}

func (s *customerHandlerImpl) Create(ctx context.Context, in *customer.Request) (*customer.Response, error) {
	log.Println("Creating customer")
	in.CreatedAt = ptypes.TimestampNow()
	s.savedCustomers = append(s.savedCustomers, in)
	return &customer.Response{Id: in.Id, Success: true}, nil
}

func (s *customerHandlerImpl) List(filter *customer.ListFilter, stream customer.Customer_ListServer) error {
	log.Println("List customers")
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
