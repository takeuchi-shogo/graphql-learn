package loader

import (
	"context"
	"fmt"
	"sync"

	"github.com/graph-gophers/dataloader"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/entity"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/service"
)

type userLoader struct{}

func newUserLoader() dataloader.BatchFunc {
	return userLoader{}.loadBatch
}

func (l userLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	userService, ok := ctx.Value(UserServiceKey).(*service.UserService)
	if !ok {
		err := fmt.Errorf("user service not found in context")
		results := make([]*dataloader.Result, len(keys))
		for i := range results {
			results[i] = &dataloader.Result{Error: err}
		}
		return results
	}
	fmt.Println("userService", userService)

	results := make([]*dataloader.Result, len(keys))
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, key := range keys {
		wg.Add(1)
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			user, err := userService.GetUser(key.String())
			fmt.Println("user", user)

			mu.Lock()
			results[i] = &dataloader.Result{Data: user, Error: err}
			mu.Unlock()
		}(i, key)
	}

	wg.Wait()
	return results
}

func LoadUser(ctx context.Context, id string) (*entity.User, error) {
	var user *entity.User

	loader, err := extract(ctx, userLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := loader.Load(ctx, dataloader.StringKey(id))()
	if err != nil {
		return nil, err
	}

	user, ok := data.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("invalid data")
	}

	return user, nil
}
