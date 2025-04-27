package middlewares

import (
	"net/http"

	"github.com/sk8sta13/rate-limiter/internal/infra/database/repository"
)

type MiddlewareStrategy interface {
	Execute(w http.ResponseWriter, r *http.Request) error
}

func Factory(r *http.Request, m *Middleware) MiddlewareStrategy {
	token := r.Header.Get("API_KEY")

	if token != "" {
		repository := repository.NewTokenRepository(*m.DB)
		return &TokenMiddleware{Repository: repository, LimitToken: &m.Limits.Token, Token: token}
	}

	repository := repository.NewIpRepository(*m.DB)
	return &IPMiddleware{Repository: repository, LimitIp: &m.Limits.Ip}
}
