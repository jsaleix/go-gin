package repositories

import (
	"api/config"
	"api/models"
	"api/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumRepository struct {
	Client *mongo.Client
}

func (r AlbumRepository) FindById(ctx context.Context, id string) (re *models.Album, ok bool) {
	return &models.Album{}, false
}

func (r AlbumRepository) FindMany(ctx context.Context) (re []models.Album, ok bool) {
	var res []models.Album
	db := r.Client
	col := db.Database(config.DB_NAME).Collection("albums")
	filter := bson.D{}
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return []models.Album{}, false
	}
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}
	return res, true
}

func (r AlbumRepository) Create(ctx context.Context, data types.CreateAlbumDto) (response models.Album, ok bool) {
	var newAlbum models.Album
	db := r.Client

	if data.Title == "" || data.Artist == "" {
		return newAlbum, false
	} else {
		col := db.Database(config.DB_NAME).Collection("albums")
		newAlbum = models.Album{Title: data.Title, Artist: data.Artist, Price: data.Price}
		_, err := col.InsertOne(ctx, newAlbum)
		if err != nil {
			return newAlbum, false
		}
		return newAlbum, true
	}

}
