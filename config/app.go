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

func Init() {
	godotenv.Load(".env")
	if _, err := os.Stat(".env"); err == nil {
		log.Printf("Loading .env file")
		godotenv.Load(".env")
	}
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
}
