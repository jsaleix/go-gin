package repositories

import (
	"context"

	"github.com/golangcompany/JWT-Authentication/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Client *mongo.Client
}

func (r UserRepository) FindById(ctx context.Context, id string) (res models.User, ok bool) {
	return models.User{}, false
}

func (r UserRepository) findBy(_ context.Context, _ map[string]string) (res models.User, ok bool) {
	return models.User{}, false
}

func (r UserRepository) FindMany(ctx context.Context) (res []models.User, ok bool) {
	return []models.User{}, false
}

func (r UserRepository) Create(ctx context.Context, data any) (res models.User, ok bool) {
	return models.User{}, false
}
