package handler

import (
	"context"
	"log"
	"sandbox/grpc/pb/todo"
	"strings"

	"github.com/golang/protobuf/ptypes"
)

// TodoHandler ..
type TodoHandler interface {
	List(filter *todo.ListFilter, stream todo.Todo_ListServer) error
	Create(ctx context.Context, in *todo.Request) (*todo.Response, error)
}

type todoHandlerImpl struct {
	savedTodo []*todo.Request
}

// NewTodoHandler ..
func NewTodoHandler() TodoHandler {
	return &todoHandlerImpl{}
}

func (s *todoHandlerImpl) Create(ctx context.Context, in *todo.Request) (*todo.Response, error) {
	log.Println("Creating todo")
	in.Reminder = ptypes.TimestampNow()
	s.savedTodo = append(s.savedTodo, in)
	return &todo.Response{Id: in.Id, Success: true}, nil
}

func (s *todoHandlerImpl) List(filter *todo.ListFilter, stream todo.Todo_ListServer) error {
	log.Println("List todos")
	for _, todo := range s.savedTodo {
		if filter.Keyword != "" {
			if !strings.Contains(todo.Title, filter.Keyword) {
				continue
			}
		}

		if err := stream.Send(todo); err != nil {
			return err
		}
	}
	return nil
}
