package handler

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/resolver"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/service/user"
)

type UserArgs struct {
	ID string
}

func (q *Query) User(ctx context.Context, args UserArgs) *resolver.UserResolver {
	userService := user.NewGetUserService()
	user, err := userService.GetUser(args.ID)
	if err != nil {
		log.Println(err)
		return nil
	}

	return resolver.NewUserResolver(user)
}
