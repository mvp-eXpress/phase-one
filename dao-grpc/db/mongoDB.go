package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mvp-eXpress/phase-one/dao-grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongoConn(url string) (*MongoDB, error) {
	fmt.Println("Connecting to MongoDB")

	// connect to MongoDB
	client, err := mongo.NewClient(url)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo Mongo DB Collection created")
	collection := client.Database("todo").Collection("todo")
	return &MongoDB{collection}, nil
}

func (m *MongoDB) CreateTodo(ctx context.Context, i *todo.NewTodoInput) (*todo.Todo, error) {
	fmt.Println("Starting CreateTodo in MongoDB")

	res, err := m.collection.InsertOne(context.Background(), i)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &todo.Todo{
		ID:          oid.String(),
		Description: i.Description,
		Title:       i.Title,
		Completed:   false,
	}, nil
}
