package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceAttachments struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	WorkspaceID uuid.UUID  `gorm:"type:uuid;" json:"workspace_id"`
	Workspace   Workspaces `gorm:"foreignKey:WorkspaceID;references:ID;constraint:OnDelete:CASCADE" json:"workspace"`
	FileName    string     `gorm:"type:text" json:"file_name"`
	FileSize    string     `gorm:"type:text" json:"file_size"`
	FileExt     string     `gorm:"type:text" json:"file_ext"`
	FilePath    string     `gorm:"type:text" json:"file_path"`
	Deleted     bool       `gorm:"default:false" json:"deleted"`
	CreatedAt   time.Time  `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp" json:"updated_at"`
}

func (WorkspaceAttachments) TableName() string {
	return "workspace_attachments" // name of table will created
}

func (workspace_attachment *WorkspaceAttachments) BeforeCreate(tx *gorm.DB) (err error) {
	workspace_attachment.ID = uuid.New()
	return
}
