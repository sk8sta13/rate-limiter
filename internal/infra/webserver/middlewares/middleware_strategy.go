package middlewares

import (
	"net/http"
)

type MiddlewareStrategy interface {
	Execute(w http.ResponseWriter, r *http.Request) error
}

func Factory(r *http.Request, m *Middleware) MiddlewareStrategy {
	token := r.Header.Get("API_KEY")

	if token != "" {
		return &TokenMiddleware{RedisClient: m.RedisClient, LimitToken: &m.Limits.Token, Token: token}
	}

	return &IPMiddleware{RedisClient: m.RedisClient, LimitIp: &m.Limits.Ip}
}
