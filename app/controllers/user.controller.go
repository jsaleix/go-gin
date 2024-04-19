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
)

type UserController struct {
	Repository interfaces.UserRepositoryI
	Service    interfaces.UserServiceI
}

// SignUp
// @Summary SignUp
// @Produce json
// @Param body body types.SignUpDto true "Email and password"
// @Router /auth/signup [post]
// @Success 201
// @Accept json
// @Tags auth
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

	_, err := ctrller.Service.CreateUser(req, config.USER_ROLE)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusCreated)
	}

}

// Login
// @Summary Login
// @Produce json
// @Param body body types.LoginDto true "Email and password"
// @Router /auth/login [post]
// @Success 200 {object} types.LoginResponse
// @Accept json
// @Tags auth
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

	// I'm a bit paranoid
	if foundUser.Email == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	if passwordIsValid := helpers.VerifyPassword(*foundUser.Password, req.Password); !passwordIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	response, ok := ctrller.Service.LogUser(foundUser)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred, could not create user"})
		return
	}

	c.JSON(http.StatusOK, response)
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
		return
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
