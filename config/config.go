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

type Ip struct {
	MaxRequests          int
	MaxRequestsInSeconds int
	BloquedSeconds       int
}

type Token struct {
	Token                string
	MaxRequests          int
	MaxRequestsInSeconds int
	BloquedSeconds       int
}

type Limits struct {
	Ip    Ip
	Token []Token
}

type Settings struct {
	DB     DB
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
		DB: DB{
			Host: viper.GetString("DB_HOST"),
			Port: viper.GetInt("DB_PORT"),
			Pass: viper.GetString("DB_PASS"),
		},
		Limits: Limits{
			Ip: Ip{
				MaxRequests:          viper.GetInt("IP_MAX_REQUESTS"),
				MaxRequestsInSeconds: viper.GetInt("IP_MAX_REQUESTS_IN_SECONDS"),
				BloquedSeconds:       viper.GetInt("IP_BLOCKED_FOR_SECONDS"),
			},
			Token: []Token{
				{
					Token:                viper.GetString("TONEK_P"),
					MaxRequests:          viper.GetInt("TONEK_P_MAX_REQUESTS"),
					MaxRequestsInSeconds: viper.GetInt("TOKEN_P_MAX_REQUESTS_IN_SECONDS"),
					BloquedSeconds:       viper.GetInt("TOKEN_P_BLOCKED_FOR_SECONDS"),
				},
				{
					Token:                viper.GetString("TONEK_M"),
					MaxRequests:          viper.GetInt("TONEK_M_MAX_REQUESTS"),
					MaxRequestsInSeconds: viper.GetInt("TOKEN_M_MAX_REQUESTS_IN_SECONDS"),
					BloquedSeconds:       viper.GetInt("TOKEN_M_BLOCKED_FOR_SECONDS"),
				},
				{
					Token:                viper.GetString("TONEK_G"),
					MaxRequests:          viper.GetInt("TONEK_G_MAX_REQUESTS"),
					MaxRequestsInSeconds: viper.GetInt("TOKEN_G_MAX_REQUESTS_IN_SECONDS"),
					BloquedSeconds:       viper.GetInt("TOKEN_G_BLOCKED_FOR_SECONDS"),
				},
			},
		},
	}
}
