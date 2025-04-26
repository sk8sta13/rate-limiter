package middlewares

import (
	"net"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
)

type Middleware struct {
	RedisClient *redis.Client
	Limits      *config.Limits
}

func (m *Middleware) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//token := r.Header.Get("API_KEY")
			/*ip := getIP(r.RemoteAddr)
			key := fmt.Sprintf("%s%s", ip, token)

			data, _ := json.Marshal(map[string]interface{}{
				"Qtd":    12,
				"Moment": time.Now().Unix(),
			})
			ctx := context.Background()
			err := m.RedisClient.Set(ctx, key, data, 0).Err()
			if err != nil {
				panic(err)
			}*/

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
