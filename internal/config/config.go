package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort        string
	DatabaseURL       string
	Env               string
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime int
}

func Load() (*Config, error) {
	return &Config{
		ServerPort:        getEnv("SERVER_PORT", "8090"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://finance:finance@localhost:5432/finance?sslmode=disable"),
		Env:               getEnv("ENV", "development"),
		DBMaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLifetime: getEnvAsInt("DB_CONN_MAX_LIFETIME_MIN", 5),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return fallback
}
