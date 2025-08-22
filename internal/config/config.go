package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	DBUrl               string
	JWTSecret           string
	JWTExpires          time.Duration
	RefreshTokenExpires time.Duration
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
	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) < 32 {
		panic("JWT_SECRET must be at least 32 characters long")
	}
	jwtExpires, err := time.ParseDuration(os.Getenv("JWT_EXPIRES"))
	if err != nil {
		jwtExpires = 24 * time.Hour
	}
	refreshTokenExpires, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRES"))
	if err != nil {
		refreshTokenExpires = 168 * time.Hour
	}
	return Config{
		Port:                port,
		DBUrl:               dbUrl,
		JWTSecret:           jwtSecret,
		JWTExpires:          jwtExpires,
		RefreshTokenExpires: refreshTokenExpires,
	}
}
