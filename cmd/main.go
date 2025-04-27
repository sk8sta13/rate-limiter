package main

import (
	"github.com/sk8sta13/rate-limiter/config"
	"github.com/sk8sta13/rate-limiter/internal/infra/database"
	"github.com/sk8sta13/rate-limiter/internal/infra/webserver"
)

func main() {
	var settings config.Settings
	config.LoadSettings(&settings)

	db := &database.DB{}
	db.SetDB(&database.Redis{}, &settings.DB)

	webserver := webserver.NewWebServer(&settings.Limits, db)
	webserver.Start()
}
