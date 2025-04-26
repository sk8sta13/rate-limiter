package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/entity"
	"github.com/sk8sta13/rate-limiter/internal/infra/database"
	"github.com/sk8sta13/rate-limiter/usecase"
)

type IPMiddleware struct {
	RedisClient *redis.Client
	LimitIp     *config.Ip
}

func (m *IPMiddleware) Execute(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	db := database.NewIPRedis(m.RedisClient)
	uc := usecase.NewIPUseCase(&ctx, m.LimitIp, &db)

	iprequest := dto.IPRequest{
		IP:            getIP(r.RemoteAddr),
		CurrentMoment: time.Now().Unix(),
	}
	err := uc.Execute(&iprequest)

	if errors.Is(err, entity.ErrIPExceededMaxRequests) {
		log.Printf("Error executing NewRegisterIPUseCase: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return err
	}

	return nil
}
