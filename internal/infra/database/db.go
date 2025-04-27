package database

import (
	"context"

	"github.com/sk8sta13/rate-limiter/config"
)

type DB struct {
	strategy DBInterface
}

func (DB *DB) SetDB(db DBInterface, config *config.DB) {
	DB.strategy = db.Connect(config)
}

func (DB *DB) Set(ctx context.Context, key string, data []byte) error {
	return DB.strategy.Set(ctx, key, data)
}

func (DB *DB) Get(ctx context.Context, key string) (string, error) {
	return DB.strategy.Get(ctx, key)
}

func (DB *DB) Del(ctx context.Context, key string) error {
	return DB.strategy.Del(ctx, key)
}
