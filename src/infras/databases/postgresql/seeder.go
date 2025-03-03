package postgresql

import (
	"klikform/src/infras/databases/postgresql/seeders"
	"log"

	"gorm.io/gorm"
)

func Seeders(db *gorm.DB) {
	// run all seeders
	seeders.RoleSeed(db)
	seeders.UserSeed(db)

	// write log
	log.Println("Seeder data is success")
}
