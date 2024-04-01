package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEFAULT_PORT = "8080"
)

var PORT string = DEFAULT_PORT
var DB_URL string

func Init() {
	godotenv.Load(".env")
	if _, err := os.Stat(".env"); err == nil {
		log.Printf("Loading .env file")
		godotenv.Load(".env")
	}
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DB_URL") != "" {
		DB_URL = os.Getenv("DB_URL")
	}
}
