package routes

import (
	"api/controllers"
	"api/db"
	"api/middlewares"
	"api/repositories"

	"github.com/gin-gonic/gin"
)

func affectUsersRoutes(r *gin.Engine) {
	repository := repositories.UserRepository{Client: db.Client}
	controller := controllers.UserController{Repository: repository}

	authRoutes := r.Group("/auth")
	authRoutes.POST("/login", controller.Login)
	authRoutes.POST("/signup", controller.SignUp)

	userRoutes := r.Group("/users")
	userRoutes.Use(middlewares.Authenticate())
	userRoutes.GET("/self", controller.GetSelf)
	userRoutes.GET("/all", controller.GetUsers)

}
