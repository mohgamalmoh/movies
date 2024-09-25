package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type API struct {
	HTTPPort int `env:"HTTP_PORT" envDefault:"8080"`
	DB       Database
}

func NewAPIConfig() (API, error) {
	c := API{}
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	err := env.Parse(&c)
	return c, err
}
