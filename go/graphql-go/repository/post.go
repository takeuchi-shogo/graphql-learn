package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type PostRepository struct{}

type PostRepositoryImpl interface {
	GetPosts(first *int, after *string) ([]*entity.Post, error)
	GetPost(id string) (*entity.Post, error)
}

func NewPostRepository() PostRepositoryImpl {
	return &PostRepository{}
}

func (r *PostRepository) GetPosts(first *int, after *string) ([]*entity.Post, error) {
	data, err := os.ReadFile("mock/data.json")
	if err != nil {
		return nil, err
	}

	var mockData struct {
		Posts []struct {
			ID       string `json:"id"`
			Content  string `json:"content"`
			AuthorID string `json:"author_id"`
		} `json:"posts"`
	}

	if err := json.Unmarshal(data, &mockData); err != nil {
		return nil, err
	}

	var posts []*entity.Post

	fmt.Println("first", first)
	for i, mockPost := range mockData.Posts {
		if first != nil && i >= *first {
			break
		}
		posts = append(posts, entity.NewPost(mockPost.ID, mockPost.Content, mockPost.AuthorID))
	}

	return posts, nil
}

func (r *PostRepository) GetPost(id string) (*entity.Post, error) {
	data, err := os.ReadFile("mock/data.json")
	if err != nil {
		return nil, err
	}

	var mockData struct {
		Posts []struct {
			ID       string `json:"id"`
			Content  string `json:"content"`
			AuthorID string `json:"author_id"`
		} `json:"posts"`
	}

	if err := json.Unmarshal(data, &mockData); err != nil {
		return nil, err
	}

	var post *entity.Post

	for _, mockPost := range mockData.Posts {
		if mockPost.ID == id {
			post = entity.NewPost(mockPost.ID, mockPost.Content, mockPost.AuthorID)
		}
	}

	if post == nil {
		return nil, fmt.Errorf("post not found: %s", id)
	}

	return post, nil
}
