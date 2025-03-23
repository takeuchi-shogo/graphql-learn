package loaders

import (
	"context"
	"net/http"

	"github.com/takeuchi-shogo/graphql-learn/go/gql-gen/generated/model"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// userReader reads Users from a database
type userReader struct {
}

func NewUserReader() *userReader {
	return &userReader{}
}

// GetUser バッチローダー用のメソッド
func (r *userReader) GetUser(ctx context.Context, keys []string) ([]*model.User, []error) {
	users := make([]*model.User, len(keys))
	errors := make([]error, len(keys))

	// 実際にはDBから取得する処理を実装
	for i, key := range keys {
		users[i] = &model.User{
			ID:   key,
			Name: "User " + key,
		}
	}

	return users, errors
}

type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

func NewLoaders() *Loaders {
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(NewUserReader().GetUser),
	}
}

// Middleware injects data loaders into the context
func Middleware(next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders()
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

// GetUsers returns many users by ids efficiently
func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
