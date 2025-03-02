package postgresql

import (
	"fmt"
	"klikform/src/infras/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	configs := configs.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.DB_HOST, configs.DB_USER, configs.DB_PASS, configs.DB_NAME, configs.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open((dsn)), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect database", err)
		return nil, err
	}

	return db, nil
}
