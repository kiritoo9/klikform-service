package repos

import (
	"errors"
	"klikform/src/applications/models"
	"klikform/src/infras/databases/postgresql"

	"gorm.io/gorm"
)

func GetUsers(page int, limit int, keywords string) (any, error) {
	var users []models.Users
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	offset := (page - 1) * limit
	query := postgresql.DB.Where("deleted = ?", false)
	if keywords != "" {
		query = query.Where("LOWER(email) LIKE ?", "%"+keywords+"%")
	}
	result := query.Offset(offset).Limit(limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil

}

func GetCountUser(keywords string) (any, error) {
	var count int64
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	query := postgresql.DB.Model(&models.Users{}).Where("deleted = ?", false)
	if keywords != "" {
		query = query.Where("LOWER(email) LIKE ?", "%"+keywords+"%")
	}
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	return count, nil
}

func GetUserById(id string) (*models.Users, error) {
	var user models.Users
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	result := postgresql.DB.Where("deleted = ? AND id = ?", false, id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("Data is not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByEmail(email string, optionalID ...string) (*models.Users, error) {
	var user models.Users
	var id string
	if len(optionalID) > 0 {
		id = optionalID[0]
	} else {
		id = ""
	}

	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	query := postgresql.DB.Where("deleted = ? AND email = ?", false, email)
	if id != "" {
		query = query.Where("id != ?", id)
	}
	result := query.First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Data is not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func GetRoleByUser(user_id string) (*models.UserRoles, error) {
	var userRole models.UserRoles
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	err := postgresql.DB.
		Preload("Role").
		Where("deleted = ? and user_id = ?", false, user_id).
		First(&userRole).Error
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}

func CreateUser(user models.Users, userRole models.UserRoles) (*models.Users, error) {
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}

	err := postgresql.DB.Transaction(func(tx *gorm.DB) error {
		// insert user
		if err := tx.Create(&user).Error; err != nil {
			return err // rollback
		}

		// insert user role
		if err := tx.Create(&userRole).Error; err != nil {
			return err // rollback
		}

		return nil // commit
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *models.Users, userRole map[string]any) (*models.Users, error) {
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}

	err := postgresql.DB.Transaction(func(tx *gorm.DB) error {
		// update user
		if err := tx.Save(&user).Error; err != nil {
			return err // rollback
		}

		// insert user role
		if userRole != nil {
			if _, ok := userRole["updated_at"]; ok {
				if err := tx.Model(&models.UserRoles{}).Where("user_id = ?", user.ID).Updates(userRole).Error; err != nil {
					return err // rollback
				}
			} else {
				if err := tx.Create(userRole).Error; err != nil {
					return err // rollback
				}
			}
		}

		return nil // commit
	})

	if err != nil {
		return nil, err
	}
	return nil, nil
}
