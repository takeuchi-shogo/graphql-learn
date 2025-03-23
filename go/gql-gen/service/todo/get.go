package todo

import (
	"context"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/repository"
)

type GetTodoService struct {
	todoRepository repository.ITodoRepository
}

func NewGetTodoService(todoRepository repository.ITodoRepository) *GetTodoService {
	return &GetTodoService{todoRepository: todoRepository}
}

func (s *GetTodoService) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	todo, err := s.todoRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
