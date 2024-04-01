package interfaces

import (
	"api/types"
	"context"

	"github.com/gin-gonic/gin"
)

type AlbumRepository interface {
	FindById(ctx context.Context, id string) (*types.Album, bool)
	FindMany(ctx context.Context) ([]types.Album, bool)
	Create(ctx context.Context, dto types.CreateAlbumDto) (types.Album, bool)
}

type AlbumController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Create(c *gin.Context)
}
