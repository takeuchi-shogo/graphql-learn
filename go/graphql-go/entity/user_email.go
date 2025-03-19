package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserEmail struct {
	id        string
	userId    string
	email     string
	verified  bool
	isDefault bool
	createdAt time.Time
	updatedAt time.Time
}

func newUserEmail(id, userId, email string, verified bool, isDefault bool, createdAt, updatedAt time.Time) *UserEmail {
	return &UserEmail{id: id, userId: userId, email: email, verified: verified, isDefault: isDefault, createdAt: createdAt, updatedAt: updatedAt}
}

func NewUserEmail(userId, email string, verified bool, isDefault bool) *UserEmail {
	id := uuid.New().String()
	createdAt := time.Now()
	updatedAt := time.Now()
	return newUserEmail(id, userId, email, verified, isDefault, createdAt, updatedAt)
}

func ReconstructUserEmail(id, userId, email string, verified bool, isDefault bool, createdAt, updatedAt time.Time) *UserEmail {
	return newUserEmail(id, userId, email, verified, isDefault, createdAt, updatedAt)
}

func (u *UserEmail) Email() string {
	return u.email
}

func (u *UserEmail) Verified() bool {
	return u.verified
}
