package config

import "os"

type Config struct {
	ServerPort string
	DBPath     string
	Env        string
}

func Load() (*Config, error) {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8090"),
		DBPath:     getEnv("DB_PATH", "./transactions.db"),
		Env:        getEnv("ENV", "development"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
