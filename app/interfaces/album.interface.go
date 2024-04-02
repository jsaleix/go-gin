package interfaces

import (
	"api/models"
	"api/types"
	"context"

	"github.com/gin-gonic/gin"
)

type AlbumRepositoryI interface {
	FindById(ctx context.Context, id string) (*models.Album, bool)
	FindMany(ctx context.Context) ([]models.Album, bool)
	Create(ctx context.Context, dto types.CreateAlbumDto) (models.Album, bool)
}

type AlbumControllerI interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Create(c *gin.Context)
}
