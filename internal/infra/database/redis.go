package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
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

func (db *Redis) GetData(ctx context.Context, key string, IPDB *dto.IPDB) {
	val, err := db.Client.Get(ctx, key).Result()

	if err != nil {
		log.Println(err)
	}

	if errors.Is(err, redis.Nil) {
		return
	}

	err = json.Unmarshal([]byte(val), &IPDB)

	if err != nil {
		log.Println(err)
	}
}

func (db *Redis) SetData(ctx context.Context, IPDB *dto.IPDB) {
	data, _ := json.Marshal(map[string]interface{}{
		"Qtd":         IPDB.Qtd,
		"FirstMoment": IPDB.FirstMoment,
		"LastMoment":  IPDB.LastMoment,
	})

	err := db.Client.Set(ctx, IPDB.Key, data, 0).Err()

	if err != nil {
		log.Println(err)
	}
}

func (db *Redis) DelData(ctx context.Context, key string) {
	err := db.Client.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
	}
}
