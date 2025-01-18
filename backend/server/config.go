package server

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DbConnectionString string `env:"DB_CONNECTION_STRING"`
}

func LoadCondiguration(ctx context.Context) Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	var cfg Config

	err = envconfig.Process(ctx, &cfg)
	if err != nil {
		log.Fatal("failed to read environment variables")
	}

	return cfg
}
