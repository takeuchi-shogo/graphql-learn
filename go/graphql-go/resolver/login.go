package resolver

import "context"

type LoginPayloadResolver struct {
	token string
}

func NewLoginPayloadResolver(ctx context.Context, token string) *LoginPayloadResolver {
	return &LoginPayloadResolver{token: token}
}

func (r *LoginPayloadResolver) Token() string {
	return r.token
}
