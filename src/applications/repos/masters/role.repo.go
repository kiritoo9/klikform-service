package repos

import (
	"errors"
	"klikform/src/applications/models"
	"klikform/src/infras/databases/postgresql"
)

func GetRoles(page int, limit int, keywords string) (any, error) {
	offset := (page - 1) * limit
	var roles []models.Roles
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}

	query := postgresql.DB.Where("deleted = ?", false)
	if keywords != "" {
		query.Where("LOWER(name) LIKE ?", "%"+keywords+"%")
	}

	result := query.Offset(offset).Limit(limit).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return &roles, nil
}

func GetCountRoles(keywords string) (any, error) {
	var count int64
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}

	query := postgresql.DB.Model(&models.Roles{}).Where("deleted = ?", false)
	if keywords != "" {
		query.Where("LOWER(name) LIKE ?", "%"+keywords+"%")
	}

	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return count, nil
}
