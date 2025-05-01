package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port         string
	DatabasePath string
	JWTSecret    string
)

func LoadConfig(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load env file at %s: %v", filePath, err))
	}

	Port = os.Getenv("PORT")
	DatabasePath = os.Getenv("DATABASE_PATH")
	JWTSecret = os.Getenv("JWT_SECRET")
}
