package handler

import (
	"context"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/resolver"
)

func (q *Query) Posts(ctx context.Context) *resolver.PostConnectionResolver {
	return resolver.NewPostConnectionResolver(ctx, []*entity.Post{
		entity.NewPost("1", "Post 1", "1"),
		entity.NewPost("2", "Post 2", "1"),
	})
}
