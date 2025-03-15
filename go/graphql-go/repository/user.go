package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type UserRepository struct{}

type UserRepositoryImpl interface {
	GetUser(id string) (*entity.User, error)
}

func NewUserRepository() UserRepositoryImpl {
	return &UserRepository{}
}

func (r *UserRepository) GetUser(id string) (*entity.User, error) {
	// JSONデータから該当するユーザーを検索
	data, err := os.ReadFile("mock/data.json")
	if err != nil {
		return nil, err
	}

	var mockData struct {
		Users []struct {
			ID          string `json:"id"`
			UserName    string `json:"user_name"`
			DisplayName string `json:"display_name"`
			Bio         string `json:"bio"`
			AvatarURL   string `json:"avatar_url"`
		} `json:"users"`
	}

	if err := json.Unmarshal(data, &mockData); err != nil {
		return nil, err
	}

	var user *entity.User

	// IDに一致するユーザーを探す
	for _, mockUser := range mockData.Users {
		if mockUser.ID == id {
			user = entity.NewUser(
				mockUser.ID,
				mockUser.UserName,
				mockUser.DisplayName,
				mockUser.Bio,
				mockUser.AvatarURL,
			)
			break
		}
	}

	if user == nil {
		return nil, fmt.Errorf("user not found: %s", id)
	}

	return user, nil
}
