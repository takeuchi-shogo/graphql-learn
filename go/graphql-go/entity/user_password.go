package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserPassword struct {
	id        string
	userId    string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func newUserPassword(id, userId, password string, createdAt, updatedAt time.Time) *UserPassword {
	return &UserPassword{id: id, userId: userId, password: password, createdAt: createdAt, updatedAt: updatedAt}
}

func NewUserPassword(userId, password string) *UserPassword {
	id := uuid.New().String()
	createdAt := time.Now()
	updatedAt := time.Now()
	return newUserPassword(id, userId, password, createdAt, updatedAt)
}

func ReconstructUserPassword(id, userId, password string, createdAt, updatedAt time.Time) *UserPassword {
	return newUserPassword(id, userId, password, createdAt, updatedAt)
}

func (u *UserPassword) Password() string {
	return u.password
}

func NewHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *UserPassword) Equals(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}
