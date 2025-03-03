package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoles struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      Users     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user"`
	RoleID    uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	Role      Roles     `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE" json:"role"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (UserRoles) TableName() string {
	return "user_roles" // name of table will created
}

func (user_role *UserRoles) BeforeCreate(tx *gorm.DB) (err error) {
	user_role.ID = uuid.New()
	return
}
