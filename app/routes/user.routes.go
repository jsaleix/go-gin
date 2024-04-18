package routes

import (
	"api/controllers"
	"api/db"
	"api/middlewares"
	"api/repositories"
	"api/services"

	"github.com/gin-gonic/gin"
)

func affectUsersRoutes(r *gin.Engine) {
	repository := repositories.UserRepository{Client: db.Client}
	service := services.UserService{Repository: repository}
	controller := controllers.UserController{Repository: repository, Service: service}

	authRoutes := r.Group("/auth")
	authRoutes.POST("/login", controller.Login)
	authRoutes.POST("/signup", controller.SignUp)

	userRoutes := r.Group("/users")
	userRoutes.Use(middlewares.Authenticate())
	userRoutes.GET("/self", controller.GetSelf)
	userRoutes.GET("/all", middlewares.IsAdmin(), controller.GetUsers)
	userRoutes.GET("/:id", middlewares.IsAdmin(), controller.GetUser)

}
