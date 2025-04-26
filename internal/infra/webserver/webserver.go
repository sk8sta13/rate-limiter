package webserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/infra/webserver/handlers"
	"github.com/sk8sta13/rate-limiter/internal/infra/webserver/middlewares"
)

type HandlerProps struct {
	Method string
	Path   string
	Func   http.HandlerFunc
}

type WebServer struct {
	Router             chi.Router
	Handlers           []HandlerProps
	InternalMiddleware middlewares.Middleware
}

func NewWebServer(limits *config.Limits, redisCli *redis.Client) *WebServer {
	newWebServer := WebServer{
		Router:   chi.NewRouter(),
		Handlers: make([]HandlerProps, 0),
	}
	newWebServer.InternalMiddleware = middlewares.Middleware{
		RedisClient: redisCli,
		Limits:      limits,
	}

	newWebServer.AddHandler(http.MethodGet, "/", handlers.HelloWorld)

	return &newWebServer
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, HandlerProps{
		Method: method,
		Path:   path,
		Func:   handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(s.InternalMiddleware.RateLimiter)
	for _, h := range s.Handlers {
		s.Router.Method(h.Method, h.Path, h.Func)
	}

	if err := http.ListenAndServe("0.0.0.0:8080", s.Router); err != nil {
		log.Printf("Error starting the server.")
		return
	}
}
