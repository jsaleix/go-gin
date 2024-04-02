package repositories

import (
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

func (r UserRepository) FindBy(_ context.Context, _ map[string]string) (res models.User, ok bool) {
	return models.User{}, false
}

func (r UserRepository) FindAll(ctx context.Context) (res []models.User, ok bool) {
	return []models.User{}, false
}

func (r UserRepository) Create(ctx context.Context, data any) (res models.User, ok bool) {
	return models.User{}, false
}
