package albums

import (
	"api/sse"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Stream *sse.Event

func AffectRoutes(r *gin.Engine) {
	group := r.Group("/albums")
	group.GET("", getAll)
	group.GET("/:id", getOne)
	group.POST("", createOne)
}

func getAll(c *gin.Context) {
	res := getAlbums()
	c.IndentedJSON(http.StatusOK, res)
}

func getOne(c *gin.Context) {
	id := c.Param("id")
	res, ok := getAlbumById(id)

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func createOne(c *gin.Context) {
	var dto CreateAlbumDto
	if err := c.BindJSON(&dto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	res, ok := postAlbums(dto)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
	} else {
		c.IndentedJSON(http.StatusCreated, res)
		sse.Stream.Message <- "OK"
	}
}
