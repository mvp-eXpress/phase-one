package todo

import (
	"context"
)

type IService interface {
	CreateTodo(ctx context.Context, i *NewTodoInput) (*Todo, error)
	CreateTodos(ctx context.Context, i *[]Todo) (*Todo, error)
	GetTodo(ctx context.Context) (*Todo, error)
	GetTodos(ctx context.Context) ([]Todo, error)
}

type NewTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoService struct {
	// db DB
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
