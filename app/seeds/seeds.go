package main

import (
	"api/config"
	"api/db"
	"log"
)

func main() {
	log.Print("Starting...")
	config.Init()
	db.Init()

	SeedUsers(db.Client)

	log.Print("Seed completed")

}
