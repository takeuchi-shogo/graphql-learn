package handler

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/resolver"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/service/auth"
)

type LoginArgs struct {
	Email    string
	Password string
}

func (q *Query) Login(ctx context.Context, args LoginArgs) *resolver.LoginPayloadResolver {
	userRepository := repository.NewUserRepository()
	userService := auth.NewLoginService(userRepository)

	token, err := userService.Login(auth.LoginInput{
		Email:    args.Email,
		Password: args.Password,
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return resolver.NewLoginPayloadResolver(ctx, token)
}
