package services

import (
	"api/config"
	"api/helpers"
	"api/interfaces"
	"api/models"
	"api/types"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repository interfaces.UserRepositoryI
}

func (service UserService) CreateUser(createUserDto types.SignUpDto, role string) (models.User, error) {
	if role == "" {
		role = config.USER_ROLE
	}
	hashedPassword := helpers.HashPassword(createUserDto.Password)
	user := models.User{}
	user.Email = new(string)
	*user.Email = createUserDto.Email
	user.Password = &hashedPassword
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.User_type = new(string)
	*user.User_type = role

	if validationErr := helpers.Validate.Struct(user); validationErr != nil {
		return user, validationErr
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	userCreated := service.Repository.Create(ctx, user)
	if !userCreated {
		return user, errors.New("an error occurred, could not create user")
	}
	return user, nil
}

func (service UserService) LogUser(user models.User) (res types.LoginResponse, ok bool) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.User_type, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if _, ok := service.Repository.UpdateOne(ctx, user.User_id, user); !ok {
		return res, false
	}

	var loginResponse types.LoginResponse
	loginResponse.ID = user.ID
	loginResponse.User_id = user.User_id
	loginResponse.Email = user.Email
	loginResponse.Token = user.Token
	loginResponse.User_type = user.User_type
	loginResponse.Refresh_token = user.Refresh_token
	loginResponse.Updated_at = user.Updated_at

	return loginResponse, true
}
