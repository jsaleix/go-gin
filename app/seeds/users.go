package main

import (
	"api/config"
	"api/repositories"
	"api/services"
	"api/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserToCreate struct {
	Mail     string
	Role     string
	Password string
}

var users = []UserToCreate{
	{
		Mail:     "user@user.com",
		Role:     config.USER_ROLE,
		Password: "User123",
	},
	{
		Mail:     "admin@admin.com",
		Role:     config.ADMIN_ROLE,
		Password: "Admin123",
	},
}

func SeedUsers(db *mongo.Client) {
	log.Print("Starting seeding users...")

	repository := repositories.UserRepository{Client: db}
	service := services.UserService{Repository: repository}
	for _, user := range users {
		log.Printf("Creating user with email: %s", user.Mail)
		dto := types.SignUpDto{Email: user.Mail, Password: user.Password}
		_, err := service.CreateUser(dto, user.Role)
		if err != nil {
			log.Printf("X Error creating user: %s", err.Error())
		} else {
			log.Printf("User created")
		}
	}
	log.Print("Done seeding users")

}
