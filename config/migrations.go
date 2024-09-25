package config

import (
	"github.com/caarlos0/env/v6"
)

func NewMigrationsConfig() (Database, error) {
	c := Database{}
	err := env.Parse(&c)
	return c, err
}
