package interfaces

import (
	"api/models"
	"api/types"
	"context"

	"github.com/gin-gonic/gin"
)

type UserRepositoryI interface {
	FindById(context.Context, string) (models.User, bool)
	FindBy(context.Context, map[string]string) (models.User, bool)
	FindAll(context.Context) ([]models.User, bool)
	Create(context.Context, models.User) bool
	UpdateOne(context.Context, string, models.User) (models.User, bool)
}

type UserControllerI interface {
	SignUp(*gin.Context)
	Login(*gin.Context)
	GetSelf(*gin.Context)
	GetUser(*gin.Context)
	GetUsers(*gin.Context)
}

type UserServiceI interface {
	CreateUser(types.SignUpDto, string) (models.User, error)
	LogUser(models.User) (types.LoginResponse, bool)
}
