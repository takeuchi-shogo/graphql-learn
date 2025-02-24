package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type PostEdge struct {
	cursor graphql.ID
	model  *entity.Post
	ctx    context.Context
}

func (r *PostEdge) Cursor() graphql.ID {
	return r.cursor
}

func (r *PostEdge) Node(ctx context.Context) *PostResolver {
	return NewPostResolver(r.ctx, r.model)
}
