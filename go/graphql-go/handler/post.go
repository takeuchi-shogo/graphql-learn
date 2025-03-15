package handler

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/resolver"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/service/post"
)

type PostsArgs struct {
	First *float64
	After *string
}

func (q *Query) Posts(ctx context.Context, args PostsArgs) *resolver.PostConnectionResolver {
	postRepository := repository.NewPostRepository()
	postService := post.NewGetListService(postRepository)

	var first *int
	if args.First != nil {
		val := int(*args.First)
		first = &val
	}

	posts, err := postService.GetList(post.GetListInput{
		First: first,
		After: args.After,
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return resolver.NewPostConnectionResolver(ctx, posts)
}
