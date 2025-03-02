package postgresql

import (
	"klikform/src/infras/databases/postgresql/migrations"
	"log"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	// load all data-models here from migrations/ directory
	// it's NOT AUTOMATIC, so you need to add each models
	err := db.AutoMigrate(
		&migrations.Users{},
	)
	if err != nil {
		log.Fatal("Error when migrating database", err)
	}
	log.Println("Migration database is success")
}
