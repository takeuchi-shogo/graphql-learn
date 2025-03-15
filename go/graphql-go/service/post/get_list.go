package post

import (
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
)

type GetListService struct {
	postRepository repository.PostRepositoryImpl
}

type GetListInput struct {
	First *int
	After *string
}

func NewGetListService(postRepository repository.PostRepositoryImpl) *GetListService {
	return &GetListService{postRepository: postRepository}
}

func (s *GetListService) GetList(input GetListInput) ([]*entity.Post, error) {
	posts, err := s.postRepository.GetPosts(input.First, input.After)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
