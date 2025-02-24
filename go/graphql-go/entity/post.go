package entity

type Post struct {
	id     string
	title  string
	userID string
}

func NewPost(id, title, userID string) *Post {
	return &Post{id: id, title: title, userID: userID}
}

func (p Post) ID() string {
	return p.id
}

func (p Post) Title() string {
	return p.title
}

func (p Post) UserID() string {
	return p.userID
}
