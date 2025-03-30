package resolver

import (
	"context"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type CreatePostPayloadResolver struct {
	post *entity.Post
	ctx  context.Context
}

func NewCreatePostPayloadResolver(ctx context.Context, post *entity.Post) *CreatePostPayloadResolver {
	return &CreatePostPayloadResolver{
		ctx:  ctx,
		post: post,
	}
}

func (r *CreatePostPayloadResolver) Post() *PostResolver {
	return NewPostResolver(r.ctx, r.post)
}
