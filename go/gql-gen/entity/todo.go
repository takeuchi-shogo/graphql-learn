package entity

import (
	"time"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodo(todo *model.Todo) *Todo {
	return &Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: *todo.Description,
		Status:      string(todo.Status),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
