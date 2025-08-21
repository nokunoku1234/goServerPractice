package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() Config {
	_ = godotenv.Load() // .envが無くてもOK
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	return Config{Port: port}
}
