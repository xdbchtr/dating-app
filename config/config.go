package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
}

func LoadConfig() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}
