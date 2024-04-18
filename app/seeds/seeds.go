package main

import (
	"api/config"
	"api/db"
	"api/repositories"
	"api/services"
	"api/types"
	"log"
)

type UserToCreate struct {
	Mail     string
	Role     string
	Password string
}

// var users = []string{"user@user.com", "admin@admin.com"}

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

func main() {
	log.Print("Starting...")
	config.Init()
	db.Init()

	repository := repositories.UserRepository{Client: db.Client}
	service := services.UserService{Repository: repository}
	for _, user := range users {
		log.Printf("Creating user with email: %s", user.Mail)
		dto := types.SignUpDto{Email: user.Mail, Password: user.Password}
		_, err := service.CreateUser(dto, user.Role)
		if err != nil {
			log.Printf("Error creating user: %s", err.Error())
		} else {
			log.Printf("User created")
		}
	}
	log.Print("Seed completed")

}
