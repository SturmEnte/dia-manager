package config

import (
	"os"
	"strconv"
)

type Config struct {
    ServerPort string
    DatabaseUri string
    TokenLifetime int
    TokenSecret string
}

func Load() *Config {
    
    tokenLifetimeStr := getEnv("TOKEN_LIFETIME", "5")
    tokenLifetime, err := strconv.Atoi(tokenLifetimeStr)
    
    if err != nil {
        tokenLifetime = 5 // Default
    }

    return &Config{
        ServerPort: getEnv("PORT", "8369"),
        DatabaseUri: getEnv("DATABASE_URI", "postgres://test:test@localhost:5432/testdb?sslmode=disable"),
        TokenLifetime: tokenLifetime,
        TokenSecret: getEnv("TOKEN_SECRET", "1234"),
    }
}

func getEnv(key, defaultVal string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return defaultVal
}
