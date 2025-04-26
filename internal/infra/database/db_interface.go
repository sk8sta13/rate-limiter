package database

import (
	"context"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
)

type DBInterface interface {
	Connect(config *config.DB) DBInterface
	SetData(ctx context.Context, IPDB *dto.IPDB)
	GetData(ctx context.Context, key string, IPDB *dto.IPDB)
	DelData(ctx context.Context, key string)
}
