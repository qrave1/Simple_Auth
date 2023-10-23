package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	RdbAddress string
	RdbPort    string
	ServerPort string
	Secret     []byte
}

func NewConfig() *Config {
	return &Config{
		RdbAddress: os.Getenv("REDIS_ADDR"),
		RdbPort:    os.Getenv("REDIS_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
		Secret:     []byte(os.Getenv("SECRET_KEY")),
	}
}
