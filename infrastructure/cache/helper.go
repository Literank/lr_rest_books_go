package cache

import "context"

// Helper reads and writes from cache.
type Helper interface {
	Save(ctx context.Context, key, value string) error
	Load(ctx context.Context, key string) (string, error)
}
