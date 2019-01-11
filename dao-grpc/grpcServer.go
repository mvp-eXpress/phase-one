package todo

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mvp-eXpress/phase-one/dao-grpc/todopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service IService
}

func ListenGRPC(s IService, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	todopb.RegisterTodoServiceServer(server, &grpcServer{s})
	reflection.Register(server)
	return server.Serve(lis)

}

func (s *grpcServer) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {

	todoItem := NewTodoInput{
		Title:       req.Item.Title,
		Description: req.Item.Description,
		Completed:   req.Item.Completed,
	}

	a, err := s.service.CreateTodo(ctx, &todoItem)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &todopb.CreateTodoResponse{
		Todo: &todopb.Todo{
			Id:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			Completed:   a.Completed,
		},
	}, nil

}
