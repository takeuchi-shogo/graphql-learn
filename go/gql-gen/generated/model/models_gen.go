// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
	GetID() string
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token      string       `json:"token"`
	Entity     *User        `json:"entity"`
	UserErrors []*UserError `json:"userErrors"`
}

type Mutation struct {
}

type NewTodo struct {
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      TodoStatus `json:"status"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type Query struct {
}

type Todo struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      TodoStatus `json:"status"`
	User        *User      `json:"user"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
}

func (Todo) IsNode()            {}
func (this Todo) GetID() string { return this.ID }

type TodoConnection struct {
	Edges      []*TodoEdge  `json:"edges"`
	Nodes      []*Todo      `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	UserErrors []*UserError `json:"userErrors"`
}

type TodoCreatePayload struct {
	Entity     *Todo        `json:"entity"`
	UserErrors []*UserError `json:"userErrors"`
}

type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

type User struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type UserCreateInput struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type UserCreatePayload struct {
	Entity     *User        `json:"entity"`
	UserErrors []*UserError `json:"userErrors"`
}

type UserError struct {
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Path    []string `json:"path"`
}

type TodoStatus string

const (
	TodoStatusPending    TodoStatus = "PENDING"
	TodoStatusTodo       TodoStatus = "TODO"
	TodoStatusInProgress TodoStatus = "IN_PROGRESS"
	TodoStatusCompleted  TodoStatus = "COMPLETED"
)

var AllTodoStatus = []TodoStatus{
	TodoStatusPending,
	TodoStatusTodo,
	TodoStatusInProgress,
	TodoStatusCompleted,
}

func (e TodoStatus) IsValid() bool {
	switch e {
	case TodoStatusPending, TodoStatusTodo, TodoStatusInProgress, TodoStatusCompleted:
		return true
	}
	return false
}

func (e TodoStatus) String() string {
	return string(e)
}

func (e *TodoStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoStatus", str)
	}
	return nil
}

func (e TodoStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleUser  UserRole = "USER"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleUser,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleUser:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
