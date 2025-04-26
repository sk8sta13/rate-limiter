package middlewares

import (
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
)

type TokenMiddleware struct {
	RedisClient *redis.Client
	LimitToken  *[]config.Token
	Token       string
}

func (TokenMiddleware *TokenMiddleware) Execute(w http.ResponseWriter, r *http.Request) error {
	ip := getIP(r.RemoteAddr)

	println("Token")
	println(ip)

	return nil
}
