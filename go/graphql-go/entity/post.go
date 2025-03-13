package entity

import "time"

type Post struct {
	id        string
	content   string
	authorID  string
	createdAt time.Time
	updatedAt time.Time
}

func newPost(id, content, authorID string, createdAt, updatedAt time.Time) *Post {
	return &Post{
		id:        id,
		content:   content,
		authorID:  authorID,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewPost(id, content, authorID string) *Post {
	createdAt := time.Now()
	updatedAt := time.Now()
	return newPost(id, content, authorID, createdAt, updatedAt)
}

func (p Post) ID() string {
	return p.id
}

func (p Post) Content() string {
	return p.content
}

func (p Post) AuthorID() string {
	return p.authorID
}
