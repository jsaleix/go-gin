package routes

import (
	"api/albums"

	"github.com/gin-gonic/gin"
)

func AffectRoutes(r *gin.Engine) {
	albums.AffectRoutes(r)
}
