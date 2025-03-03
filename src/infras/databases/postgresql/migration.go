package postgresql

import (
	"klikform/src/applications/models"
	"log"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	// load all data-models here from migrations/ directory
	// it's NOT AUTOMATIC, so you need to add each models
	err := db.AutoMigrate(
		&models.Users{},
		&models.Roles{},
		&models.UserRoles{},
		&models.Workspaces{},
		&models.WorkspaceAttachments{},
		&models.WorkspaceUsers{},
		&models.Campaigns{},
		&models.CampaignSeos{},
		&models.CampaignForms{},
		&models.CampaignFormAttributes{},
		&models.FormAttributes{},
	)
	if err != nil {
		log.Fatal("Error when migrating database", err)
	}
	log.Println("Migration database is success")
}
