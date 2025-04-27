package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/infra/database"
)

type IpRepository struct {
	DB database.DB
}

func NewIpRepository(db database.DB) *IpRepository {
	return &IpRepository{
		DB: db,
	}
}

func (i *IpRepository) GetData(ctx context.Context, key string, IPDB *dto.IPDB) {
	data, err := i.DB.Get(ctx, key)

	if err != nil {
		log.Println(err)
		return
	}

	if data == "" {
		return
	}

	err = json.Unmarshal([]byte(data), &IPDB)

	if err != nil {
		log.Println(err)
	}
}

func (i *IpRepository) SetData(ctx context.Context, IPDB *dto.IPDB) {
	data, _ := json.Marshal(map[string]interface{}{
		"Qtd":         IPDB.Qtd,
		"FirstMoment": IPDB.FirstMoment,
		"LastMoment":  IPDB.LastMoment,
	})

	err := i.DB.Set(ctx, IPDB.Key, data)

	if err != nil {
		log.Println(err)
	}
}

func (i *IpRepository) DelData(ctx context.Context, key string) {
	err := i.DB.Del(ctx, key)

	if err != nil {
		log.Println(err)
	}
}
