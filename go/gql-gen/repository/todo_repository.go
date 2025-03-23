package repository

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
)

type ITodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Get(ctx context.Context, id string) (*model.Todo, error)
	GetList(ctx context.Context, first *int, after *string, last *int, before *string, isReverse *bool) ([]*model.Todo, error)
}

type TodoRepository struct {
	// db *sql.DB
}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	todoEntity := entity.NewTodo(todo)
	log.Println("Create todo", todoEntity)
	return todo, nil
}

func (r *TodoRepository) Get(ctx context.Context, id string) (*model.Todo, error) {
	todoEntity := entity.NewTodo(&model.Todo{
		ID:    id,
		Title: "test",
	})
	return &model.Todo{
		ID:    todoEntity.ID,
		Title: todoEntity.Title,
	}, nil
}

func (r TodoRepository) GetList(ctx context.Context, first *int, after *string, last *int, before *string, isReverse *bool) ([]*model.Todo, error) {
	log.Println("Get list", first, after, last, before, isReverse)
	return []*model.Todo{
		{
			ID:    "1",
			Title: "test",
		},
	}, nil
}
