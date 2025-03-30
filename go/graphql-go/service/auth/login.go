package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginService struct {
	userRepository repository.UserRepositoryImpl
}

func NewLoginService(userRepository repository.UserRepositoryImpl) *LoginService {
	return &LoginService{userRepository: userRepository}
}

func (s *LoginService) Login(input LoginInput) (string, *entity.User, error) {
	user, err := s.userRepository.GetUserByEmail(input.Email)
	if err != nil {
		return "", nil, err
	}

	if !user.Password.Equals(input.Password) {
		return "", nil, errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
