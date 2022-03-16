package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"CreatedAt"`
	UpdatedAt   time.Time      `json:"UpdatedAt"`
	DeletedAt   gorm.DeletedAt `json:"DeletedAt"`
}
