package entity

import "time"

type User struct {
	id          string
	userName    string
	displayName string
	bio         string
	avatarURL   string
	createdAt   time.Time
	updatedAt   time.Time
}

func newUser(id, userName, displayName, bio, avatarURL string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:          id,
		userName:    userName,
		displayName: displayName,
		bio:         bio,
		avatarURL:   avatarURL,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}
func NewUser(id, userName, displayName, bio, avatarURL string) *User {
	createdAt := time.Now()
	updatedAt := time.Now()
	return newUser(id, userName, displayName, bio, avatarURL, createdAt, updatedAt)
}

func ReconstructUser(id, userName, displayName, bio, avatarURL string, createdAt, updatedAt time.Time) *User {
	return newUser(id, userName, displayName, bio, avatarURL, createdAt, updatedAt)
}

func (u User) ID() string {
	return u.id
}

func (u User) UserName() string {
	return u.userName
}

func (u User) DisplayName() string {
	return u.displayName
}

func (u User) Bio() string {
	return u.bio
}

func (u User) AvatarURL() string {
	return u.avatarURL
}
