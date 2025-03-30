package handler

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/resolver"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/service/auth"
)

type LoginArgs struct {
	Input struct {
		Email    string
		Password string
	}
}

func (q *Query) Login(ctx context.Context, args LoginArgs) *resolver.LoginPayloadResolver {
	userRepository := repository.NewUserRepository()
	userService := auth.NewLoginService(userRepository)

	token, user, err := userService.Login(auth.LoginInput{
		Email:    args.Input.Email,
		Password: args.Input.Password,
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return resolver.NewLoginPayloadResolver(ctx, token, user)
}

type LoginInput struct {
	Input struct {
		Email    string
		Password string
	}
}

type LoginService struct {
	userService auth.LoginService
}

func NewLoginService(userService auth.LoginService) *LoginService {
	return &LoginService{userService: userService}
}

func (s *LoginService) Login(ctx context.Context, args LoginInput) (*resolver.LoginPayloadResolver, error) {
	token, user, err := s.userService.Login(auth.LoginInput{
		Email:    args.Input.Email,
		Password: args.Input.Password,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewLoginPayloadResolver(ctx, token, user), nil
}
