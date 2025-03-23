package utils

import (
	"strconv"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
)

func CreateTodoConnection(todos []*model.Todo, userErrors []*model.UserError) *model.TodoConnection {
	edges := make([]*model.TodoEdge, len(todos))
	for i, todo := range todos {
		edges[i] = &model.TodoEdge{
			Node:   todo,
			Cursor: strconv.Itoa(i),
		}
	}
	return &model.TodoConnection{
		Edges: edges,
		Nodes: todos,
		PageInfo: &model.PageInfo{
			HasNextPage:     true,
			HasPreviousPage: true,
		},
		UserErrors: userErrors,
	}
}
