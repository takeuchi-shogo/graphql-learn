package resolver

import (
	"context"

	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
)

type LoginPayloadResolver struct {
	token string
	user  *entity.User
}

func NewLoginPayloadResolver(ctx context.Context, token string, user *entity.User) *LoginPayloadResolver {
	return &LoginPayloadResolver{
		token: token,
		user:  user,
	}
}

func (r *LoginPayloadResolver) Token() *string {
	return &r.token
}

func (r *LoginPayloadResolver) Entity() *UserResolver {
	return NewUserResolver(r.user)
}

func (r *LoginPayloadResolver) UserErrors() []*UserErrorResolver {
	return []*UserErrorResolver{
		NewUserErrorResolver("Invalid email or password", "INVALID_CREDENTIALS", []string{"email"}),
	}
}
