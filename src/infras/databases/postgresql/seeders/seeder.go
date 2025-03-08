package seeders

import (
	"klikform/src/infras/databases/postgresql"
	"log"
)

func Seeders() {
	// ensure database is initialized
	if postgresql.DB == nil {
		log.Fatal("Database not initialized")
	}

	// run all seeders
	RoleSeed(postgresql.DB)
	UserSeed(postgresql.DB)

	// write log
	log.Println("Seeder data is success")
}
