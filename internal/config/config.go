package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() Config {
	_ = godotenv.Load() // .envが無くてもOK
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DB NOT FOUND")
	}
	return Config{Port: port, DBUrl: dbUrl}
}
