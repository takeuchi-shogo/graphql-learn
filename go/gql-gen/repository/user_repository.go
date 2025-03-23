package repository

import (
	"context"
	"log"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Get(ctx context.Context, id string) (*model.User, error)
}

type UserRepository struct {
	// db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	userEntity := entity.NewUser(user)
	log.Println("Create user", userEntity)
	return user, nil
}

func (r *UserRepository) Get(ctx context.Context, id string) (*model.User, error) {
	userEntity := entity.NewUser(&model.User{
		ID:    id,
		Name:  "test",
		Email: "test@example.com",
	})
	log.Println("Get user", userEntity)
	return &model.User{
		ID:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}, nil
}
