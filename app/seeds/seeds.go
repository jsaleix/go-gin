package main

import (
	"api/config"
	"api/db"
	"log"
)

var users = []string{"user@user.com", "admin@admin.com"}

func main() {
	log.Print("Starting...")
	config.Init()
	db.Init()

	for _, mail := range users {
		log.Printf("Creating user with email: %s", mail)
	}

}
