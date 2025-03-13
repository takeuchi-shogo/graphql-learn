package handler

type Query struct{}

func (q *Query) Hello() string {
	return "Hello, World!"
}
