package database

import (
	"context"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
)

type DB struct {
	strategy DBInterface
}

func (DB *DB) SetDB(db DBInterface, config *config.DB) {
	DB.strategy = db.Connect(config)
}

func (DB *DB) SetData(ctx context.Context, IPDB *dto.IPDB) {
	DB.strategy.SetData(ctx, IPDB)
}

func (DB *DB) GetData(ctx context.Context, key string, IPDB *dto.IPDB) {
	DB.strategy.GetData(ctx, key, IPDB)
}

func (DB *DB) DelData(ctx context.Context, key string) {
	DB.strategy.DelData(ctx, key)
}
