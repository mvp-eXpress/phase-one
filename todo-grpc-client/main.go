package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mvp-eXpress/phase-one/dao-grpc/todopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Todo client initializing...")
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := todopb.NewTodoServiceClient(cc)

	doCreateTodo(c)

}

func doCreateTodo(c todopb.TodoServiceClient) {
	fmt.Println("Starting createtodo rpc...")
	req := &todopb.CreateTodoRequest{
		Item: &todopb.NewTodoRequest{
			Title:       "3 yo yo title",
			Description: "3 as;hdha; sdesc",
		},
	}

	res, err := c.CreateTodo(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling createtodo RPC: %v", err)
	}
	log.Printf("Response from createtodo rpc: %v", res.Todo)
}
