package controllers

import (
	"api/helpers"
	"api/interfaces"
	"api/models"
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
	var user models.User
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	validationErr := helpers.Validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	_, accountExists := ctrller.Repository.FindBy(ctx, map[string]string{"email": *user.Email})
	if accountExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "account already existing"})
		return
	}

	password := helpers.HashPassword(*user.Password)
	user.Password = &password
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	if userCreated := ctrller.Repository.Create(ctx, user); userCreated != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred, could not create user"})
	} else {
		c.Status(http.StatusCreated)
	}

}

func (ctrller UserController) Login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}

func (ctrller UserController) GetSelf(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}

func (ctrller UserController) GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}

func (ctrller UserController) GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, ok := ctrller.Repository.FindAll(ctx)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No album found"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}
