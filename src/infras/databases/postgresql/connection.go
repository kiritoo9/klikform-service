package postgresql

import (
	"fmt"
	"klikform/src/infras/configs"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func OpenDB() {
	once.Do(func() {
		configs := configs.LoadConfig()

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			configs.DB_HOST, configs.DB_USER, configs.DB_PASS, configs.DB_NAME, configs.DB_PORT,
		)

		var err error
		DB, err = gorm.Open(postgres.Open((dsn)), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect database", err)
		}

		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatal("Failed to get database instance", err)
		}
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(50)
		log.Println("Database connected successfully")
	})
}
