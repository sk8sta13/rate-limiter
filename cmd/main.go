package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/infra/webserver"
)

func main() {
	var settings config.Settings
	config.LoadSettings(&settings)

	webserver := webserver.NewWebServer(
		&settings.Limits,
		redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", settings.Redis.Host, settings.Redis.Port),
		}),
	)
	webserver.Start()

	/*redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", settings.Redis.Host, settings.Redis.Port),
	})
	x := database.NewIPRedis(redisClient)
	ipdb := dto.IPDB{Key: "172.27.0.1"}
	x.GetData(context.Background(), &ipdb)
	//println(ipdb)
	fmt.Println(ipdb)*/

	/*ctx := context.Background()
	data, _ := json.Marshal(map[string]interface{}{
		"Qtd":    12,
		"Moment": time.Now().Unix(),
	})
	err := redisClient.Set(ctx, "127.0.0.1", data, 0).Err()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)*/
}
