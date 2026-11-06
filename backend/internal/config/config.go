package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port                  string
	DatabaseURL           string
	DatabaseMaxConns      int32
	DatabaseMinConns      int32
	DatabaseMaxConnLife   time.Duration
	DatabaseMaxConnIdle   time.Duration
	DatabaseHealthTimeout time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		Port:                  getEnv("PORT", "8080"),
		DatabaseURL:           os.Getenv("DATABASE_URL"),
		DatabaseMaxConns:      int32(getEnvInt("DATABASE_MAX_CONNS", 10)),
		DatabaseMinConns:      int32(getEnvInt("DATABASE_MIN_CONNS", 0)),
		DatabaseMaxConnLife:   time.Duration(getEnvInt("DATABASE_MAX_CONN_LIFETIME_SECONDS", 1800)) * time.Second,
		DatabaseMaxConnIdle:   time.Duration(getEnvInt("DATABASE_MAX_CONN_IDLE_SECONDS", 300)) * time.Second,
		DatabaseHealthTimeout: time.Duration(getEnvInt("DATABASE_HEALTH_TIMEOUT_SECONDS", 3)) * time.Second,
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.DatabaseMaxConns < 1 {
		return Config{}, fmt.Errorf("DATABASE_MAX_CONNS must be at least 1")
	}
	if cfg.DatabaseMinConns < 0 || cfg.DatabaseMinConns > cfg.DatabaseMaxConns {
		return Config{}, fmt.Errorf("DATABASE_MIN_CONNS must be between 0 and DATABASE_MAX_CONNS")
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
