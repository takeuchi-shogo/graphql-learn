package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type PostConnectionResolver struct {
	ctx   context.Context
	posts []*entity.Post
}

func NewPostConnectionResolver(ctx context.Context, posts []*entity.Post) *PostConnectionResolver {
	return &PostConnectionResolver{ctx: ctx, posts: posts}
}

func (r *PostConnectionResolver) Edges() []*PostEdge {
	p := make([]*PostEdge, len(r.posts))
	for i, post := range r.posts {
		p[i] = &PostEdge{
			cursor: graphql.ID(post.ID()),
			model:  post,
			ctx:    r.ctx,
		}
	}
	return p
}

func (r *PostConnectionResolver) Nodes() []*PostResolver {
	p := make([]*PostResolver, len(r.posts))
	for i, post := range r.posts {
		p[i] = NewPostResolver(r.ctx, post)
	}
	return p
}
