package repositories

import (
	"api/config"
	"api/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Client *mongo.Client
}

func (r UserRepository) FindById(ctx context.Context, id string) (res models.User, ok bool) {
	return models.User{}, false
}

func (r UserRepository) FindBy(ctx context.Context, filters map[string]string) (res models.User, ok bool) {
	col := r.Client.Database(config.DB_NAME).Collection("user")
	err := col.FindOne(ctx, filters).Decode(&res)
	if err != nil {
		return res, false
	} else {
		return res, true
	}
}

func (r UserRepository) FindAll(ctx context.Context) (res []models.User, ok bool) {
	return []models.User{}, false
}

func (r UserRepository) Create(ctx context.Context, user models.User) bool {
	col := r.Client.Database(config.DB_NAME).Collection("user")
	_, errs := col.InsertOne(ctx, user)
	return errs != nil
}
