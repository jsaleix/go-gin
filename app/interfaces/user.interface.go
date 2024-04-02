package interfaces

import (
	"api/models"
	"context"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	FindById(context.Context, string) (models.User, bool)
	FindBy(context.Context, map[string]string) (models.User, bool)
	FindAll(context.Context) ([]models.User, bool)
	Create(context.Context, any) (models.User, bool)
}

type UserController interface {
	SignUp(*gin.Context)
	Login(*gin.Context)
	GetSelf(*gin.Context)
	GetUser(*gin.Context)
	GetUsers(*gin.Context)
}