package seeders

import (
	"klikform/src/applications/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RoleSeed(db *gorm.DB) {
	// check data already seeded
	var count int64
	db.Model(&models.Roles{}).Count(&count)
	if count > 0 {
		return // stop process so seeder won't run
	}

	roles := []models.Roles{
		{ID: uuid.New(), Name: "admin", Descriptions: "Role for admin"},
		{ID: uuid.New(), Name: "user", Descriptions: "Role for user"},
	}
	db.Create(&roles)
}
