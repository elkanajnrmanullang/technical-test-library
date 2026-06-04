package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	Secret string
	Expires int
	RefreshSecret string
	RefreshExpires int
}

func GetJwtConfig() JwtConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtExpires, err := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	if err != nil {
		log.Fatal("Error parsing JWT_EXPIRES from .env")
	}

	jwtRefreshExpires, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRES"))
	if err != nil {
		log.Fatal("Error parsing JWT_REFRESH_EXPIRES from .env")
	}

	return JwtConfig{
		Secret: os.Getenv("JWT_SECRET"),
		Expires: jwtExpires,
		RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		RefreshExpires: jwtRefreshExpires,
	}
}

func (j *JwtConfig) GetJwtExpirationTime() time.Duration {
	return time.Duration(j.Expires) * time.Minute
}

func (j *JwtConfig) GetJwtRefreshExpirationTime() time.Duration {
	return time.Duration(j.RefreshExpires) * time.Minute
}