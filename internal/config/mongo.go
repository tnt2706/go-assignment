package config

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type MongoConfig struct {
	DbCommonConnectString string `env:"DB_COMMON_CONNECT_STRING" required:"true"`
}

func GetMongoConfig() *MongoConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Mongo unable to load .env file: %e", err)
	}

	cfg := MongoConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	return &cfg
}