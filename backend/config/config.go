package config

import "os"

type Config struct {
    ServerPort string
}

func Load() *Config {
    return &Config{
        ServerPort: getEnv("PORT", "8369"),
    }
}

func getEnv(key, defaultVal string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return defaultVal
}
