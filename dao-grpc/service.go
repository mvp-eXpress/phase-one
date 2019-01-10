package todo

import (
	"context"
)

type IService interface {
	CreateTodo(ctx context.Context, i *NewTodoInput) (*Todo, error)
	// CreateTodos(ctx context.Context, i *[]Todo) (*Todo, error)
	// GetTodo(ctx context.Context) (*Todo, error)
	// GetTodos(ctx context.Context) ([]Todo, error)
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoService struct {
	r IRepo
}

type IRepo interface {
	CreateTodo(ctx context.Context, i *NewTodoInput) (*Todo, error)
}

func NewService(r IRepo) IService {
	return &TodoService{r}
}

func (s *TodoService) CreateTodo(ctx context.Context, n *NewTodoInput) (*Todo, error) {

	t, err := s.r.CreateTodo(ctx, n)
	if err != nil {
		return nil, err
	}
	return t, nil

}

type NewTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
