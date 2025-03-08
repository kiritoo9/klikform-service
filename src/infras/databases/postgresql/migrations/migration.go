package migrations

import (
	"klikform/src/applications/models"
	"klikform/src/infras/databases/postgresql"
	"log"
)

func Migrations() {
	// load all data-models here from migrations/ directory
	// it's NOT AUTOMATIC, so you need to add each models
	if postgresql.DB == nil {
		log.Fatal("Database not initialized")
	}

	err := postgresql.DB.AutoMigrate(
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
