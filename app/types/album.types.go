package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID     primitive.ObjectID `json:"_id" bson: "_id"`
	Title  string             `json:"title" bson: "title"`
	Artist string             `json:"artist" bson: "artist"`
	Price  float64            `json:"price" bson: "price"`
}

type CreateAlbumDto struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
