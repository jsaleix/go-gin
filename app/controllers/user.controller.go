package controllers

import (
	"api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repository repositories.UserRepository
}

func (ctrller UserController) SignUp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
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
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}
