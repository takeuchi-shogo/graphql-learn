package post

import (
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
)

type CreateService struct {
	postRepository repository.PostRepositoryImpl
}

type CreateInput struct {
	Content  string
	AuthorID string
}

func NewCreateService(postRepository repository.PostRepositoryImpl) *CreateService {
	return &CreateService{
		postRepository: postRepository,
	}
}

func (s *CreateService) Create(input CreateInput) (*entity.Post, error) {
	post := entity.NewPost("", input.Content, input.AuthorID)
	// TODO: 実際のDBへの保存処理を実装
	return post, nil
}
