package config

import "os"

type Config struct {
	AppName     string
	LogLevel    string
	ServerPort  string
	DatabaseURL string
	RedisURL    string
}

func Load() *Config {
	return &Config{
		AppName:     os.Getenv("APP_NAME"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
	}
}
