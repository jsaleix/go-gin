package albums

import (
	"net/http"
	"sse/sse"

	"github.com/gin-gonic/gin"
)

var Stream *sse.Event

func AffectRoutes(r *gin.Engine, s *sse.Event) {
	Stream = s
	r.GET("/albums", getAll)
	r.GET("/albums/:id", getOne)
	r.POST("/albums", createOne)
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
		c.IndentedJSON(http.StatusBadRequest, nil)
	}
	res, ok := postAlbums(dto)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, nil)
	} else {
		c.IndentedJSON(http.StatusCreated, res)
		Stream.Message <- "OK"
	}
}
