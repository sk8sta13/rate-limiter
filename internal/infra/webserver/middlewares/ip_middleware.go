package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/dto"
	"github.com/sk8sta13/rate-limiter/internal/entity"
	"github.com/sk8sta13/rate-limiter/internal/infra/database/repository"
	"github.com/sk8sta13/rate-limiter/usecase"
)

type IPMiddleware struct {
	Repository *repository.IpRepository
	LimitIp    *config.Ip
}

func (m *IPMiddleware) Execute(w http.ResponseWriter, r *http.Request) error {
	iprequest := dto.IPRequest{
		IP:            getIP(r.RemoteAddr),
		CurrentMoment: time.Now().Unix(),
	}

	uc := usecase.NewIPUseCase(r.Context(), m.LimitIp, m.Repository)
	err := uc.Execute(&iprequest)

	if errors.Is(err, entity.ErrIPExceededMaxRequests) {
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return err
	}

	return nil
}
