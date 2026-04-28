package config

import "os"

type Config struct {
    KafkaBootstrap string
    ServerPort     string
}

func Load() *Config {
    return &Config{
        KafkaBootstrap: getEnv("KAFKA_BOOTSTRAP", "localhost:29092"),
        ServerPort:     getEnv("SERVER_PORT", "8081"),
    }
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}