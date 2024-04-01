package repositories

import (
	"api/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumRepository struct {
	Client *mongo.Client
}

func (r AlbumRepository) FindById(ctx context.Context, id string) (re *types.Album, ok bool) {
	return &types.Album{}, false
}

func (r AlbumRepository) FindMany(ctx context.Context) (re []types.Album, ok bool) {
	var res []types.Album
	db := r.Client
	col := db.Database("app").Collection("albums")
	filter := bson.D{}
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return []types.Album{}, false
	}
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}
	return res, true
}

func (r AlbumRepository) Create(ctx context.Context, data types.CreateAlbumDto) (response types.Album, ok bool) {
	var newAlbum types.Album
	db := r.Client

	if data.Title == "" || data.Artist == "" {
		return newAlbum, false
	} else {
		col := db.Database("app").Collection("albums")
		newAlbum = types.Album{Title: data.Title, Artist: data.Artist, Price: data.Price}
		_, err := col.InsertOne(ctx, newAlbum)
		if err != nil {
			return newAlbum, false
		}
		return newAlbum, true
	}

}
