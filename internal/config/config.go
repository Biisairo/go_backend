package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Port         string
	DatabasePath string
	JWTSecret    string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Fail to load .env")
	}

	Port = os.Getenv("PORT")
	DatabasePath = os.Getenv("DATABASE_PATH")
	JWTSecret = os.Getenv("JWT_SECRET")
}
