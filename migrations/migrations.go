package migrations

import (
	"github.com/henrique-sulimann/golang-restapi/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Game{})
}
