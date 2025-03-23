package auth

import (
	"context"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/repository"
)

type SignupService struct {
	userRepository repository.IUserRepository
}

func NewSignupService(userRepository repository.IUserRepository) *SignupService {
	return &SignupService{userRepository: userRepository}
}

func (s *SignupService) Signup(ctx context.Context, input model.UserCreateInput) (*model.UserCreatePayload, error) {
	user, err := s.userRepository.Create(ctx, &model.User{
		ID:    "1",
		Name:  input.Name,
		Email: input.Email,
		Role:  model.UserRoleUser,
	})
	if err != nil {
		return nil, err
	}
	return &model.UserCreatePayload{
		Entity: user,
	}, nil
}
