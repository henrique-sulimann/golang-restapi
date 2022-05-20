package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMongo struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name" validate:"required"`
	Password  string             `json:"password" validate:"required"`
	CreatedAt time.Time          `json:"CreatedAt"`
	UpdatedAt time.Time          `json:"UpdatedAt"`
	// DeletedAt   time.Time          `json:"DeletedAt"`
}
