package loader

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/dataloader"
)

type LoaderCollection struct {
	dataloaderFuncMap map[key]dataloader.BatchFunc
}

func NewLoaderCollection() LoaderCollection {
	return LoaderCollection{
		dataloaderFuncMap: map[key]dataloader.BatchFunc{
			userLoaderKey: newUserLoader(),
		},
	}
}

func (c LoaderCollection) Attach(ctx context.Context) context.Context {
	for k, f := range c.dataloaderFuncMap {
		newLoader := dataloader.NewBatchedLoader(f)
		ctx = context.WithValue(ctx, k, newLoader)
	}
	return ctx
}

func extract(ctx context.Context, k key) (*dataloader.Loader, error) {
	// コンテキストの中身をデバッグ
	if ctx == nil {
		log.Printf("Context is nil")
		return nil, fmt.Errorf("nil context")
	}

	loader, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		log.Printf("Failed to extract loader for key: %v, value type: %T", k, ctx.Value(k))
		return nil, fmt.Errorf("loader not found for key: %v", k)
	}

	log.Printf("Successfully extracted loader for key: %v", k)
	return loader, nil
}
