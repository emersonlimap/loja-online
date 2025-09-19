package config

import (
	"os"
)

type Config struct {
	DatabaseURL   string
	JWTSecret     string
	Environment   string
	Port          string
	SessionSecret string
}

func Load() *Config {
	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://admin:admin@localhost:5432/loja_online?sslmode=disable"),
		JWTSecret:     getEnv("JWT_SECRET", "minha_chave_secreta_jwt_para_desenvolvimento"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		Port:          getEnv("PORT", "8080"),
		SessionSecret: getEnv("SESSION_SECRET", "minha_chave_secreta_de_sessao"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
