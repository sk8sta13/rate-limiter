package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/infra/database"
)

type TokenRepository struct {
	DB database.DB
}

func NewTokenRepository(db database.DB) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

func (i *TokenRepository) GetData(ctx context.Context, key string, TokenDB *dto.TokenDB) {
	data, err := i.DB.Get(ctx, key)

	if err != nil {
		log.Println(err)
		return
	}

	if data == "" {
		return
	}

	err = json.Unmarshal([]byte(data), &TokenDB)

	if err != nil {
		log.Println(err)
	}
}

func (i *TokenRepository) SetData(ctx context.Context, TokenDB *dto.TokenDB) {
	data, _ := json.Marshal(map[string]interface{}{
		"Qtd":         TokenDB.Qtd,
		"FirstMoment": TokenDB.FirstMoment,
		"LastMoment":  TokenDB.LastMoment,
	})

	err := i.DB.Set(ctx, TokenDB.Key, data)

	if err != nil {
		log.Println(err)
	}
}

func (i *TokenRepository) DelData(ctx context.Context, key string) {
	err := i.DB.Del(ctx, key)

	if err != nil {
		log.Println(err)
	}
}
