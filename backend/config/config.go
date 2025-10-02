package config

import "os"

type Config struct {
    ServerPort string
    DatabaseUri string
}

func Load() *Config {
    return &Config{
        ServerPort: getEnv("PORT", "8369"),
        DatabaseUri: getEnv("DATABASE_URI", "postgres://test:test@localhost:5432/testdb?sslmode=disable"),
    }
}

func getEnv(key, defaultVal string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return defaultVal
}
