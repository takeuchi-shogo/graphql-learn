package user

import (
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type GetUserService struct{}

func NewGetUserService() *GetUserService {
	return &GetUserService{}
}

func (s *GetUserService) GetUser(id string) (*entity.User, error) {
	return entity.NewUser(
		id,
		"john_doe",
		"John Doe",
		"I'm a developer",
		"https://example.com/avatar.png",
	), nil
}
