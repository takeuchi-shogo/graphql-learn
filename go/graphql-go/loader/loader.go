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
		log.Printf("Attaching loader for key: %v", k)
		newLoader := dataloader.NewBatchedLoader(f)
		ctx = context.WithValue(ctx, k, newLoader)

		// 確認用ログ
		if ctx.Value(k) != nil {
			log.Printf("Successfully attached loader for key: %v", k)
		}
		// ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(f))
	}
	return ctx
}

func extract(ctx context.Context, k key) (*dataloader.Loader, error) {
	// loader, ok := ctx.Value(k).(*dataloader.Loader)
	// if !ok {
	// 	log.Printf("Loader not found for key: %v", k)
	// 	return nil, fmt.Errorf("loader not found for key: %v", k)
	// }
	// return loader, nil
	log.Printf("Extracting loader for key: %v", k)

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
