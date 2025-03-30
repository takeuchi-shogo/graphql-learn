package resolver

type UserErrorResolver struct {
	message string
	code    string
	path    []string
}

func NewUserErrorResolver(message string, code string, path []string) *UserErrorResolver {
	return &UserErrorResolver{message: message, code: code, path: path}
}

func (r *UserErrorResolver) Message() string {
	return r.message
}

func (r *UserErrorResolver) Code() string {
	return r.code
}

func (r *UserErrorResolver) Path() []string {
	return r.path
}
