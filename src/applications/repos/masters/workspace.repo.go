package repos

import (
	"errors"
	"klikform/src/applications/models"
	"klikform/src/infras/databases/postgresql"

	"gorm.io/gorm"
)

func GetWorkspaces(page int, limit int, keywords string, userID string) (any, error) {
	var workspaces []models.Workspaces
	offset := (page - 1) * limit

	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	query := postgresql.DB.Model(&models.Workspaces{}).Where("deleted = ?", false)
	if userID != "" {
		query = query.Where("id IN (SELECT workspace_id FROM workspace_users WHERE deleted = ? AND user_id = ?)", false, userID)
	}
	if keywords != "" {
		query = query.Where("(LOWER(title) LIKE ? OR LOWER(description) LIKE ?)", "%"+keywords+"%", "%"+keywords+"%")
	}
	result := query.Offset(offset).Limit(limit).Find(&workspaces)
	if result.Error != nil {
		return nil, result.Error
	}

	return &workspaces, nil
}

func GetCountWorkspace(keywords string, userID string) (any, error) {
	var count int64
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	query := postgresql.DB.Model(&models.WorkspaceUsers{}).Where("deleted = ?", false)
	if userID != "" {
		query = query.Where("id IN (SELECT workspace_id FROM workspace_users WHERE deleted = ? AND user_id = ?)", false, userID)
	}
	if keywords != "" {
		query = query.Where("(LOWER(title) LIKE ? OR LOWER(description) LIKE ?)", "%"+keywords+"%", "%"+keywords+"%")
	}
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return count, nil
}

func GetWorkspaceById(id string, userID string) (*models.Workspaces, error) {
	var workspace models.Workspaces
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	query := postgresql.DB.Where("deleted = ? AND id = ?", false, id)
	if userID != "" {
		query = query.Where("id IN (SELECT workspace_id FROM workspace_users WHERE deleted = ? AND user_id = ?)", false, userID)
	}
	result := query.First(&workspace)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("Data is not found")
		}
		return nil, result.Error
	}
	return &workspace, nil
}

func CreateWorkspace(workspace models.Workspaces, workspaceUser models.WorkspaceUsers) (*models.Workspaces, error) {
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}

	err := postgresql.DB.Transaction(func(tx *gorm.DB) error {
		// insert workspace
		if err := tx.Create(&workspace).Error; err != nil {
			return err
		}

		// insert workspace user
		if err := tx.Create(&workspaceUser).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func UpdateWorkspace(workspace *models.Workspaces) (*models.Workspaces, error) {
	if postgresql.DB == nil {
		return nil, errors.New("Database not initialized")
	}
	err := postgresql.DB.Transaction(func(tx *gorm.DB) error {
		// update workspace
		if err := tx.Save(&workspace).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return workspace, nil
}
