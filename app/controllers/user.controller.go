package controllers

import (
	"api/config"
	"api/helpers"
	"api/interfaces"
	"api/models"
	"api/types"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	Repository interfaces.UserRepositoryI
}

func (ctrller UserController) SignUp(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var req types.SignUpDto
	defer cancel()

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	validationErr := helpers.Validate.Struct(req)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	_, accountExists := ctrller.Repository.FindBy(ctx, map[string]string{"email": req.Email})
	if accountExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred, could not create user"})
		return
	}

	password := helpers.HashPassword(req.Password)
	user := models.User{}
	user.Email = &req.Email
	user.Password = &password
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.User_type = new(string)
	*user.User_type = config.USER_ROLE
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.User_type, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	if validationErr := helpers.Validate.Struct(req); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if userCreated := ctrller.Repository.Create(ctx, user); !userCreated {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred, could not create user"})
	} else {
		c.Status(http.StatusCreated)
	}

}

func (ctrller UserController) Login(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var req types.LoginDto
	var foundUser models.User
	defer cancel()

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, ok := ctrller.Repository.FindBy(ctx, map[string]string{"email": req.Email})
	defer cancel()

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	if passwordIsValid := helpers.VerifyPassword(*foundUser.Password, req.Password); !passwordIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	if foundUser.Email == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.User_type, foundUser.User_id)
	foundUser.Token = &token
	foundUser.Refresh_token = &refreshToken
	foundUser.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if _, ok := ctrller.Repository.UpdateOne(ctx, foundUser.User_id, foundUser); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	var loginResponse types.LoginResponse
	loginResponse.ID = foundUser.ID
	loginResponse.User_id = foundUser.User_id
	loginResponse.Email = foundUser.Email
	loginResponse.Token = foundUser.Token
	loginResponse.User_type = foundUser.User_type
	loginResponse.Refresh_token = foundUser.Refresh_token
	loginResponse.Updated_at = foundUser.Updated_at

	c.JSON(http.StatusOK, loginResponse)
}

func (ctrller UserController) GetSelf(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userId := c.GetString("uid")

	if userId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return
	}
	user, ok := ctrller.Repository.FindById(ctx, userId)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No user found"})
		return
	}

	res := types.ConvertToPublicUser(&user)
	c.IndentedJSON(http.StatusOK, res)
}

func (ctrller UserController) GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return
	}
	user, ok := ctrller.Repository.FindById(ctx, id)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No user found"})
	}

	res := types.ConvertToPublicUser(&user)
	c.IndentedJSON(http.StatusOK, res)

}

func (ctrller UserController) GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, ok := ctrller.Repository.FindAll(ctx)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No user found"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}
