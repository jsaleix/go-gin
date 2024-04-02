package controllers

import (
	"api/interfaces"
	"api/types"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	Repository interfaces.AlbumRepositoryI
}

func (ctrller AlbumController) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, ok := ctrller.Repository.FindMany(ctx)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No album found"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func (ctrller AlbumController) GetOne(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, ok := ctrller.Repository.FindById(ctx, id)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func (ctrller AlbumController) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var dto types.CreateAlbumDto
	if err := c.BindJSON(&dto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	res, ok := ctrller.Repository.Create(ctx, dto)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
	} else {
		c.IndentedJSON(http.StatusCreated, res)
	}
}
