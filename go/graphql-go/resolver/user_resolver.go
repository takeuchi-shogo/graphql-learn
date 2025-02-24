package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type UserResolver struct {
	user *entity.User
}

func NewUserResolver(user *entity.User) *UserResolver {
	return &UserResolver{user: user}
}

func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID())
}

func (r *UserResolver) Name() string {
	return r.user.Name()
}
