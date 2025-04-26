package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DB struct {
	Host string
	Port int
	Pass string
}

type Redis struct {
	Host string
	Port int
}

type Ip struct {
	MaxRequests          int
	BloquedSeconds       int
	MaxRequestsInSeconds int
}

type Token struct {
	Token          string
	MaxRequests    int
	BloquedSeconds int
}

type Limits struct {
	Ip    Ip
	Token []Token
}

type Settings struct {
	Redis  Redis
	Limits Limits
}

func LoadSettings(s *Settings) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	ReadSettings(s)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		ReadSettings(s)
	})
}

func ReadSettings(s *Settings) {
	*s = Settings{
		Redis: Redis{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetInt("REDIS_PORT"),
		},
		Limits: Limits{
			Ip: Ip{
				MaxRequests:          viper.GetInt("IP_MAX_REQUESTS"),
				BloquedSeconds:       viper.GetInt("IP_BLOCKED_FOR_SECONDS"),
				MaxRequestsInSeconds: viper.GetInt("IP_MAX_REQUESTS_IN_SECONDS"),
			},
			Token: []Token{
				{
					Token:          viper.GetString("TONEK_P"),
					MaxRequests:    viper.GetInt("TOKEN_P_MAX_REQEUSTS_PER_SECOND"),
					BloquedSeconds: viper.GetInt("TOKEN_P_BLOCKED_SECONDS"),
				},
				{
					Token:          viper.GetString("TONEK_M"),
					MaxRequests:    viper.GetInt("TOKEN_M_MAX_REQUESTS_PER_SECOND"),
					BloquedSeconds: viper.GetInt("TOKEN_M_BLOCKED_SECONDS"),
				},
				{
					Token:          viper.GetString("TONEK_G"),
					MaxRequests:    viper.GetInt("TOKEN_G_MAX_REQUESTS_PER_SECOND"),
					BloquedSeconds: viper.GetInt("TOKEN_g_BLOCKED_SECONDS"),
				},
			},
		},
	}
}
