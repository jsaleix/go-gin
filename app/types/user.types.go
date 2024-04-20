package types

import (
	"api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpDto struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	ID            primitive.ObjectID `bson:"_id"`
	Email         *string            `json:"email" validate:"email,required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}

type UserPublic struct {
	ID         primitive.ObjectID `bson:"_id"`
	Email      *string            `json:"email" validate:"email,required"`
	User_type  *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	User_id    string             `json:"user_id"`
}

func ConvertToPublicUser(user *models.User) *UserPublic {
	return &UserPublic{
		ID:         user.ID,
		Email:      user.Email,
		User_type:  user.User_type,
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
		User_id:    user.User_id,
	}
}

type UpdateUserDto struct {
	Email *string `json:"email" validate:"email,required"`
}

type UpdatePasswordDto struct {
	Password string `json:"password" validate:"required,min=6"`
}
