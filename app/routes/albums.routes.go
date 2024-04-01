package routes

import (
	"api/controllers"
	"api/db"
	"api/repositories"

	"github.com/gin-gonic/gin"
)

func affectAlbumsRoutes(r *gin.Engine) {
	repository := repositories.AlbumRepository{Client: db.Client}
	controller := controllers.AlbumController{Repository: repository}

	group := r.Group("/albums")
	group.GET("", controller.GetAll)
	group.GET("/:id", controller.GetOne)
	group.POST("", controller.Create)
}
