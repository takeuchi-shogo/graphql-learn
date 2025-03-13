package entity

import "time"

type Follow struct {
	id          string
	followerID  string
	followingID string
	createdAt   time.Time
}

func newFollow(id, followerID, followingID string, createdAt time.Time) *Follow {
	return &Follow{
		id:          id,
		followerID:  followerID,
		followingID: followingID,
		createdAt:   createdAt,
	}
}

func NewFollow(id, followerID, followingID string) *Follow {
	createdAt := time.Now()
	return newFollow(id, followerID, followingID, createdAt)
}

func (f Follow) ID() string {
	return f.id
}
