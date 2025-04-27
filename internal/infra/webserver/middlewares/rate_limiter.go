package middlewares

import (
	"net"
	"net/http"

	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/infra/database"
)

type Middleware struct {
	DB     *database.DB
	Limits *config.Limits
}

func (m *Middleware) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			strategy := Factory(r, m)
			if err := strategy.Execute(w, r); err != nil {
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}

func getIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return ""
	}

	return ip
}
