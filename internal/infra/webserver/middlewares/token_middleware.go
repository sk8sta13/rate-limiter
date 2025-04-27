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

type TokenMiddleware struct {
	Repository *repository.TokenRepository
	LimitToken *[]config.Token
	Token      string
}

func (t *TokenMiddleware) Execute(w http.ResponseWriter, r *http.Request) error {
	tokenrequest := dto.TokenRequest{
		IP:            getIP(r.RemoteAddr),
		Token:         t.Token,
		CurrentMoment: time.Now().Unix(),
	}

	uc := usecase.NewTokenUseCase(r.Context(), t.Repository)
	err := uc.Execute(&tokenrequest, t.LimitToken)

	if errors.Is(err, entity.ErrTokenUnauthorized) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}

	if errors.Is(err, entity.ErrTokenExceededMaxRequest) {
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return err
	}

	return nil
}
