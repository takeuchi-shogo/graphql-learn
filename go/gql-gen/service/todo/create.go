package todo

import (
	"context"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/repository"
)

type CreateTodoService struct {
	todoRepository repository.ITodoRepository
}

func NewCreateTodoService(todoRepository repository.ITodoRepository) *CreateTodoService {
	return &CreateTodoService{
		todoRepository: todoRepository,
	}
}

func (s *CreateTodoService) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := s.todoRepository.Create(ctx, &model.Todo{
		ID:    "1",
		Title: input.Title,
	})
	if err != nil {
		return nil, err
	}
	return todo, nil
}
