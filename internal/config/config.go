package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL env variable not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variable not set")
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
	}
}
