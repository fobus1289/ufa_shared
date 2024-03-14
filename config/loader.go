package config

import (
	"github.com/caarlos0/env/v6"

	_ "github.com/joho/godotenv/autoload"
)

func Load[T any](instance *T) *T {
	if err := env.Parse(instance); err != nil {
		panic(err)
	}
	return instance
}
