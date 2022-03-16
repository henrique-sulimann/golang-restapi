package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameMongo struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"CreatedAt"`
	UpdatedAt   time.Time          `json:"UpdatedAt"`
	// DeletedAt   time.Time          `json:"DeletedAt"`
}
