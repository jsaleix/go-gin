package routes

import (
	"api/albums"
	"api/sse"

	"github.com/gin-gonic/gin"
)

func AffectRoutes(r *gin.Engine, s *sse.Event) {
	albums.AffectRoutes(r, s)
}
