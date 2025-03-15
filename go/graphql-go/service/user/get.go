package user

import (
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
)

type GetUserService struct {
	userRepository repository.UserRepositoryImpl
}

func NewGetUserService(userRepository repository.UserRepositoryImpl) *GetUserService {
	return &GetUserService{userRepository: userRepository}
}

func (s *GetUserService) GetUser(id string) (*entity.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
