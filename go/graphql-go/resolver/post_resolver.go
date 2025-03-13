package resolver

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/graphql-go"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/loader"
)

type PostResolver struct {
	ctx  context.Context
	post *entity.Post
}

func NewPostResolver(ctx context.Context, post *entity.Post) *PostResolver {
	return &PostResolver{ctx: ctx, post: post}
}

func (r *PostResolver) ID() graphql.ID {
	return graphql.ID(r.post.ID())
}

func (r *PostResolver) Content() string {
	return r.post.Content()
}

func (r *PostResolver) Author() *UserResolver {
	fmt.Println("resolver.User()", r.post.AuthorID())
	user, err := loader.LoadUser(r.ctx, r.post.AuthorID())
	if err != nil {
		log.Println(err)
		return nil
	}

	return NewUserResolver(user)
}
