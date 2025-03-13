package entity

import "time"

type Like struct {
	id        string
	postID    string
	userID    string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func newLike(id, postID, userID string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Like {
	return &Like{
		id:        id,
		postID:    postID,
		userID:    userID,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

func NewLike(id, postID, userID string) *Like {
	createdAt := time.Now()
	updatedAt := time.Now()
	return newLike(id, postID, userID, createdAt, updatedAt, nil)
}

func ReconstructLike(id, postID, userID string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Like {
	return newLike(id, postID, userID, createdAt, updatedAt, deletedAt)
}

func (l Like) ID() string {
	return l.id
}

func (l Like) PostID() string {
	return l.postID
}

func (l Like) UserID() string {
	return l.userID
}

func (l Like) CreatedAt() time.Time {
	return l.createdAt
}
