package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceUsers struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID        uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User          Users      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user"`
	WorkspaceID   uuid.UUID  `gorm:"type:uuid;" json:"workspace_id"`
	Workspace     Workspaces `gorm:"foreignKey:WorkspaceID;references:ID;constraint:OnDelete:CASCADE" json:"workspace"`
	IsOwner       bool       `gorm:"default:false" json:"is_owner"`
	AccessControl string     `gorm:"type:char(2);default:'A1';comment:'A1=VIEW_ONLY,A2=FULL_ACCESS'" json:"access_control"`
	Remark        string     `gorm:"type:text" json:"remark"`
	Deleted       bool       `gorm:"default:false" json:"deleted"`
	CreatedAt     time.Time  `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"type:timestamp" json:"updated_at"`
}

func (WorkspaceUsers) TableName() string {
	return "workspace_users" // name of table will created
}

func (workspace_user *WorkspaceUsers) BeforeCreate(tx *gorm.DB) (err error) {
	workspace_user.ID = uuid.New()
	return
}
