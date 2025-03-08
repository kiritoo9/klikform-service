package repos

import (
	"errors"
	"klikform/src/applications/models"
	"klikform/src/infras/databases/postgresql"

	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	result := postgresql.DB.Where("deleted = ? AND email = ?", false, email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Data is not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
