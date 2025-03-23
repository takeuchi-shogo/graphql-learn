package todo

import (
	"context"
	"errors"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/repository"
)

type GetTodoListService struct {
	todoRepository repository.ITodoRepository
}

type GetTodoListInput struct {
	First     *int
	After     *string
	Last      *int
	Before    *string
	IsReverse *bool
}

func NewGetTodoListService(todoRepository repository.ITodoRepository) *GetTodoListService {
	return &GetTodoListService{
		todoRepository: todoRepository,
	}
}

func NewGetTodoListInput(first *int32, after *string, last *int32, before *string, isReverse *bool) GetTodoListInput {
	var firstInt *int
	if first != nil {
		value := int(*first)
		firstInt = &value
	}
	var lastInt *int
	if last != nil {
		value := int(*last)
		lastInt = &value
	}

	return GetTodoListInput{
		First:     firstInt,
		After:     after,
		Last:      lastInt,
		Before:    before,
		IsReverse: isReverse,
	}
}

func (s *GetTodoListService) GetTodoList(ctx context.Context, input GetTodoListInput) ([]*model.Todo, error) {
	if s.todoRepository == nil {
		return nil, errors.New("todoRepository is nil")
	}
	todoList, err := s.todoRepository.GetList(ctx, input.First, input.After, input.Last, input.Before, input.IsReverse)
	if err != nil {
		return nil, err
	}
	return todoList, nil
}
