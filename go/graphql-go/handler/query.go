package handler

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/loader"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/resolver"
)

type Query struct{}

func (q *Query) Hello() string {
	return "Hello, World!"
}

func (q *Query) Posts(ctx context.Context) *resolver.PostConnectionResolver {
	return resolver.NewPostConnectionResolver(ctx, []*entity.Post{
		entity.NewPost("1", "Post 1", "1"),
		entity.NewPost("2", "Post 2", "1"),
	})
}

func (q *Query) User(args struct{ ID string }) *resolver.UserResolver {
	user, err := loader.LoadUser(context.Background(), args.ID)
	if err != nil {
		log.Println(err)
		return nil
	}

	return resolver.NewUserResolver(user)
}
