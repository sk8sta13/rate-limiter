package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
)

type Redis struct {
	Client *redis.Client
}

func (db *Redis) Connect(config *config.DB) DBInterface {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		}),
	}
}

func (db *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := db.Client.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func (db *Redis) Set(ctx context.Context, key string, data []byte) error {
	return db.Client.Set(ctx, key, data, 0).Err()
}

func (db *Redis) Del(ctx context.Context, key string) error {
	return db.Client.Del(ctx, key).Err()
}
