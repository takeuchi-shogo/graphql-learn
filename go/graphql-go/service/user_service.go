package service

import (
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	return entity.NewUser(id, "John Doe"), nil
}
