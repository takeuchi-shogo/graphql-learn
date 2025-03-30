package handler

import (
	"context"
	"fmt"
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

type CreatePostArgs struct {
	Input struct {
		Content  string
		AuthorID string
	}
}

func (q *Query) CreatePost(ctx context.Context, args CreatePostArgs) *resolver.CreatePostPayloadResolver {
	fmt.Println("CreatePost", args)

	postRepository := repository.NewPostRepository()
	postService := post.NewCreateService(postRepository)

	post, err := postService.Create(post.CreateInput{
		Content:  args.Input.Content,
		AuthorID: args.Input.AuthorID,
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return resolver.NewCreatePostPayloadResolver(ctx, post)
}
