package routes

import "github.com/gin-gonic/gin"

func AffectRoutes(r *gin.Engine) {
	affectAlbumsRoutes(r)
	affectUsersRoutes(r)
}
