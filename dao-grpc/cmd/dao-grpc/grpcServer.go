package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mvp-eXpress/phase-one/dao-grpc/todopb"
	"github.com/mvp-eXpress/phase-one/dao-grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service todo.IService
}

// func ListenGRPC(s IService, port int) error {
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	server := grpc.NewServer()
// 	todo.RegisterTodoServiceServer(server, &grpcServer{s})
// 	reflection.Register(server)
// 	return server.Serve(lis)
// }

func (s *grpcServer) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {

	 todoItem:=todo.NewTodoInput{
		 Title:req.Item.Title,
		 Description: req.Item.Description,
	 }
	 
	a, err := s.service.CreateTodo(ctx, &todoItem)
		if err != nil {
			log.Fatal(err)
		return nil, err
	}
	return &todopb.CreateTodoResponse{
		Id:"1A",
	},nil

}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Initializing Mongo service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	todopb.RegisterTodoServiceServer(server, &grpcServer{s})
	reflection.Register(server)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

		// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing MongoDB Connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of Program")
}
