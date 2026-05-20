package config

import "os"

type Config struct {
	ServerPort  string
	DatabaseURL string
	Env         string
}

func Load() (*Config, error) {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8090"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://finance:finance@localhost:5432/finance?sslmode=disable"),
		Env:         getEnv("ENV", "development"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
