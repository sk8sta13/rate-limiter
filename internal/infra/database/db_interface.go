package database

import (
	"context"

	"github.com/sk8sta13/rate-limiter/config"
)

type DBInterface interface {
	Connect(config *config.DB) DBInterface
	Set(ctx context.Context, key string, data []byte) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
