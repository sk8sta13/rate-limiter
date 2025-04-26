package database

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/internal/dto"
)

type IP struct {
	Db *redis.Client
}

func NewIPRedis(db *redis.Client) IP {
	return IP{Db: db}
}

func (IP *IP) GetData(ctx context.Context, key string, IPDB *dto.IPDB) {
	val, err := IP.Db.Get(ctx, key).Result()

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

func (IP *IP) SetData(ctx context.Context, IPDB *dto.IPDB) {
	data, _ := json.Marshal(map[string]interface{}{
		"Qtd":         IPDB.Qtd,
		"FirstMoment": IPDB.FirstMoment,
		"LastMoment":  IPDB.LastMoment,
	})

	err := IP.Db.Set(ctx, IPDB.Key, data, 0).Err()

	if err != nil {
		log.Println(err)
	}
}

func (IP *IP) DelData(ctx context.Context, key string) {
	err := IP.Db.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
	}
}
